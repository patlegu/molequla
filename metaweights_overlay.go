package main

import "math"

// ═══════════════════════════════════════════════════════════════════════════════
// Q-style metaweights overlay — raw-probability, dynamic-gate, coherence-without-training
//
// Ported from ~/arianna/q/postgpt_q.c:1305-1395, the reference implementation
// of «coherence emerges from corpus statistics» pattern. The mechanism mixes
// the model's untrained logits with five statistical signals (Hebbian / prophecy
// / destiny / bigram / trigram) added as RAW probability values, not log-probs.
//
// Coefficients are dynamic — magnitude-detector picks weightless or trained
// bundle based on average |logit|. While the transformer's weights are weak,
// the overlay dominates and produces coherent output. As weights strengthen,
// overlay fades, ceding to model voice. Q's auto-curriculum.
//
// This is the load-bearing fix for the regression where Phase B used
// math.Log(prob) instead of raw prob and multiplied by coefficients 15/10,
// which created -30 to -200 logit penalties that swamped untrained model
// logits in the wrong direction — every previous run output gibberish.
// ═══════════════════════════════════════════════════════════════════════════════

// MetaCoeffs holds the five Dario-field overlay coefficients for one regime.
type MetaCoeffs struct {
	Heb, Pro, Ds, Bg, Tg float64
}

// Reference values verbatim from postgpt_q.c:1358-1359.
var (
	metaCoeffsWeightless = MetaCoeffs{Heb: 1.0, Pro: 0.7, Ds: 0.15, Bg: 15.0, Tg: 10.0}
	metaCoeffsTrained    = MetaCoeffs{Heb: 0.6, Pro: 0.4, Ds: 0.3, Bg: 5.0, Tg: 3.0}
)

// Q's transformer-magnitude gate threshold (postgpt_q.c:1356).
const metaTFGateThreshold = 0.1

// MetaweightsOverlay applies the Q-style dynamic logit overlay in place on
// `logits` (model's raw pre-temperature logits over the vocabulary).
//
//   - `ids` — the running token context for this generation step (last token
//     is the immediate predecessor of the position we're predicting).
//   - `field` — the organism's CooccurField (must not be nil).
//   - `model` — used for destiny term (purpose direction projected onto wte).
//   - `prophecyField` — persistent expectation field; nil means «not yet seeded».
//     The caller owns this slice and re-passes it across steps for age + collapse.
//   - Returns the modified logits slice (same backing array — caller can use
//     it directly for softmax) and the (possibly seeded) prophecy field.
//
// Mechanism (mirrors postgpt_q.c:1354-1395 step-for-step):
//  1. Compute tmag = mean(|logits|). If tmag > 0.1 → trained coeffs, else weightless.
//  2. Look up bigram[prev][i], trigram[prev2,prev1][i], normalised per-context.
//  3. Aggregate Hebbian window counts over recent CooccurWindowSize ctx tokens.
//  4. Seed/age prophecy field, take prophecy[i] as bias.
//  5. Project wte rows onto gamma direction (destiny[i] = cosine).
//  6. logits[i] += c_heb*heb[i] + c_pro*pro[i] + c_ds*ds[i] + c_bg*bg[i] + c_tg*tg[i].
//  7. Unigram damping: very rare tokens get -2.0 penalty, common tokens get
//     linear suppression — mirror postgpt_q.c:1393-1394.
func MetaweightsOverlay(
	logits []float64,
	ids []int,
	field *CooccurField,
	model *GPT,
	prophecyField []float64,
) ([]float64, []float64) {
	V := len(logits)
	if V == 0 || field == nil || len(ids) < 1 {
		return logits, prophecyField
	}

	// 1. Magnitude gate — pick coefficient bundle.
	var tmag float64
	for _, v := range logits {
		if v < 0 {
			tmag -= v
		} else {
			tmag += v
		}
	}
	tmag /= float64(V)
	coeffs := metaCoeffsWeightless
	if tmag > metaTFGateThreshold {
		coeffs = metaCoeffsTrained
	}

	// 2. Bigram + trigram per-context — normalised probabilities.
	bigramProb := make([]float64, V)
	trigramProb := make([]float64, V)
	field.mu.RLock()
	prev := ids[len(ids)-1]
	if ctx, ok := field.BigramByFirst[prev]; ok {
		var total float64
		for _, v := range ctx {
			total += v
		}
		if total > 0 {
			for tid, v := range ctx {
				if tid < V {
					bigramProb[tid] = v / total
				}
			}
		}
	}
	if len(ids) >= 2 {
		a, b := ids[len(ids)-2], ids[len(ids)-1]
		if ctx, ok := field.TrigramByContext[[2]int{a, b}]; ok {
			var total float64
			for _, v := range ctx {
				total += v
			}
			if total > 0 {
				for tid, v := range ctx {
					if tid < V {
						trigramProb[tid] = v / total
					}
				}
			}
		}
	}

	// 3. Hebbian — window-walked co-occurrence, max-normalised (postgpt_q.c:196).
	hebbianStrength := make([]float64, V)
	windowSize := CFG.CooccurWindowSize
	if windowSize <= 0 || windowSize > len(ids) {
		windowSize = len(ids)
	}
	var hebMax float64
	for j := len(ids) - windowSize; j < len(ids); j++ {
		if neighbors, ok := field.CooccurWindow[ids[j]]; ok {
			for tid, cnt := range neighbors {
				if tid < V {
					hebbianStrength[tid] += cnt
					if hebbianStrength[tid] > hebMax {
						hebMax = hebbianStrength[tid]
					}
				}
			}
		}
	}
	if hebMax > 0 {
		for i := range hebbianStrength {
			hebbianStrength[i] /= hebMax
		}
	}

	// 4. Unigram for damping pass.
	unigramProb := make([]float64, V)
	var uniTotal float64
	for _, v := range field.Unigram {
		uniTotal += v
	}
	if uniTotal > 0 {
		for tid, v := range field.Unigram {
			if tid < V {
				unigramProb[tid] = v / uniTotal
			}
		}
	}
	field.mu.RUnlock()

	// 5. Prophecy — seed once, then age per step.
	prophecyProb := make([]float64, V)
	if prophecyField == nil {
		prophecyField = make([]float64, V)
		// Seed from trigram (primary) + half-weight bigram fallback.
		for i := 0; i < V; i++ {
			prophecyField[i] = trigramProb[i] + 0.5*bigramProb[i]
		}
		// Normalise.
		var pt float64
		for _, v := range prophecyField {
			pt += v
		}
		if pt > 0 {
			for i := range prophecyField {
				prophecyField[i] /= pt
			}
		} else {
			prophecyField = nil
		}
	} else {
		decay := CFG.MetaProphecyDecay
		if decay <= 0 || decay > 1 {
			decay = 0.95
		}
		for i := range prophecyField {
			prophecyField[i] *= decay
		}
	}
	if prophecyField != nil {
		var pt float64
		for _, v := range prophecyField {
			pt += v
		}
		if pt > 0 {
			for i := 0; i < V && i < len(prophecyField); i++ {
				prophecyProb[i] = prophecyField[i] / pt
			}
		}
	}

	// 6. Destiny — cosine(wte[i], gammaDir). Lazy compute once would be nicer
	// but Q recomputes per step from current destiny vector. Mirror that.
	destinyScore := make([]float64, V)
	if model != nil {
		if gammaDir, mag := model.GammaContrastiveProjection(); mag > 0 && len(gammaDir) > 0 {
			if wte := model.Base["wte"]; wte != nil {
				D := wte.Nin
				if D > 0 && D <= len(gammaDir) {
					for v := 0; v < V && v < wte.Nout; v++ {
						row := wte.Rows[v].Data
						var dot, en float64
						for d := 0; d < D; d++ {
							dot += gammaDir[d] * row[d]
							en += row[d] * row[d]
						}
						en = math.Sqrt(en + 1e-10)
						if en > 1e-8 {
							destinyScore[v] = dot / en
						}
					}
				}
			}
		}
	}

	// 7. The actual overlay — Q's line 1383, raw values:
	//    raw[i] += c_heb*heb[i] + c_pro*pro[i] + c_ds*ds + c_bg*bg + c_tg*tg
	for i := 0; i < V; i++ {
		logits[i] += coeffs.Heb*hebbianStrength[i] +
			coeffs.Pro*prophecyProb[i] +
			coeffs.Ds*destinyScore[i] +
			coeffs.Bg*bigramProb[i] +
			coeffs.Tg*trigramProb[i]

		// Unigram damping (postgpt_q.c:1393-1394) — suppress noise tokens.
		if unigramProb[i] < 1e-6 {
			logits[i] -= 2.0
		} else if unigramProb[i] > 0.01 {
			logits[i] -= 0.3 * (unigramProb[i] - 0.01) * 100.0
		}
	}

	return logits, prophecyField
}

