package gobet

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type (
	Odd     string
	OddType string
)

const (
	Decimal    OddType = "decimal"
	Fractional OddType = "fractional"
	Moneyline  OddType = "moneyline"
)

func (o *Odd) Type() OddType {
	for _, r := range *o {
		switch r {
		case '.':
			return Decimal
		case ',':
			*o = Odd(strings.Replace(string(*o), ",", ".", 1))
			return Decimal
		case '+', '-':
			return Moneyline
		case '/':
			return Fractional
		default:
			if !unicode.IsNumber(r) {
				return ""
			}
		}
	}
	return Decimal
}
func (o Odd) Valid() bool { return o.Type() != "" }

func Float(s string, t OddType) float64 {
	if t == Fractional {
		num, denum := Fraction(s)
		return num / denum
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Sprintf("could not parse %q (%s): %v", s, t, err))
	}
	return f
}

func Fraction(s string) (numerator, denominator float64) {
	i := strings.Index(s, "/")
	n, d := s[:i], s[i+1:]
	num, err := strconv.ParseFloat(n, 64)
	if err != nil {
		panic(fmt.Sprintf("could not parse %q numerator %q: %v", s, n, err))
	}
	den, err := strconv.ParseFloat(d, 64)
	if err != nil {
		panic(fmt.Sprintf("could not parse %q denumerator %q: %v", s, d, err))
	}
	return num, den
}

func (o Odd) Float() float64 { return Float(string(o), o.Type()) }

func (o Odd) Convert(to OddType) (Odd, float64) {
	from := o.Type()
	switch from {
	case Fractional:
		switch to {
		case Decimal:
			f := Float(string(o), from) + 1
			s := strconv.FormatFloat(f, 'f', 2, 64)
			return Odd(s), f
		case Moneyline:
			num, denum := Fraction(string(o))
			var f float64
			if num > denum {
				f = num / denum
			} else {
				f = denum / num
			}
			f /= 100
			s := strconv.FormatFloat(f, 'f', 0, 64)
			if f > 0 {
				s = "+" + s
			} else {
				s = "-" + s
			}
			return Odd(s), f
		case Fractional:
			return o, Float(string(o), from)
		}
	case Moneyline:
		switch to {
		case Decimal:
			f := Float(string(o), from)
			switch {
			case f > 0:
				f = f/100 + 1
				return Odd(strconv.FormatFloat(f, 'f', 2, 64)), f
			case f < 0:
				f = 100/math.Abs(f) + 1
				return Odd(strconv.FormatFloat(f, 'f', 2, 64)), f
			}
			return "0", 0
		case Fractional:
			f := Float(string(o), from)
			switch {
			case f > 0:
				s := strconv.FormatFloat(f, 'f', 2, 64)
				return Odd(s + "/100"), f / 100
			case f < 0:
				s := strconv.FormatFloat(f, 'f', 2, 64)
				return Odd("-100/" + s), -100 / f
			}
			return "0", 0
		case Moneyline:
			return o, Float(string(o), from)
		}
	case Decimal:
		switch to {
		case Moneyline:
			f := Float(string(o), from)
			switch {
			case f > 2:
				f = (f - 1) * 100
			case f > 1 && f < 2:
				f = -100 / (f - 1)
			default:
				return "0", 0
			}
			s := strconv.FormatFloat(f, 'f', 0, 64)
			if f > 0 {
				s = "+" + s
			}
			return Odd(s), f
		case Fractional:
			f := Float(string(o), from)
			return Odd(strconv.FormatFloat(f-1, 'f', 2, 64) + "/1"), f
		case Decimal:
			return o, Float(string(o), from)
		}
	}
	panic(fmt.Sprintf("unsupported conversion %s %q to %s", from, o, to))
}
