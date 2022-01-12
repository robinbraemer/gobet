package gobet

import "math"

// ArbPct returns the probability that both events happen.
// A value below 100% means an arbitrage opportunity.
//
// Arbitrage % = ((1 / decimal odds for outcome A)) + ((1 / decimal odds for outcome B))
func ArbPct(oddA float64, oddB float64) float64 { return Pct(oddA) + Pct(oddB) }

func Pct(odd float64) float64 { return 1 / odd }

// ArbProfit returns the expected arbitrage profit.
//
// Having found a surebet, we then need to calculate the profit we will receive based
// on the amount of money we are willing to invest.
//
// If, for example, you are wanting to place $500 stake on the tennis surebet above,
// you would calculate the profit using the following formula:
//
// Profit = (Investment / Arbitrage %) – Investment
func ArbProfit(arbPct, stake float64) float64 { return stake/arbPct - stake }

// ArbStakes distributes the total stake across two bets.
//
// It calculates how your investment needs to be broken down in terms of stakes across both bets.
// This is so that you are returning the same profit regardless of which outcome wins.
// The idea is to return the same profit regardless of whether the first or second outcome is successful,
// so it is critical to use the correct stakes – if not, you could find that one outcome is more profitable
// than the other or that you actually lose money if one outcome wins.
//
// To calculate the individual stakes:
// Individual bets = (Investment x Individual Arbitrage %) / Total Arbitrage %
func ArbStakes(totalStake float64, individualPcts []float64) (individualStakes []float64) {
	var totalPct float64
	for _, pct := range individualPcts {
		totalPct += pct
	}
	individualStakes = make([]float64, len(individualPcts))
	for i, pct := range individualPcts {
		individualStakes[i] = totalStake * pct / totalPct
	}
	// sum of individualStakes should equal to totalStake
	return
}

// Round rounds x to the specified decimal places.
func Round(x float64, places uint) float64 {
	p := math.Pow(10, float64(places))
	return math.Floor(x*p) / p
}
