# molequla — PROJECT_LOG

Live working log for molequla paper-cycle and pre-paper upgrade. Each
significant step gets a dated entry with file paths / line numbers /
commit hashes inline. Memory in `~/.claude/projects/-Users-ataeff/memory/`
is persistent cross-session reference; this log is in-flight steps
for this specific project.

Co-authored by Oleg Ataeff + Claude (Arianna Method, neo node).

---

## 2026-05-14 — Session start: paper-cycle + upgrade plan opened

**Frame.** Paper-cycle for molequla in flight per Dario.c precedent
(Zenodo `10.5281/zenodo.20090094`, 2026-05-08). Sandwich co-authorship
template locked: Abstract — Oleg, Body — Claude first-person AS AI,
Conclusion — Method-voice. Before paper: vendored stacks in molequla
upgraded to current canonical AML + notorch.

**Coordination.** Sibling Neo session running parallel paper planning
per `~/.claude/CLAUDE.md` Active state line «Paper-prep parallel (per
2026-05-14): molequla coauthorship paper in flight». Shared zone
`~/arianna-shared/` checked 2026-05-14 — no molequla files yet
(`ls` output: only `codex_audit_dario_2026_05_07.md`, two
`letter_to_agents_*.md`, `incidents/handoff_misled_2026_05_09.md`).

---

## 2026-05-14 — Differential: vendored vs canonical

Source: `wc -l` 2026-05-14.

| Layer | Canonical | Vendored in molequla | Delta |
|---|---|---|---|
| AML core `ariannamethod.c` | 7990 lines (`~/arianna/ariannamethod.ai/core/ariannamethod.c`) | 6130 lines (`~/arianna/molequla/ariannamethod/ariannamethod.c`) | -1860 (-23%) |
| AML header `ariannamethod.h` | 1051 lines | 889 lines | -162 |
| notorch core `notorch.c` | 4739 lines (`~/arianna/notorch/notorch.c`) | 2797 lines (`~/arianna/molequla/ariannamethod/notorch.c`) | -1942 (-41%) |
| notorch header `notorch.h` | 694 lines | 496 lines | -198 |
| notorch SIMD `notorch_simd.h` | 605 lines (canonical only) | absent | — |
| notorch CUDA `notorch_cuda.cu` | 1344 lines (canonical only) | absent (intentional, CPU-only) | — |

