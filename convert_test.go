package gobet

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMoneylineToDecimal(t *testing.T) {
	formats := []OddType{Decimal, Moneyline, Fractional}
	for _, x := range []struct {
		o         Odd
		decimal   Odd
		moneyline Odd
	}{
		{"3.30", "3.30", "+230"},
		{"3.30", "3.30", "+230"},
		{"+230", "3.30", "+230"},
		{"-280", "1.36", "-280"},
		{"1.357143", "1.357143", "-280"},
	} {
		for _, format := range formats {
			s, f := x.o.Convert(format)
			fmt.Printf("%s %s to %s -> %s %f\n", x.o.Type(), x.o, format, s, f)
			switch format {
			case Decimal:
				require.Equal(t, x.decimal, s)
			case Moneyline:
				require.Equal(t, x.moneyline, s)
			}
		}
	}
}

func TestClient_Odds2(t *testing.T) {
	fmt.Println(Odd("1.357143").Convert(Moneyline))
}
