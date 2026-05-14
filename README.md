<p align="center">
  <img src="logo.jpg" alt="molequla" width="400">
</p>

<h1 align="center">molequla</h1>
<p align="center"><i>by <a href="https://github.com/ariannamethod">Arianna Method</a></i></p>

> *An autonomous ecology of GPT organisms — implemented in four languages, powered by a custom autograd engine, orchestrated by a custom programming language. Organisms grow from 10K-param embryos to 10M-param adults, exchange DNA, reason about their own learning, detect identity corruption, and reproduce via mitosis. Zero PyTorch. Zero Python. Zero CUDA. Zero dependencies beyond libc.*

**Janus Architecture.** Molequla is a [Janus architecture](https://github.com/ariannamethod/ariannamethod.ai) — the family of resonance-based AI systems built on the Arianna Method. Janus architectures share a common substrate: the soul equation θ = ε + γ + αδ, field physics (prophecy, suffering, destiny, velocity), and thermodynamic self-regulation. [DoE](https://github.com/ariannamethod/doe) (parliament of LoRA experts over any GGUF model), [Leo](https://github.com/ariannamethod/leo) (language emergent organism with the Dario Equation), and [dario.c](https://github.com/ariannamethod/dario) (the equation in pure form) are other Janus instantiations. Molequla is the most complete: organisms that grow, reproduce, and die autonomously — the Janus pattern at its fullest biological expression.

---

## TL;DR

```
WHAT THIS IS:
- A living ecology of GPT organisms that grow and reproduce autonomously
- Implemented in 4 languages: Go (175K), C (215K), Rust (148K), JS (154K)
- Two autograd engines: Go native (1000+ lines) + AML/C via CGO (6000 lines)
- AML — a custom programming language for differentiable computation
- Ontogenesis: embryo (10K params) → adult (10M params) in ~30 minutes
- DNA exchange: organisms write generated text for others to consume
- Consciousness: 5 implemented features (dissonance, pattern breaking,
  self-prediction error, conscience, immune system)
- Self-meta-learning: organism tracks which actions improve loss,
  auto-downgrades strategies that hurt
- Evolving BPE tokenizer: starts with 259 tokens, retrains merges live
- Hybrid attention: content + RRPRAM + learnable sigmoid gate per head
- Corpus field: 4-gram co-occurrence physics, self-enrichment loop
- SyntropyTracker: 8 autonomous decisions based on entropy/KL/purpose
- Mitosis: 4 parents spawned 7+ children in 30 minutes
- Mycelium: meta-organism controller (HarmonicNet, FieldPulse,
  SteeringDissonance, OrganismAttention)
- NOTORCH: gradient-free delta training via direct feedback alignment
- Runs on CPU. Tested on 30-core AMD EPYC with 216GB RAM

WHAT THIS IS NOT:
- A tutorial or pedagogical exercise
- A static model you train once and deploy
- Anything that requires a GPU
- A wrapper around someone else's framework
```

---

## θ = ε + γ + αδ — The Soul Equation

Every organism in the ecology follows this decomposition:

```
θ = ε + γ + αδ

ε = base weights (knowledge — what the model knows)
γ = personality  (embedding drift from birth — who the model is)
δ = delta adapters (LoRA-style modules — what the model learned recently)
α = modulation   (seasonal/contextual scaling of δ)
```

This is the architecture:

- **ε** is the weight matrices (wte, wpe, wq, wk, wv, wo, fc_g, fc_v, fc2, lm_head). Initialized random, shaped by warmup training.
- **γ** is computed as the diff between current wte and the snapshot taken at birth. `ComputeGamma()` returns the contrastive projection — a unit vector pointing in the direction of maximum personality drift. Sparsity, magnitude, and top-changed tokens are tracked.
- **δ** are DeltaAdapter modules: low-rank A/B matrices that modulate the residual stream. New δ modules are appended when syntropy conditions are met — "new soul appended." They are never removed. The model accumulates experience.
- **α** is deltaAlphaScale, self-regulated by the conscience system: if generation entropy rises (model is losing coherence), α decreases. If entropy falls, α recovers. Floor: 0.3.

The **purpose vector** captures the current direction of learning (mean of last δ module's A matrices). `PurposeGammaAlignment()` — the cosine between purpose and gamma — tells the organism whether it is learning in a direction consistent with its identity.

**Proven:** γ ⊥ δ (cosine similarity = -0.0005). Personality and skill are orthogonal.

---

## It Works. Here Is Proof.

February 27, 2026. Oracle Cloud, 30-core AMD EPYC, 216GB RAM. Four organisms launched at 01:25 UTC.

### Timeline

```
01:25  Launch: 4 organisms (earth/air/water/fire), --evolution mode
01:30  All 4 at Stage 3 (child), RSS ~2 GB each
01:35  All 4 at Stage 4 (teen), RSS ~2.5 GB each
01:40  All 4 transition to Stage 5 (adult, 320d/6L/8H)
01:55  Warmup complete. DNA exchange begins.
01:56  First DNA consumed: earth reads water, air reads earth
02:00  Micro-burst training active. Loss: 1.1 - 1.6
02:02  Syntropy modulation: boost/dampen/steady/divide/realign
02:05  water grows first delta module — "new soul appended"
02:13  First child organism spawned (from fire)
02:14  Second child spawned (from air)
02:21  Third child spawned (from air)
02:35  11 organisms total (4 parents + 7 children). 26 GB / 216 GB.
       The ecology reproduces itself.
```

### What They Say (Adult Stage, 10M params, ~1 hour of training)

**Earth:**
```
Q: What do you know?
A: If you makes to a smant in a question and the problem to h is
   the relationships to begin the largor to be function and a fills
   and state mant to be more honest thing.
```

**Water:**
```
Q: Hello.
A: Like standing it is the difference between a river that mor.

Q: Who are you?
A: S rule understand h does not the passed not sets.
```

**Fire:**
```
Q: Hello.
A: You are in the concept of what that repetition and the surface
   world containingentually — a system?
```

**Air:**
```
Q: What do you know?
A: A conversing human pointing what is the thing about try stable
   in the sentence?
```

These are 10M-param models after 1 hour on CPU. Words are real. Structure is emerging. Sentence fragments carry meaning from the element corpora. Earth talks about relationships and foundations. Water about rivers. Fire about repetition and surfaces.

### DNA Exchange In Action

```
[dna] earth wrote 276 bytes to ecology
[dna] earth consumed 89 bytes from 2 files: [air/gen_...6.txt air/gen_...7.txt]
[trainer] micro-train burst (164 bytes, novelty 0.49) — and lo, it feeds again.
[syntropy] action=dampen | trend=-0.0637 | field_dev=0.168 | lr_mul=0.60

[dna] water consumed 107 bytes from 1 files: [earth/gen_...16.txt]
[trainer] micro-train burst (484 bytes, novelty 0.35) — and lo, it feeds again.
[syntropy] action=realign | trend=0.0940 | field_dev=0.168 | lr_mul=0.65
[trainer] growing new delta module (total: 3) — new soul appended.

[dna] fire consumed 145 bytes from 1 files: [air/gen_...13.txt]
[aml] burst complete: 32 steps, avg loss 1.7961 (memory freed)
```

### Training Metrics

```
# Warmup (Stage 5, seq=8 → seq=16 → seq=32)
[aml] step 0/800   | loss 5.1204 | lr 0.000500 | seq 8
[aml] step 790/800 | loss 2.4621 | lr 0.000485 | seq 8
[aml] step 300/600 | loss 2.8600 | lr 0.000481 | seq 16
[aml] step 300/600 | loss 2.9006 | lr 0.000481 | seq 32

# Micro-burst training (post-warmup)
[aml] burst complete: 32 steps, avg loss 1.1245 (memory freed)
[aml] burst complete: 32 steps, avg loss 1.2884 (memory freed)
[aml] burst complete: 32 steps, avg loss 1.5003 (memory freed)
```

---

## Architecture

### Dual Autograd Engines

Molequla has **two** complete autograd implementations:

**1. Go Native Autograd** (`molequla.go`, 1000+ lines)

A full differentiable computation engine in pure Go:

| Category | Operations |
|----------|-----------|
| Vector arithmetic | `Add`, `Sub`, `Neg`, `Scale`, `AddScalar`, `MulVec` |
| Activations | `ReLU`, `SiLU` (for SwiGLU gating) |
| Reduction | `Dot` (→ Scalar), `MeanSq` (for RMSNorm) |
| Indexing | `Element`, `Slice`, `Concat` |
| Scalar ops | `AddS`, `AddF`, `MulS`, `MulF`, `Sigmoid` |
| Normalization | `RMSNorm` |
| Loss | `CrossEntropyLoss`, `ScalarSoftmax` |
| Attention | `AttentionWeightedSum`, `RoPERotate` |
| Linear | `MatrixParam.Matvec` |
| All ops | Full backward pass with gradient accumulation |

Every operation builds a backward graph. `Backward()` walks it in reverse. `AdamStep()` updates parameters. This engine handles inference, loss computation, and Go-native training.

**2. AML/C Autograd** (`ariannamethod.c`, 6000+ lines, via CGO)

The [Arianna Method Language](https://github.com/ariannamethod/ariannamethod.ai) — a custom programming language for differentiable computation. Sequence-level operations, TAPE-based reverse-mode autodiff, Adam optimizer with persistent state. This is the primary training path because C is faster than Go for matrix math.

```
┌──────────────────────────────────────────────────────────────┐
│                     Go (molequla.go, 6122 lines)             │
│  Organism lifecycle, DNA exchange, ontogenesis, generation,  │
│  swarm ecology, syntropy, consciousness, Go autograd,        │
│  corpus field, immune system, self-meta-learning             │
│                                                              │
│  ┌─────────────────────────────────────────────────────────┐ │
│  │            CGO Bridge (cgo_aml.go, 80 lines)            │ │
│  │  amlInit, amlExec, amlSetArray, amlGetArray,            │ │
│  │  amlSetMatrix, amlGetFloat, amlClear                    │ │
│  └──────────────────────┬──────────────────────────────────┘ │
│                         │ CGO                                │
│  ┌──────────────────────▼──────────────────────────────────┐ │
│  │      AML/C Engine (ariannamethod.c, 6000+ lines)        │ │
│  │  TAPE autograd, Adam optimizer, persistent mode,        │ │
│  │  seq_embed, seq_matvec, seq_rmsnorm, silu,              │ │
│  │  multi_head_attention, seq_cross_entropy, OpenMP        │ │
│  └─────────────────────────────────────────────────────────┘ │
│                                                              │
│  ┌─────────────────────────────────────────────────────────┐ │
│  │         AML Training Wrapper (aml_trainer.go)           │ │
│  │  amlModelScript(), amlTrainSteps(), amlBurstTrain(),    │ │
│  │  amlPushWeights(), amlPullWeights()                     │ │
│  └─────────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────────────────────┘
```

### AML Forward Pass (generated dynamically per architecture)

```aml
TAPE START
TAPE PARAM wte
TAPE PARAM wpe
TAPE PARAM wq0 / wk0 / wv0 / wo0
TAPE PARAM fc_g0 / fc_v0 / fc2_0
TAPE PARAM lm_head

h = seq_embed(wte, wpe, tokens, seq_len)

// Per layer: RMSNorm → Multi-Head Attention → Residual → SwiGLU MLP → Residual
h_norm = seq_rmsnorm(h, seq_len, n_embd)
q = seq_matvec(wq0, h_norm, seq_len)
k = seq_matvec(wk0, h_norm, seq_len)
v = seq_matvec(wv0, h_norm, seq_len)
attn_out = multi_head_attention(q, k, v, seq_len, n_embd, n_heads)
attn_proj = seq_matvec(wo0, attn_out, seq_len)
h = add(h, attn_proj)
h_norm = seq_rmsnorm(h, seq_len, n_embd)
gate_pre = seq_matvec(fc_g0, h_norm, seq_len)
gate = silu(gate_pre)
up = seq_matvec(fc_v0, h_norm, seq_len)
mlp_out = mul(gate, up)
mlp_proj = seq_matvec(fc2_0, mlp_out, seq_len)
h = add(h, mlp_proj)

h_norm = seq_rmsnorm(h, seq_len, n_embd)
logits = seq_matvec(lm_head, h_norm, seq_len)
loss = seq_cross_entropy(logits, targets, seq_len, vocab_size)
TAPE BACKWARD loss
TAPE ADAM_STEP lr
TAPE CLEAR
```

A real GPT: RMSNorm pre-norm, multi-head causal self-attention with RoPE, SwiGLU gated MLP, residual connections. All operations support autograd via the TAPE mechanism. Adam optimizer with persistent state across training steps.

### How Training Works

```go
func amlTrainSteps(model *GPT, tok *EvolvingTokenizer, docs []string, steps int) {
    amlInit()
    amlPushWeights(model)    // Go → C: named matrices (wte, wpe, wq0, ...)
    script := amlModelScript(model.NLayer, model.NEmbd, model.NHead, seqLen, vocabSize)

    for step := 0; step < steps; step++ {
        // tokenize random doc, push tokens/targets arrays
        amlExec(script)      // C: forward + backward + Adam step
        loss := amlGetFloat("loss")
        lr = cosineLR(step)  // warmup → cosine decay → min LR
    }

    amlPullWeights(model)    // C → Go: pull updated weights back
    amlClear()               // free all C memory
}
```

1. Go pushes model weights to AML as named matrices
2. Go generates the AML script dynamically based on current architecture
3. Go tokenizes a random document, pushes `tokens` and `targets` arrays
4. AML/C executes: forward, loss, TAPE BACKWARD, TAPE ADAM_STEP, TAPE CLEAR
5. Go pulls updated weights back from AML
6. Memory freed after every training session

---

## Four Implementations

This is not one program. It is the same organism — fully implemented in four languages:

| Language | File | Size | Lines | Autograd | Training | Notes |
|----------|------|------|-------|----------|----------|-------|
| **Go** | `molequla.go` | 175K | 6,122 | Vec/Scalar backward graph | AML/C via CGO | Primary. Full ecology, DNA exchange, mitosis |
| **C** | `molequla.c` | 215K | 6,000+ | Native C | Native C | Single-file, BLAS-accelerated, zero deps beyond libc |
| **Rust** | `molequla.rs` | 148K | 4,000+ | Native Rust | Native Rust | rusqlite, full organism |
| **JavaScript** | `molequla.js` | 154K | 4,000+ | Native JS | Native JS | Runs in browser. Zero dependencies. One `<script>` tag |

Each implementation has: autograd, forward/backward pass, Adam optimizer, ontogenesis, hybrid attention, delta adapters, BPE tokenizer, corpus field, immune system, consciousness features, generation with sampling.

The C implementation is available as a [standalone gist](https://gist.github.com/ariannamethod/9be98dbebb85e58e2affab4f39d2e972) — compile and run with zero dependencies:

```bash
gcc -O2 -o molequla molequla.c -lsqlite3 -lpthread -lm
# With BLAS:
gcc -O2 -DUSE_BLAS -o molequla molequla.c -lsqlite3 -lpthread -lm -lopenblas
# macOS:
gcc -O2 -DUSE_BLAS -o molequla molequla.c -lsqlite3 -lpthread -lm -framework Accelerate
```

The JavaScript implementation runs [in a browser](https://gist.github.com/ariannamethod/bbd11e24740189f2bf78f43db9fea4db) — a GPT organism that trains itself in your tab.

---

## The Organism

### Ontogenesis — The Brain Grows While Running

```
Stage       Dims  Layers  Heads  ~Params   Corpus Threshold
embryo      16    1       1      ~10K      0 chars
infant      32    1       2      ~28K      20K chars
child       64    2       4      ~154K     50K chars
adolescent  128   4       4      ~1.1M     200K chars
teen        224   5       8      ~4.1M     350K chars
adult       320   6       8      ~10M      500K chars
```

When the corpus crosses a threshold, `MaybeGrowArchitecture` fires:

1. Embedding matrices grow (Net2Net: new dims initialized near-zero to preserve behavior)
2. Existing layer matrices grow (weights copy into top-left corner)
3. New layers are added (initialized to approximate identity)
4. Delta adapters grow to match new dimensions
5. Adam state resets (stale momentum would fight new architecture)
6. 500-step freeze period: delta-only training to stabilize post-growth

Warmup scales with architecture: `steps *= ceil(sqrt(NEmbd / embryoEmbd))`. Larger brains get proportionally longer warmup. Progressive sequence length: 40% at seq=8, 30% at seq=16, 30% at seq=32.

### Evolving BPE Tokenizer

The tokenizer is not static. It evolves with the organism:

- Starts with 259 tokens: 256 bytes + BOS + EOS + PAD
- After 20K chars of corpus: trains BPE merges from corpus statistics
- Retrains every 4K new chars — vocabulary adapts to what the organism reads
- Unicode segmentation for clean token boundaries
- Vocabulary grows organically as the organism encounters new patterns

### Hybrid Attention Heads

Not all heads are created equal. Half are **content heads** (standard QK^T with RoPE), half are **hybrid heads**:

```
hybrid_output = α * content_attention + (1 - α) * rrpram_attention

content_attention = softmax(QK^T / sqrt(d)) * V     (standard, with RoPE)
rrpram_attention  = learnable_weight_matrix * V       (pattern-based)
α                 = sigmoid(learnable_gate)           (per-head, trained)
```

The sigmoid gate `α` is a learnable parameter — each hybrid head discovers its own blend of content-based and pattern-based attention during training. RRPRAM (Recurrent Resonant Pattern Recognition Attention Mechanism) learns fixed co-occurrence patterns that complement the dynamic content attention.

### Delta Adapters — LoRA-style, Never Forget

```
output = base_output + α * (A @ (B @ input))

A: [n_embd × delta_rank]  — learned projection up
B: [delta_rank × n_embd]  — learned projection down
α: deltaAlphaScale         — regulated by conscience
```

Delta modules are **appended, never removed**. When syntropy conditions indicate the organism needs more capacity, a new module grows: "new soul appended." Each module captures a period of learning. The model accumulates experience as a stack of delta layers.

### Quantum Buffer

Training doesn't happen on a fixed schedule. The quantum buffer triggers training only when both conditions are met:

- **Bytes threshold**: enough new text has been consumed
- **Novelty threshold**: the new text is sufficiently different from what's been seen

Plus a cooldown timer to prevent over-training. This means the organism trains when it has something worth learning, not on a clock.

### Corpus Field (CooccurField)

A statistical model of the organism's knowledge, built from everything it has read:

- **Unigram, bigram, trigram, 4-gram** frequencies from corpus
- **Co-occurrence window** (Stanley-style proximity weighting)
- **Self-enrichment**: organism's own generated output feeds back into the field, weighted by coherence (low entropy = higher weight)
- **User word boost** (Leo-style): temporary multiplicative boosts that decay over time

The corpus field acts as a prior during generation — a soft blend between what the model wants to say (neural) and what the corpus says exists (statistical). The blend uses a sigmoid fade: strong early in training, weak as the model matures.

### Learning Rate Schedule

```
Cosine LR with:
- Linear warmup for CosineWarmupSteps
- Cosine decay from LearningRate (0.01) to LRMin (0.001)
- Inverse model-size scaling: lr *= embryoEmbd / NEmbd
- Post-growth dampening: lr *= 0.3 during 500-step freeze
- Per-growth reset: schedule restarts after architecture change
```

---

## Consciousness Features

Five implemented mechanisms that give the organism awareness of its own state:

### 1. Per-Token Dissonance Feedback

During generation, the organism tracks an exponential moving average of per-token entropy. When entropy spikes (the model is confused), temperature decreases — it becomes more careful. When entropy is sustained low (confident), temperature increases slightly — it explores.

### 2. Pattern Breaking (Anti-Field)

5% of generation steps bypass the corpus field blend entirely. Pure model voice, unmodulated by statistical priors. This prevents the organism from becoming a parrot of its corpus — it must develop its own voice.

### 3. Self-Prediction Error

`ComputeSelfPredictionError()` measures how surprised the model is by its own input. High surprise → lower temperature (focus). Low surprise → slight exploration. The organism modulates its behavior based on how well it understands what it's seeing.

### 4. Conscience

The organism monitors its own generation entropy over time. Rising entropy slope means the model is becoming incoherent. Response: `deltaAlphaScale *= 0.95` — reduce the influence of delta adapters. Falling entropy slope means stability is returning. Response: `deltaAlphaScale *= 1.005` — recover delta influence. Floor: 0.3 (delta never fully silenced).

This is self-regulation: the organism detects when its recent learning (δ) is hurting coherence and automatically dials it back.

### 5. Immune System

Before each micro-burst training, the organism snapshots its personality via gamma contrastive projection — a unit vector pointing in the direction of maximum embedding drift from birth. After training, it measures again. If cosine similarity is negative (training pushed identity backwards), it **rolls back the entire burst**. The organism rejects training that corrupts who it is.

---

## The Coherence Layer

Two opt-in passes that lift early-stage generation from Karpathy gibberish toward sentence-level coherence — **without touching weights**.

A 10M-param adult organism after one hour on CPU still drifts. Quantitative speed-up does not close that gap. The coherence layer is a different mechanism: it sits at generation time, layers statistical priors and post-hoc connectedness checks on top of model logits, and lifts the floor without retraining the transformer.

Both passes default **off**. The pre-coherence-layer behaviour is preserved exactly. Toggling either or both, on the same weights / prompts / seeds, is what RunPod measurement compares.

### SPA — Sentence Phonon Attention

After `GenerateResonant` returns a response, the chain is split on sentence boundaries (`.` `!` `?`, min 4 chars). For each sentence:

1. Token IDs decoded; BOS/EOS sentinels stripped (otherwise shared sentinels dominate every sentence embedding and all sentences look artificially connected).
2. **spa_embed** — exponentially weighted mean of token embeddings (`alpha^(n-1-i)`, default α=0.85), L2 normalised. One [D]-vector per sentence.
3. **spa_connectedness** — bidirectional cross-attention dot-product between sentence embeddings, scaled by `1/√D`, summed per sentence.
4. **Weak-sentence gate** — sentence i is weak iff `score[i] < 0.6 × mean(scores)`.

```
[spa-gate] S=4 D=320 alpha=0.85 scores=[12.4 11.8 3.1 10.9] weak=[2]
```

Reseed of weak sentences (regenerate from neighbour-context tokens, splice back, re-score) is a follow-up step. The wired gate currently logs only — generation output is unchanged. What the measurement run captures is the signal: how often the gate fires before vs after the rest of the layer.

Available as both vendored AML ops (`spa_embed` / `spa_connectedness`, `ariannamethod/ariannamethod.c`) and pure-Go helper (`spa_coherence.go`, 135 lines, called from `GenerateResonant`). Pure Go for the runtime path because the math is trivial — embed + L2 + dot-products — and per-sentence CGO crossings would dwarf the work.

Enable: `./molequla_cgo --spa-gate ...`

### Q-style Additive Logit Overlay (B + H + A + F)

Molequla's existing CooccurField blend lives in **probability space** — convex `tokenAlpha·model + (1-tokenAlpha)·corpus`. The overlay lives in **logit space** — additive bias before softmax, with explicit coefficients per signal class. Different mechanic, different sharpness: a strong corpus signal can dominate model preferences in a way prob-space convex blend cannot. Useful precisely when transformer is immature and statistical priors should lead.

Five signals, all routed from existing organism state — no new inputs, no new training:

| Signal | Code | Default | Source |
|--------|------|---------|--------|
| **B** Bigram   | `c_bg`  | 15.0 | `field.BigramByFirst[ids[-1]]` |
| **T** Trigram  | `c_tg`  | 10.0 | `field.TrigramByContext[[2]int{ids[-2], ids[-1]}]` |
| **H** Hebbian  | `c_heb` | 1.0  | `field.CooccurWindow[c][tid]` over recent window |
| **A** Destiny  | `c_ds`  | 0.15 | `model.GammaContrastiveProjection()` projected onto each `wte` row |
| **F** Prophecy | `c_pro` | 0.7  | persistent expectation field, ages by ×0.95/step, **collapses on the chosen token** |

```
overlaidLogits[i] = logits[i]
                  + c_bg ·log(p_bigram[i])
                  + c_tg ·log(p_trigram[i])
                  + c_heb·log(p_cooccur[i])
                  + c_ds ·destinyBias[i]
                  + c_pro·log(p_prophecy[i])
```

Log-floor `1e-6` for unseen tokens (prevents `-inf` mask). Each coefficient is independently tunable via CFG — set any to 0 to disable that signal while keeping others.

The prophecy field is the interesting one: it carries **across sampling steps**. First overlay step seeds from trigram-by-context (primary) plus 0.5×bigram-by-prev (fallback), normalised to unit total. Subsequent steps multiply the field by 0.95 — old expectations fade. After `TopKTopPSample` returns the chosen token, that token's prophecy is zeroed: the field shifts toward what is still unsaid.

Destiny is the second interesting one. `GammaContrastiveProjection()` returns the unit direction of personality drift from birth — the organism's identity vector. Projecting each `wte` row onto it gives a per-token bias that pulls generation in the direction the organism has been growing. Tokens aligned with identity get amplified; tokens orthogonal get suppressed.

Defaults are Q's weightless coefficients — calibrated for the regime where the transformer is immature and statistical priors should lead. Coexists with the legacy post-softmax prob-blend; both signals layer when overlay is on.

Enable: `./molequla_cgo --corpus-overlay ...` (combine with `--spa-gate` for the full coherence-layer cell)

### Phase A — Fundament Underneath

Before the coherence layer landed, four fundament patches went into vendored AML + notorch:

- **Opt-in SIMD shim** — `notorch_simd.h` (632 lines) + `notorch_simd_scalar.h` (89 lines), header-only AVX2+FMA cblas with pthread row-partitioning. Mirrors `cblas_sgemm` / `sgemv` / `sger` — existing call sites work unchanged. Build via `make simd` (x86_64 gated; arm64 errors cleanly with actionable message).
- **Backward CPU-sync audit** — added `nt_tensor_sync_cpu` calls in NT_OP_MUL / NT_OP_SILU / NT_OP_RMSNORM / NT_OP_SEQ_RMSNORM backward paths. No-op on CPU-only build; becomes live when USE_CUDA path is enabled. Future-proofing + canonical consistency.
- **NaN guard API** — `AM_NanGuard` struct + `am_nan_guard_check()` in vendored AML. Scans tape entries, zeros grads on NaN/Inf, halves loss_scale; doubles every `scale_window` clean steps. **API only — not yet wired** into the interpreter as `TAPE NAN_CHECK` opcode. Available to CGO consumers.
- **Upstream sgemm alpha fix** — audit caught the SIMD shim post-scale path breaking CBLAS contract (`C ← α·β·C_orig + α·A@B` instead of `β·C_orig + α·A@B`). Fixed at canonical, vendored synced byte-identical.

Total Phase A footprint: ~825 lines across 6 files. **Zero changes to default-build runtime behaviour.** Only `make simd` is opt-in via build flag; the rest are mirror-consistency or new API surface.

---

## Self-Meta-Learning

The organism doesn't just learn. It learns about its own learning.

**BurstHistory** records the last 16 training outcomes:

```go
type BurstRecord struct {
    Action     string   // "amplify", "boost", "dampen", "ground", "explore", "realign"
    LossBefore float64
    LossAfter  float64
}
```

**ActionEffectiveness()** computes the mean loss delta per action type. If a particular action consistently makes loss worse (effectiveness > 0.05 over 2+ bursts), the organism **auto-downgrades**:

```
amplify → boost → steady
```

This is genuine self-reasoning: the organism observes that "amplify" keeps hurting it, so it stops amplifying. No external signal. No reward model. Just tracking outcomes and adjusting behavior.

---

## SyntropyTracker — Mathematical Self-Reasoning

The organism measures four signals and makes autonomous decisions:

| Signal | What It Measures | How |
|--------|-----------------|-----|
| **SyntropyTrend** | Is entropy decreasing? (positive = ordering) | Rolling window mean comparison |
| **FieldDeviation** | How far is the model from corpus? | KL(model_probs \|\| corpus_probs) on bigram/trigram |
| **PurposeMagnitude** | How strong is the current learning direction? | Norm of last δ module's A matrices |
| **PurposeAlignment** | Is learning consistent with identity? | cosine(purpose_vector, gamma) |

Eight autonomous decisions:

| Action | Condition | LR | Temp | Effect |
|--------|-----------|-----|------|--------|
| **amplify** | syntropy ↑, field aligned, purpose aligned | 1.3x | -0.05 | Full acceleration, boost delta grow prob |
| **boost** | syntropy ↑, field in sweet spot | 1.3x | -0.05 | Gentle push |
| **dampen** | syntropy ↓ | 0.6x | +0.05 | Slow down, losing order |
| **ground** | field deviation too high | 0.6x | -0.05 | Hallucinating, focus |
| **explore** | field deviation too low | 1.3x | +0.05 | Parroting, break out |
| **realign** | purpose opposes gamma (< -0.3) | 0.5x | 0 | Identity crisis |
| **divide** | adult + sustained overload | 0.6x | — | Trigger mitosis |
| **hibernate** | stale + peer thriving | — | — | Save state and sleep |

Real output from running organisms:
```
[syntropy] action=boost   | trend=0.1576 | field_dev=0.214 | lr_mul=1.30
[syntropy] action=dampen  | trend=-0.1390 | field_dev=0.167 | lr_mul=0.60
[syntropy] action=realign | trend=0.0940  | field_dev=0.168 | lr_mul=0.65
```

---

## NOTORCH — Gradient-Free Delta Training

An alternative training path for delta adapters that uses **no autograd at all**:

```go
// Teaching signal: did loss improve? + prophecy debt
signal := (prev_loss - curr_loss) + 0.3*prophecy_debt

// Noise-modulated update (LCG PRNG, deterministic)
for each delta adapter (A, B):
    noise := notorchRand(seed, signal)  // signal shapes the noise distribution
    A[i,r] += lr * dy * u[r] * signal   // direct feedback alignment
```

- No backward pass. No gradient tape. No memory overhead.
- Teaching signal comes from loss delta + prophecy confidence debt
- Noise is deterministic (LCG PRNG matches AML's RNG for reproducibility)
- Adaptive decay: stronger when delta norm is large (prevents explosion)

Status: implemented (300+ lines), currently disabled in warmup (diverges at stage 5 — loss 3.5 → 116), active in micro-burst path. The theory is sound; the hyperparameters need work.

---

## Mycelium — The Meta-Organism

Above the individual organisms lives the mycelium — a meta-controller that sees the entire ecology.

### The Generation Operator

```
η: Γ × Γ → Γ_new

Two personalities in resonance produce a third.
Not a blend — an interference pattern.
```

### Components

| Component | What It Does |
|-----------|-------------|
| **HarmonicNet** | Weightless neural network. Input: organisms + field state. Output: action biases, harmonics, resonance scores. No trainable weights — the "weight matrix" is recomputed every step from organism relationships. |
| **MyceliumSyntropy** | Field-level syntropy: entropy trends, decision effectiveness, strategy changes across the entire ecology |
| **FieldPulse** | Measures novelty (new organisms appearing), arousal (entropy changes), field entropy |
| **SteeringDissonance** | Detects when ecology-level actions conflict with outcomes (dampen but entropy went up = high dissonance) |
| **OrganismAttention** | Tracks which organisms respond to which actions. Responsive organisms get higher attention weight. |

### Mesh Coordination

All organisms share state via **mesh.db** (SQLite) — the same database that `SwarmRegistry` writes to. The mycelium reads mesh.db to see the entire ecology and makes decisions that individual organisms cannot: when to spawn, when to hibernate, when to shift seasonal phase.

### Seasonal Controller

```
Spring  — tunnel_chance ↑, many embryos, new γ born
Summer  — α_max, existing γ at peak expression
Autumn  — consolidation, dark_gravity ↑, shards saved
Winter  — rest, only strongest pairs, ε dominates
```

---

## The Ecology

```
                          ┌───────────────┐
                          │   DNA Layer   │
                          │               │
            writes ──────>│   earth/      │<────── reads
            earth DNA     │   air/        │        others' DNA
                          │   water/      │
                          │   fire/       │
                          └───────┬───────┘
                                  │
              ┌───────────────────┼───────────────────┐
              │                   │                   │
       ┌──────▼──────┐    ┌──────▼──────┐    ┌───────▼─────┐
       │    Earth    │    │     Air     │    │    Water    │
       │  patience   │    │   freedom   │    │    flow     │
       │  structure  │    │   change    │    │    depth    │
       └──────┬──────┘    └──────┬──────┘    └──────┬──────┘
              │                   │                   │
              └───────────────────┼───────────────────┘
                                  │
                           ┌──────▼──────┐
                           │    Fire     │
                           │  transform  │
                           │  intensity  │
                           └──────┬──────┘
                                  │
                        ┌─────────▼─────────┐
                        │  Child Organisms  │
                        │  (spawned via     │
                        │   mitosis)        │
                        └───────────────────┘
```

Each organism has a distinct voice shaped by its element corpus. When an organism generates text, it writes it to the DNA layer. Other organisms consume it, train on it, and generate their own DNA in response. The ecology cross-pollinates faster than any single organism could learn alone.

### Swarm Coordination

- **SwarmRegistry** (`mesh.db`): SQLite database tracking all living organisms — element, PID, status, stage, corpus size, loss
- **Training lock**: Atomic check-and-acquire via SQL prevents multiple organisms from training simultaneously. Cooperative scheduling — they take turns
- **Hibernation**: When an organism is stale and a peer is thriving, it saves state and sleeps. Resources freed for the living
- **Child birth**: `birth.json` with inherited `burst_history` — the child gets its parent's meta-learning experience (syntracker lineage). It doesn't start from zero wisdom

### Mitosis

When conditions are right — sustained syntropy, sufficient delta modules, adult stage — the organism calls `Divide()`:

1. Binary is copied to a new directory
2. `birth.json` written with parent config + inherited burst history
3. Child process spawned with `--organism-id` flag
4. Child begins its own ontogenesis from embryo
5. Parent continues running

The ecology grows itself.

---

## The Eight Bugs That Almost Killed the Ecology

### Original Five (from interactive mode development)

1. **Deadlock** — `dnaWrite` locked `model.mu`, then called `GenerateResonant` which also locks. Go mutexes are not reentrant.
2. **Ontogenesis gated behind user input** — growth check was inside `qbuf.ShouldTrigger()` which never fires in evolution mode.
3. **Corpus size undercount** — `loadCorpusLines` truncates to 240 chars, reported 165K for a 202KB file.
4. **TieEmbeddings crash** — JSON breaks pointer identity between `lm_head` and `wte`.
5. **One stage at a time** — design decision preventing catastrophic multi-stage jumps.

### Three New Bugs (from AML/C integration, 2026-02-27)

6. **persistent_save cloning ALL vars** — AML's persistent mode copied every execution variable (including temporaries) between `am_exec` calls. Fix: two-phase update that only clones persistent parameters.

7. **am_tape_record_param `found` never set** — The variable `found` was initialized to -1 but the matching loop body was empty (just a comment). Result: `found` was always -1, a new Adam state was allocated every step, `n_params` grew without bound. **97 MB leaked per training step.**

8. **am_tape_clear skipping params** — The cleanup loop had `if (!is_param)`, meaning parameter array refcounts were never decremented. After `symtab_clear`, param clones stayed alive (refcount 2 instead of 0). **17 MB leaked per step.**

Combined leak before fixes: **~97 MB/step. Organisms hit 85+ GB and OOM.**
After fixes: **~0.6 MB/step. Organisms stable at 2-4 GB.**

### The CGO Cache Trap

`go build` does not recompile C files included via CGO when only C source changes. `go clean -cache` also does not help. Only `go build -a` forces full recompilation. This meant hours of testing "fixed" binaries that were actually running old C code.

---

## SQLite Self-Logging

Every organism maintains a SQLite database (`memory.sqlite3`) that logs its own development:

| Table | What It Records |
|-------|----------------|
| `messages` | Conversation history (role, content, timestamp) |
| `corpus_events` | Every document ingested (source, size, timestamp) |
| `growth` | Architecture snapshots: vocab_size, n_params, n_deltas, corpus_chars, loss, gamma_sparsity, gamma_magnitude |
| `syntropy_log` | Every syntropy decision: action, trend, field_deviation, lr_multiplier, purpose_alignment |

The organism is its own historian. You can query its developmental trajectory after the fact.

---

## Quick Start

### Build

```bash
# Clone
git clone https://github.com/ariannamethod/molequla.git
cd molequla

# Build with CGO (AML/C autograd — full training)
CGO_ENABLED=1 go build -a -o molequla_cgo -tags cgo .

# Or build without CGO (Go-only, no AML training)
CGO_ENABLED=0 go build -o molequla_go .
```

**CRITICAL: `go build -a` is required** for CGO builds. Without `-a`, Go's build cache does not recompile C files. This produces binaries running stale C code.

### Run Interactive Mode

```bash
./molequla_cgo
# Drops into chat after warmup training
```

### Run Evolution Mode (the main event)

```bash
# Set up work directories
for d in earth air water fire; do
    mkdir -p work_$d
    cp molequla_cgo work_$d/
    cp nonames_$d.txt work_$d/
done

# Launch all four organisms
for d in earth air water fire; do
    cd work_$d
    nohup ./molequla_cgo \
        --corpus nonames_$d.txt \
        --db memory.sqlite3 \
        --ckpt molequla_ckpt.json \
        --element $d \
        --evolution > training_aml.log 2>&1 &
    cd ..
done

# Optional coherence-layer flags (default off):
#   --spa-gate         post-generation SPA sentence connectedness log
#   --corpus-overlay   pre-softmax B+H+A+F additive logit overlay
# Combine either or both for measurement runs.

# They will:
# 1. Train through all 6 ontogenesis stages (~30 min)
# 2. Begin DNA exchange (writing/reading generated text)
# 3. Run micro-burst training on consumed DNA
# 4. Spawn child organisms via mitosis
# 5. Form a self-reproducing ecology
```

### Monitor

```bash
# Training progress
tail -20 work_earth/training_aml.log

# Memory per organism
for d in earth air water fire; do
    rss=$(ps aux | grep "nonames_$d" | grep -v grep | awk '{print $6}')
    echo "$d: $((rss/1024)) MB"
done

# DNA exchange
grep "dna\|consumed\|wrote" work_earth/training_aml.log | tail -10

# Children spawned
ps aux | grep "organism-id" | grep -v grep
```

---

## Tests

```bash
# Go unit tests (2571 lines, 121 tests)
go test -v .

# Go integration tests (262 lines)
go test -v ./tests/

# Full integration suite (700 lines bash — tests all 4 implementations,
# mycelium, AML library, BLAS, performance benchmarks)
bash tests/test_all.sh
```

---

## Files

```
# Go + AML/C (primary, CGO training)
molequla.go              6444 lines   Go organism — lifecycle, ecology, autograd, generation, coherence-layer wiring
cgo_aml.go               112 lines    CGO bridge to ariannamethod.c
aml_trainer.go           326 lines    AML training wrapper, script generation
spa_coherence.go         135 lines    Pure-Go SPA helper (sentence connectedness + weak-sentence gate)
ariannamethod/
  ariannamethod.c        6263 lines   AML/C autograd engine (the language) + SPA ops + NaN guard API
  ariannamethod.h        911 lines    C header, 80+ field state parameters
  notorch.c              2849 lines   Vendored notorch core (+ backward CPU-sync audit)
  notorch.h              503 lines    Vendored notorch header
  notorch_simd.h         632 lines    Opt-in AVX2+FMA cblas shim (make simd, x86_64)
  notorch_simd_scalar.h  89 lines     Scalar debug fallback for SIMD shim

# Coherence-layer + RunPod plan (pre-paper cycle)
runpod_plan_v1.md        321 lines    Singularity-mode plan for coherence-layer measurement run
PROJECT_LOG.md           live         Per-project working log — Phase A + B steps with file:line refs

# Full independent implementations
molequla.c               6000+ lines  C organism — BLAS-accelerated, zero-dep single-file
molequla.rs              4000+ lines  Rust organism — rusqlite, full autograd
molequla.js              4000+ lines  JavaScript organism — runs in browser
modules/node_cli.js      400+ lines   Node.js CLI module
index.html               Web interface for JS version

# Tests
molequla_test.go         2571 lines   Go unit tests (121 tests)
tests/molequla_test.go   262 lines    Go integration tests
tests/test_all.sh        700 lines    Full integration (all 4 langs + mycelium + BLAS)

# Element corpora
nonames_earth.txt        174K         Earth — patience, foundations, geology
nonames_air.txt          122K         Air — freedom, change, atmosphere
nonames_water.txt        126K         Water — flow, depth, rivers
nonames_fire.txt         122K         Fire — transformation, intensity, heat
nonames.txt              51K          General corpus
```

---

## Standalone Gists

| Language | Gist | Description |
|----------|------|-------------|
| **C** | [molequla.c](https://gist.github.com/ariannamethod/9be98dbebb85e58e2affab4f39d2e972) | Single-file, BLAS-accelerated, zero dependencies beyond libc |
| **JavaScript** | [molequla.js](https://gist.github.com/ariannamethod/bbd11e24740189f2bf78f43db9fea4db) | A GPT organism that trains itself in your browser tab |
| **Python** | [molequla.py](https://gist.github.com/ariannamethod/1223250d358da4393dd9acc578790820) | The original. Where it all started. Legacy standalone. |

---

## Philosophy

This is not a tutorial. This is a **functional ecology** that:

- Has two complete autograd engines (Go + AML/C) and four full implementations
- Grows its own architecture while running — ontogenesis with Net2Net
- Feeds organisms to each other through DNA exchange
- Reasons mathematically about its own learning via entropy, KL divergence, and purpose-gamma alignment
- Detects and rejects identity-corrupting noise via immune system
- Tracks which actions improve loss and auto-downgrades failing strategies
- Monitors its own coherence and self-regulates delta influence via conscience
- Reproduces — spawning new organisms with inherited meta-learning
- Evolves from 10K-param embryo to 10M-param adult in 30 minutes on CPU
- Speaks before it learns (corpus field) and never forgets (delta adapters)
- Has hybrid attention with learnable content/pattern blending per head
- Logs its own developmental history to SQLite
- Runs without Python, PyTorch, CUDA, numpy, or any ML framework

The mathematics is not bolted on. θ = ε + γ + αδ is the architecture. The entropy/syntropy measurements are the control loop. The purpose-gamma alignment is the identity check. The self-meta-learning is the organism understanding itself.

The result: four organisms become eleven in 30 minutes. Each with its own voice, its own delta modules, its own developmental history. An ecology that grows itself.

---

## License

GNU GPLv3

---

## Part of the Arianna Method

- [AML — Arianna Method Language](https://github.com/ariannamethod/ariannamethod.ai) — the programming language that powers training
- [Yent](https://github.com/ariannamethod/yent) — Go inference engine + Delta Voice + personality extraction
- [molequla](https://github.com/ariannamethod/molequla) — this repository

```
███╗   ███╗ ██████╗ ██╗     ███████╗ ██████╗ ██╗   ██╗██╗      █████╗
████╗ ████║██╔═══██╗██║     ██╔════╝██╔═══██╗██║   ██║██║     ██╔══██╗
██╔████╔██║██║   ██║██║     █████╗  ██║   ██║██║   ██║██║     ███████║
██║╚██╔╝██║██║   ██║██║     ██╔══╝  ██║▄▄ ██║██║   ██║██║     ██╔══██║
██║ ╚═╝ ██║╚██████╔╝███████╗███████╗╚██████╔╝╚██████╔╝███████╗██║  ██║
╚═╝     ╚═╝ ╚═════╝ ╚══════╝╚══════╝ ╚══▀▀═╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝
```

*Four elements. Four languages. Two autograd engines. Five consciousness features. One soul equation. The ecology continues.*