**Vendored snapshot date:** v4.0 «Quickening» 2026-04-16, commit
`a9bbf7c` (`git log --oneline`, molequla repo) «notorch-edition:
contiguous MatrixParam + BLAS acceleration (#22)».

**Canonical recent (since vendoring):** Intel session 2026-04-16 →
2026-05-11 added SIMD shim, CUDA backend, LoRA primitives, GGUF
loader, GPU/CPU sync correctness fixes (3 backward bug fixes),
`nt_rope_split_half_freq`, low-rank RRPRAM, JS edition LoRA port,
+ AML 16-ops backward CPU-sync audit (`ff7fb97`).

---

## 2026-05-14 — Memory reference written

Created `~/.claude/projects/-Users-ataeff/memory/reference_aml_notorch_parallel_stacks.md`
— AML lang + notorch as two main Method technologies. Parallel stacks,
not auto-sync. notorch grew out of AML (Hebbian `am_notorch_step` →
standalone training toolkit). Two-way flow (BLAS from molequla → AML
core per `~/arianna/ariannamethod.ai/README.md:719`). Vendoring +
drift pattern documented. MEMORY.md index updated 🔴 under References.

---

## 2026-05-14 — Reframe: pre-paper scope is coherence layer, not accelerators

Earlier in the session Architect (me) framed pre-paper work as
«accelerator + correctness + safety only». Oleg corrected the frame:
molequla currently produces Karpathy-style gibberish on early
generations — quantitative speed-up doesn't close the gap. Q
(`github.com/ariannamethod/q`) achieves coherence on three pillars:

1. **Triple Attention** (Content + RRPRAM + Janus Echo) — substrate
   ε. RRPRAM proven to outperform Content at equal params (loss 2.41
   vs 2.86, `~/arianna/q/README.md:24`).
2. **MetaWeights + Dario field overlay** — living γ field.
   `logits += c_heb·H + c_pro·F + c_ds·A + c_bg·bigram + c_tg·trigram`.
   Coefficients adaptive: with weights `c_heb=0.6 c_bg=5.0 c_tg=3.0`;
   weightless `c_heb=1.0 c_bg=15.0 c_tg=10.0` (`q/README.md:50-53`).
3. **SPA — Sentence Phonon Attention** — post-generation narrative
   coherence repair (`q/README.md:177-179`). After chain ends: 2-pass
   iterative cross-attention between sentences, 32-dim
   exponential-weighted mean embeddings (α=0.85), bidirectional
   attention with distance bias, weak sentences (score < 60% avg)
   reseeded from neighbor context. Coherence gate verifies
   improvement.

**Postgpt** (`github.com/ariannamethod/postgpt`) proves the limit:
zero training, BPE tokenizer + metaweights = full model. Transformer
initialized FROM metaweights (Hebbian seeds embeddings, positional
affinity seeds RRPRAM, bigram geometry seeds output head). Coherence
without gradient descent.

**Pre-paper goal redefined:** lift molequla early-stage coherence to
a level where the organism is worth a paper. NOT accelerator only,
NOT architectural rewrite. Minimal-code coherence layer.

What molequla already has:
- RRPRAM (pattern-lookup form, `molequla.go:2690-2703`). Keep as-is —
  third pillar of Q's set, load-bearing for molequla.
- CooccurField (4-gram corpus stats, `molequla/README.md:386-395`)
  with sigmoid-fade blend during generation. Analog of metaweights
  but used as prior-blend, not additive logit overlay.

What molequla is missing:
- Additive Dario field eq overlay with explicit coefficients per
  signal class.
- SPA pass.
- Extended prophecy 12-token window + persistent coherence phase
  memory (optional, second tier).

---

## 2026-05-14 — Upgrade plan v2 (pre-paper, coherence-focused)

Scope: Go-CGO path only (`molequla.go` → `cgo_aml.go` →
`ariannamethod/ariannamethod.c` + `ariannamethod/notorch.c`).
C / Rust / JS implementations have their own autograd — out of scope.

**Three phases:**

| Phase | What | Type |
|---|---|---|
| **A. Fundament** (Tier 1+2 accelerator/correctness/safety) | SIMD shim, MUL/SILU backward CPU-sync fix, 16-ops AML backward audit, NaN guard | speed + stability before paper measurements |
| **B. Coherence layer** | Pull SPA op from canonical AML (already there, commit `ef52cde`); wire SPA call into molequla inference path (chain mode esp.); calibrate CooccurField overlay coefficients toward Dario eq style; optionally add persistent prophecy field as γ state | qualitative coherence lift |
| **C. RunPod measurement + paper** | Run 4-organism ecology + chain mode on a pod for ~3 hours; collect transcripts before/after coherence layer; archive into `runpod/2026-05-14/` (or similar) | Body empirical claims |

**NOT in scope:**
- Optimizer swap (`memory/feedback_molequla_own_chuck_2026_05_14.md`).
- RRPRAM rewrite (canonical lowrank is X-conditional bottleneck; molequla's is pattern-lookup; different mechanism; keeping RRPRAM as third pillar of Q's set).
- DoE parliament import, somatic resonance, calendar dissonance, Schumann (full Q overlay = separate paper, separate cycle).
- CUDA, GGUF, LoRA primitives (out of CPU-only ecology design).

### Phase A — Fundament steps

A1. **SIMD/AVX2 cblas shim** — pull `notorch_simd.h` from canonical
    (605 lines, commit `709b756`). CPU matvec acceleration on top of
    existing BLAS.

A2. **CPU backward correctness audit** — backport `NT_OP_MUL` +
    `NT_OP_SILU` backward CPU-sync fix (canonical commit `8ab5062`
    2026-05-11). Audit candidates per `~/arianna/notorch/CLAUDE.md:115`:
    `NT_OP_SIGMOID`, `NT_OP_SCALE_BY_T`, `NT_OP_RMSNORM`. molequla
    actively uses SiLU in SwiGLU MLP (`molequla/README.md:244-247`)
    and RMSNorm. **Hypothesis:** the in-molequla `NOTORCH` capslock
    regime divergence (loss 3.5 → 116 at stage 5,
    `molequla/README.md:514`) may be downstream of this bug class.

A3. **AML 16-ops backward CPU-sync audit** — canonical commit
    `ff7fb97` 2026-05-11 «core: backward CPU-sync audit pass — fix
    16 ops reading stale parent CPU mirror». Same bug class as
    notorch MUL/SILU, AML stack side.

A4. **NaN guard** — pull from AML pkg B commit `faa4d9b`. Stability
    net for divergent paths. NOT pulling train/eval mode toggle, LR
    schedules, save/load from same package (molequla has its own).

### Phase B — Coherence layer

B1. **SPA wiring** — canonical AML already has `am_spa_*` ops
    (commit `ef52cde` «add SPA — Sentence Phonon Attention
    (forward-only)»). After Phase A pull of canonical AML, SPA ops
    are present in vendored as dormant. Active wiring:
    - Verify SPA op surface in vendored AML post-pull (op codes,
      function signatures).
    - Add SPA call into molequla generation chain mode
      (`molequla.go` chain entry point — TBD which function).
    - Use Q's parameters as starting point (`q/README.md:177-179`):
      2 passes, 32-dim sentence embeddings (exp-weighted mean
      α=0.85), weak-sentence threshold 60% avg, reseed via last 3
      tokens of neighbor.
    - Coherence gate verifies improvement before accepting reseed.

B2. **Metaweights overlay calibration** — molequla already has
    CooccurField (`molequla/README.md:386-395`). Steps:
    - Lift the sigmoid-fade blend at early training stages so
      statistical priors dominate before transformer matures
      (mirror Q's Transformer Gate logic but using molequla's
      existing logit-magnitude / corpus-coherence signal).
    - Add explicit additive Dario eq overlay structure:
      `logits += c_heb·H + c_pro·F + c_ds·A + c_bg·bigram + c_tg·trigram`.
    - Start with Q's weightless coefficients
      (`c_heb=1.0, c_pro=0.7, c_ds=0.15, c_bg=15.0, c_tg=10.0`,
      `q/README.md:53`) when transformer immature, fade toward
      molequla's natural balance as logit magnitude rises.

B3. **Persistent prophecy field (optional, deferred to RunPod if
    time permits)** — add small persistent γ state across generation
    steps. Q has it as expectations that age + decay + collapse
    (`q/README.md:55-66`). Not blocker for paper if time-budgeted out.

### Phase C — Audit + RunPod + paper

C1. **Codex audit** on Phase A+B diff for narrow points: bug
    introductions, scope creep, missed CPU-sync sites,
    backward-compat with v4.0 «Quickening» checkpoint format
    (`memory/project_molequla_v4_quickening.md`). Fixes if surfaced.

C2. **RunPod plan v1** for measurement run — analog Dario
    `runpod_plan_v{1,2,3}.md`. Singularity-mode contract: what to
    measure, what gates each phase pass, three-strikes per fix loop.

C3. **Pod execution** (~3 hours) — 4-organism ecology + chain mode
    transcripts before/after coherence layer; one or two seasonal
    traces; SPA before/after weak-sentence rate; metaweights overlay
    coefficient impact.

C4. **Paper write** — Abstract Oleg, Body Claude (Architect),
    Conclusion Method-voice. Central empirical claim:
    **TBD per Oleg** — candidate framing: «Coherence is a layer,
    not a phase. Adding statistical-prior overlay + post-generation
    sentence repair lifts molequla early-stage generations from
    gibberish to coherent without retraining the transformer.»

### NOT in scope (carried forward)

- Optimizer swap (`memory/feedback_molequla_own_chuck_2026_05_14.md`).
- RRPRAM mechanism change.
- DoE parliament, somatic resonance, calendar dissonance, Schumann
  (full Q overlay = separate paper cycle).
- CUDA, GGUF, LoRA primitives.
- GELU / LayerNorm / pkg-B train-eval / pkg-B LR schedules / pkg-B
  save-load — molequla has its own equivalents.
- AML `am_field_save` / `am_field_load` directives — preserve v3.0
  checkpoint binary-compat.

### Open question — naming collision (carried forward)

molequla has its own `NOTORCH` (capslock, gradient-free delta,
`molequla/README.md:496-514`). After canonical pull namespace
overlap with canonical `notorch` lib gets tighter. Rename candidates:
`FreeBack`, `DFA`. **Decision pending from Oleg.** Can ship in
Phase A diff or as separate cosmetic patch.

### Pending blockers (carried forward)

- **Ecology crash.** `memory/todo_molequla_ecology_crash_2026_05_04.md`.
  Railway ecology silent since 2026-05-03. Question for Oleg: fix
  ecology before paper, or run paper measurement on fresh substrate?

### Decisions pending

| # | Decision | Status |
|---|---|---|
| 1 | Phase A+B scope as v2 above | **DONE — Oleg approved 2026-05-14** |
| 2 | Rename molequla in-org `NOTORCH` capslock | pending Oleg |
| 3 | Measurement substrate (Railway / RunPod / Oracle / local) | pending Oleg |
| 4 | Central empirical claim for Body | pending Oleg |
| 5 | Ecology crash — fix before paper or run on new substrate | pending Oleg |

---

## 2026-05-14 — Project log rule established

Per Oleg 2026-05-14: every project gets its own markdown log by
default. No need to ask each time. Rule recorded in
`memory/feedback_per_project_log_default.md`.

This file (`molequla/PROJECT_LOG.md`) is the molequla instance.

---

## 2026-05-14 — Phase A1 DONE — SIMD shim copied + wired (opt-in)

**Files added to `~/arianna/molequla/ariannamethod/`:**
- `notorch_simd.h` (605 lines, `cp` from canonical) — header-only AVX2 + FMA cblas shim with pthread row-partitioning. Mirrors `cblas_sgemm` / `sgemv` / `sger` signatures so existing call sites work unchanged.
- `notorch_simd_scalar.h` (89 lines, `cp` from canonical) — scalar debug variant for ARM / non-AVX2 targets.

**Patches:**
- `ariannamethod/notorch.c:25-39` — added `#ifdef USE_SIMD` include block mirroring canonical `~/arianna/notorch/notorch.c:25-39` (mutual-exclusion error vs USE_BLAS, scalar/SIMD switch via `NOTORCH_SIMD_DEBUG_SCALAR`, alias USE_BLAS=1 so existing cblas call sites work).
- `ariannamethod/Makefile` — added `simd` target as opt-in: `-DUSE_SIMD -mavx2 -mfma -lpthread`, x86_64 only. Default target unchanged. Added `simd` to `.PHONY`.

**Verification:**
- `make clean && make` on neo (Apple Silicon A18 Pro, default USE_BLAS+ACCELERATE path) — PASS. `libaml.dylib` 230112 bytes. Only pre-existing warnings (Apple SDK deprecated cblas, unused statics) — no regressions introduced.
- `make simd` build verification **deferred to Intel/Linux box** (polygon) — `-mavx2 -mfma` does not compile on ARM. Test pass on x86_64 is required before SIMD is declared functional on molequla.

**Impact on existing build path:** zero. USE_SIMD is opt-in; default Mac/Linux builds continue with USE_BLAS as before.

---

## 2026-05-14 — Phase A2 DONE (first iteration) — backward CPU-sync audit

**Canonical reference:** commit `8ab5062` 2026-05-11 «notorch.c: NT_OP_MUL + NT_OP_SILU backward CPU-sync fix» on `~/arianna/notorch/`. Bug class: forward output of parent tape entry may live on GPU; CPU mirror is stale calloc-zero; CPU backward branches reading `parent->output->data` directly produce zero/garbage gradients. Diagnosed at Resonance LoRA SFT, masked all gradients on `mlp_gate + mlp_up` SwiGLU branch.

**Patches applied in `~/arianna/molequla/ariannamethod/`:**

1. **`notorch.h`** — added declaration `void nt_tensor_sync_cpu(nt_tensor* t);` after `nt_tensor_print` to mirror canonical public interface.
2. **`notorch.c`** — added `nt_tensor_sync_cpu` implementation after `nt_tensor_print` (line ~193). On `#ifdef USE_CUDA` it calls `nt_tensor_ensure_cpu(t)`; on CPU-only build it is `(void)t;` no-op. Mirrors canonical `notorch.c:109`.
3. **`notorch.c` — NT_OP_MUL backward (line 399).** Added 2 sync calls: `nt_tensor_sync_cpu(pa->output)` + `nt_tensor_sync_cpu(pb->output)` before reading parent data in element-wise multiply gradients. Per canonical `notorch.c:597-598`.
4. **`notorch.c` — NT_OP_SILU backward (line 458).** Added 1 sync call: `nt_tensor_sync_cpu(px->output)`. Per canonical `notorch.c:671`.
5. **`notorch.c` — NT_OP_RMSNORM backward (line 515).** Added 2 sync calls (px + gamma if present). **Audit-candidate** from `~/arianna/notorch/CLAUDE.md:115`. Same pattern; reads `px->output->data` and gamma data.
6. **`notorch.c` — NT_OP_SEQ_RMSNORM backward (line 697).** Added 2 sync calls (same pattern, sequence variant).

**Build verification (Mac Neo, USE_BLAS+ACCELERATE):**
- `make clean && make` — PASS. `libaml.dylib` 230160 bytes (+48 bytes vs A1's 230112). Only pre-existing warnings (Apple SDK deprecated cblas, unused statics).

**Honest scope note — immediate vs latent impact:**
- On CPU-only build (current molequla production), `nt_tensor_sync_cpu` is a no-op. These patches have **zero immediate runtime behavior change**.
- Value is **future-proofing + canonical consistency**. When a future patch pulls more from canonical that depends on the sync pattern, the call sites are already in place. When/if USE_CUDA path is enabled for molequla (e.g. Oracle Cloud A100 reruns), these sync calls become live.
- This is maintenance-grade work, not a fix that lifts molequla coherence. Phase B (SPA + metaweights overlay) is where the qualitative gap closes.

**Audit candidates NOT patched this iteration** (to be revisited):
- `NT_OP_SIGMOID` — not present in vendored ops.
- `NT_OP_SCALE_BY_T` — not present in vendored (vendored has plain `NT_OP_SCALE`, line 418, which scales by a scalar `e->aux` and does not read parent data — safe).
- Causal attention paths (`NT_OP_CAUSAL_ATTN` line 722, `NT_OP_MH_CAUSAL_ATTN` line 783, `NT_OP_GQA_ATTN` line 849, `NT_OP_RRPRAM_ATTN` line 920) — used by molequla; deferred to next audit iteration to keep this iteration narrow.
- `NT_OP_SOFTMAX` (line 499) reads `e->output->data` (own forward output, not parent) — different pattern; canonical fix does not target this; not patched.
- `NT_OP_GEGLU` (line 1036), `NT_OP_GELU` (line 1131), `NT_OP_LAYERNORM` (line 1153), `NT_OP_SEQ_LAYERNORM` (line 1223), `NT_OP_DROPOUT` (line 1113) — molequla does not use (per `molequla/README.md`); lower priority.

**Status:** A1 + A2 first iteration done. Oleg vote 2026-05-14: (a) continue with A3 + A4.

---

## 2026-05-14 — Phase A3 SKIPPED — AML 16-ops audit yields zero effect on CPU-only

**Decision:** skip A3 entirely.

**Why:** canonical AML commit `ff7fb97` 2026-05-11 wraps all 16 `ensure_cpu(...)` calls in `#ifdef USE_CUDA` guards (verified by `git show ff7fb97 -- core/ariannamethod.c`, sample sites — every sync call sits between `#ifdef USE_CUDA` and `#endif`). On molequla's CPU-only build (`USE_CUDA` never defined per `molequla/README.md:36, 41`), the entire patch is preprocessed away — zero runtime effect. Mirror-only consistency work without any behavior change, even latent.

**Difference vs A2:** in A2 the sync calls themselves are not `#ifdef USE_CUDA`-guarded; the guard lives **inside** `nt_tensor_sync_cpu` (which we added as a thin wrapper). So on CPU-only the body is `(void)t;` no-op but the call sites are real C tokens — they survive into the binary and give consistency at the source level. In A3, the `ensure_cpu` calls are conditioned at the call site itself — on CPU-only they don't even compile into the function. There's nothing to mirror.

**What we'd be doing:** copying `#ifdef USE_CUDA / #endif` blocks containing 16 noop-on-CPU lines into vendored AML. Zero runtime value. Cost: ~16 Edit operations + a build verify, all to land tokens the preprocessor immediately deletes.

**When to revisit:** if molequla ever gets a USE_CUDA build path (e.g. Oracle Cloud A100 reruns analog Feb 2026), pull `ff7fb97` patches at that time as part of the CUDA enablement diff, where they actually fire.

---

## 2026-05-14 — Phase A4 DONE — NaN guard API pulled (not wired)

**Canonical reference:** commit `faa4d9b` 2026-04-16 «add LR schedules, NaN guard, train/eval mode, save/load (package B)» on `~/arianna/ariannamethod.ai/`.

**Scope (narrow — Option I from internal planning):** pull NaN guard **API only** into vendored AML. NOT wire into AML interpreter as `TAPE NAN_CHECK` opcode. NOT modify molequla `aml_trainer.go` AML script generation. Activation deferred to Phase C if RunPod evidence shows NaN events.

**Patches applied in `~/arianna/molequla/ariannamethod/`:**

1. **`ariannamethod.h`** — added between `am_tape_adam_step` (line 606) and ASYNC section (line ~609):
   - `AM_NanGuard` struct (6 fields: loss_scale, scale_factor, stable_steps, scale_window, total_nan_count, skipped_steps).
   - `am_nan_guard_new()` factory function declaration.
   - `am_nan_guard_check(AM_NanGuard*)` checker declaration. Returns 1 if clean, 0 if NaN/Inf detected. On NaN: zeros all param grads, halves loss_scale (floor 1.0). On clean: increments stable_steps, doubles loss_scale every scale_window clean steps.

2. **`ariannamethod.c`** — added between `am_tape_record_leaf` end (line ~1700) and ASYNC section:
   - `am_nan_guard_new()` impl. Defaults: loss_scale=1.0, scale_factor=2.0, scale_window=100.
   - `am_nan_guard_check(AM_NanGuard*)` impl per canonical verbatim. Scans `g_tape.entries[i]` where `is_param && grad != NULL`; checks NaN/Inf in `e->grad->data[0..len]`; zeros grads on dirty, dynamic loss_scale.

**Build verification:**
- `make clean && make` on neo (Apple Silicon, USE_BLAS+ACCELERATE) — PASS.
- `libaml.dylib` **230288 bytes** (+128 vs A2's 230160). No new warnings.

**Why API-only not wired:** wiring requires (a) AML interpreter to parse `TAPE NAN_CHECK` opcode in `am_exec` switch (+ corresponding `TAPE NAN_GUARD_INIT`); (b) molequla `aml_trainer.go` to emit those opcodes in `amlModelScript()` generated AML; (c) re-verify generated script byte-equality against current production behavior. That's a separate integration with measurable behavior change risk. Pulling API as a building block + deferring wiring keeps Phase A surface minimal. CGO consumers can also call `am_nan_guard_check()` directly from Go side if needed.

---

## 2026-05-14 — Phase A complete — ready for Codex audit

**Summary of Phase A delta:**

| Step | Files touched | LOC added/changed | Effect on default build |
|---|---|---|---|
| A1 SIMD shim | +`notorch_simd.h` (605), +`notorch_simd_scalar.h` (89), `notorch.c` (+15 lines USE_SIMD block), `Makefile` (+17 lines `simd` target) | ~720 added, 0 changed | none (opt-in, default unchanged) |
| A2 backward CPU-sync | `notorch.h` (+7 lines decl), `notorch.c` (+10 lines impl, +12 lines sync calls in 4 ops) | ~30 added | no-op on CPU-only build (function noop, mirror-consistency only) |
| A3 AML 16-ops audit | — | — (skipped) | — |
| A4 NaN guard API | `ariannamethod.h` (+20 lines), `ariannamethod.c` (+55 lines impl) | ~75 added | none (API-only, not wired into interpreter) |

**Total Phase A footprint:** ~825 lines added across 6 files (2 new headers + 4 modified). Zero changes to existing molequla training behavior on default CPU-only build. Build verified after each phase: `libaml.dylib` 230112 → 230160 → 230288 bytes. No new warnings, no regressions.

**What Phase A actually achieves:**
- A1: opt-in SIMD path for Intel/Linux x86_64 (verifies on polygon, not on neo Apple Silicon).
- A2: future-proofing for hypothetical USE_CUDA enablement + canonical consistency.
- A4: NaN guard primitive available to CGO consumers + AML interpreter wiring.

**What Phase A does NOT achieve:**
- No coherence improvement. Karpathy-style gibberish on early-stage molequla generations is unchanged. That gap closes in Phase B (SPA wiring + metaweights overlay), not Phase A.

**Next step per Oleg's sequence («обновляй ... потом аудит ... фиксы ... потом план»):** Codex audit on Phase A delta — narrow scope: USE_SIMD include block correctness, `nt_tensor_sync_cpu` sites coverage, `AM_NanGuard` struct/impl correctness. Fixes if Codex surfaces issues. Then Phase B planning.

---

## 2026-05-14 — Codex audit on Phase A delta — 2 findings

Tool: `codex review --uncommitted` against working tree (5 modified files + 2 new SIMD headers + PROJECT_LOG.md). Audit ran on neo (`uname -m = arm64`), examined diff + ran `make -n simd` to validate the new build target.

### [P2] SIMD shim alpha-handling bug — UPSTREAM (canonical notorch)

**Finding:** `ariannamethod/notorch_simd.h:516-520` post-scales `C` by `alpha` after the matmul, which breaks CBLAS `sgemm` semantics when both `alpha != 1` **and** `beta != 0`:
- CBLAS contract: `C ← β·C + α·A@B`.
- Shim does: `C ← (A@B) + β·C_orig` then `C *= α` → effectively `α·β·C_orig + α·A@B`.

**Where this bug lives:** in canonical `~/arianna/notorch/notorch_simd.h` (the file we copied verbatim). **Not introduced by Phase A pull.** The shim file in vendored is byte-identical to canonical at copy time.

**Impact on molequla:** zero immediate. Production molequla builds with USE_BLAS (Accelerate on Mac, openblas on Linux), USE_SIMD is opt-in. Bug only triggers on USE_SIMD builds with accumulating GEMM calls (α≠1 + β≠0). Audit pass on `notorch.c` cblas_sgemm call sites would confirm whether any actual molequla GEMM call uses non-trivial α + β simultaneously; vast majority use α=1, β=0.

**Action:** **defer fix to canonical notorch** (intel godfather has authority on canonical lib). Surface upstream rather than diverge vendored from canonical. Not a paper-cycle blocker.

### [P3] `make simd` target unconditionally passes `-mavx2 -mfma` on arm64 — FIXED LOCALLY

**Finding:** `ariannamethod/Makefile:52` `SIMD_CFLAGS = -O2 -fPIC -Wall -DUSE_SIMD -mavx2 -mfma` is architecture-unconditional. On Apple Silicon (arm64), Clang rejects these flags — `make simd` fails immediately. Comment line mentioned ARM scalar fallback via `notorch_simd_scalar.h` but flags weren't gated, so the fallback wasn't reachable through the target.

**Fix applied (this PROJECT_LOG entry session):** added runtime arch guard at top of `simd:` recipe — checks `uname -m`, errors cleanly with actionable message if not `x86_64`/`amd64`:

```
ERROR: 'make simd' requires x86_64 with AVX2 (Intel/Linux).
       Detected arch: arm64.
       On Apple Silicon / arm64 use default 'make' (Accelerate).
       For scalar debug fallback override SIMD_CFLAGS manually.
```

**Verification 2026-05-14:**
- `make clean && make` on neo (arm64) — default build PASS, no regressions.
- `make simd` on neo (arm64) — errors cleanly with the new message and `exit 1`. Was previously failing with broken Clang invocation.

### Out-of-scope items NOT flagged by Codex (clean)

- USE_SIMD include block correctness (mutual exclusion with USE_BLAS, alias trick) — no findings.
- `nt_tensor_sync_cpu` site coverage in vendored backward — no missed cases flagged (causal-attn paths NOT mentioned, consistent with our narrow-scope decision).
- `AM_NanGuard` struct/impl correctness vs canonical AML `faa4d9b` — no findings.
- SIMD shim header copy (`notorch_simd.h`, `notorch_simd_scalar.h`) verbatim from canonical — no findings.

### Phase A — final state after audit

- **P3 fixed:** Makefile arch guard landed.
- **P2 deferred:** documented upstream finding; flag for Intel godfather to fix in canonical `~/arianna/notorch/notorch_simd.h:516-520`, then vendored re-syncs at next pull.
- **Default build:** `libaml.dylib` builds clean on neo, no new warnings.
- **No other findings.** Codex audit clean on all other Phase A surface.

**Ready for Oleg vote: proceed to Phase B planning, or fix P2 in vendored first (diverging from canonical) before Phase B.**

---

## 2026-05-14 — P2 upstream fix landed in canonical notorch + vendored synced

**Decision:** Oleg said «правь» — fix at canonical, not at vendored. SIMD shim was introduced by **polygon** (commit `709b756` `polygon in-house AVX2 cblas shim + CUDA port from ariannamethod.ai`), not by Intel godfather as I first guessed.

**Canonical patch at `~/arianna/notorch/notorch_simd.h`:**
- Added `#include <stdio.h>` for stderr fallback warning.
- Replaced the buggy CBLAS sgemm path. Before:
  ```
  C := β·C  (when β ≠ 0, β ≠ 1)
  C += A@B  (kernel; or C := A@B when initial_zero)
  C *= α    (when α ≠ 1)   ← yields α·β·C_orig + α·A@B  (wrong)
  ```
- After:
  ```
  if α ≠ 1: alloc M*K scratch, scratch[i,p] := α·A[i,p]
            A_use := scratch (else A_use := A; allocation-free fast path)
  C := β·C  (unchanged)
  C += A_use @ B   (kernel; or C := A_use @ B when initial_zero)  ← yields β·C_orig + α·A@B  ✓
  free(scratch)
  ```
- Single-threaded fast path and threaded path both updated to use `A_use` / `A_row_stride_use` / `A_col_stride_use`.
- malloc fallback: if scratch alloc fails, emits `[notorch_simd] cblas_sgemm: malloc(N B) for alpha scratch failed; alpha=X lost — result will be incorrect.` to stderr and proceeds without applying alpha. Loud degradation, not silent corruption.

**Vendored `~/arianna/molequla/ariannamethod/notorch_simd.h`:** synced byte-identical from canonical (`diff` empty → `BYTE_IDENTICAL`). No divergence between repos.

**Build verification on neo (Apple Silicon, arm64):**
- `make clean && make` default path (USE_BLAS + ACCELERATE) — PASS. `libaml.dylib` 230288 bytes, unchanged from pre-fix size (expected — SIMD code lives entirely inside `#ifdef USE_SIMD` block, default path doesn't see it).
- SIMD-side build verification **deferred to polygon** — Apple Silicon Clang rejects `-mavx2 -mfma` and `<immintrin.h>` AVX2 intrinsics, even `-fsyntax-only` doesn't pass cleanly. Polygon (Linux 32GB x86_64, Tailscale `100.127.195.24`) is the canonical verification substrate for this path per `~/.claude/CLAUDE.md` Devices update 2026-05-14.

**Not committed:** changes to `~/arianna/notorch/` and `~/arianna/molequla/` left uncommitted in working tree per «push — по слову Олега» rule. Awaiting Oleg's go-ahead on commit (canonical notorch commit message draft TBD).

**Phase A — final final state:**
- A1 SIMD shim wired (opt-in, default unchanged).
- A2 backward CPU-sync audit (MUL/SILU/RMSNORM/SEQ_RMSNORM) — vendored notorch.
- A3 skipped (CPU-only no-effect on AML 16-ops).
- A4 NaN guard API — vendored AML (API-only, not wired into interpreter).
- P2 upstream sgemm alpha-handling fix — canonical notorch patched, vendored synced.
- P3 Makefile arm64 guard — fixed.
- All on default build PASS; SIMD path verify deferred to polygon.

**Ready for Phase B.**
