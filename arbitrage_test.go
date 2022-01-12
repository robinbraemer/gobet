package gobet

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArbPct(t *testing.T) {
	require.Equal(t, 0.99031, Round(ArbPct(1.18, 7), 5))
}

func TestArbProfit(t *testing.T) {
	require.Equal(t, 4.89, Round(ArbProfit(0.99031, 500), 2))
}