// MetaweightsOverlayCollapse zeroes the prophecy slot for a token that was
// just sampled — Q's «collapse on fulfilment» pattern. Called from the
// generation loop after sampling each token. nil-safe.
func MetaweightsOverlayCollapse(prophecyField []float64, sampledID int) {
	if prophecyField != nil && sampledID >= 0 && sampledID < len(prophecyField) {
		prophecyField[sampledID] = 0
	}
}

// MetaweightsRepetitionPenalty applies Q's age-graded repetition penalty +
// bigram-blocking pass on raw logits AFTER the metaweights overlay.
// Mirror of postgpt_q.c:1399-1408. Multiplicative penalties.
//
//   - Recent tokens (last 20 of ctx): logits[t] *= 0.335 + 0.035*age
//     where age=1 for most recent, age=20 for oldest in window.
//     Recent gets stronger penalty (0.335 ≈ 67% damping), old weaker
//     (0.665 ≈ 33% damping). This breaks lock-in like «Sppellllllll».
//   - Bigram blocking (cl >= 2): for every position where ctx[i]==ctx[cl-2],
//     penalise ctx[i+1] by 0.2. Prevents re-emitting the same bigram pair
//     just seen — kills two-token repetition cycles.
//
// Logits are mutated in place.
func MetaweightsRepetitionPenalty(logits []float64, ids []int) {
	V := len(logits)
	cl := len(ids)
	if V == 0 || cl == 0 {
		return
	}
	// Age-graded penalty over last 20 ctx tokens.
	start := cl - 20
	if start < 0 {
		start = 0
	}
	for ri := cl - 1; ri >= start; ri-- {
		if ids[ri] >= 0 && ids[ri] < V {
			age := float64(cl - ri) // 1 = just seen, 20 = oldest in window
			pen := 0.3 + 0.035*age
			logits[ids[ri]] *= pen
		}
	}
	// Bigram blocking — Q postgpt_q.c:1407-1408.
	if cl >= 2 {
		last := ids[cl-2]
		for ri := 0; ri < cl-1; ri++ {
			if ids[ri] == last && ids[ri+1] >= 0 && ids[ri+1] < V {
				logits[ids[ri+1]] *= 0.2
			}
		}
	}
}
