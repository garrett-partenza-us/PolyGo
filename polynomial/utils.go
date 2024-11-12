package polynomial

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func (t Term) String() string {
	superscriptDigits := map[int]string{
		0: "⁰", 1: "¹", 2: "²", 3: "³", 4: "⁴",
		5: "⁵", 6: "⁶", 7: "⁷", 8: "⁸", 9: "⁹",
	}


	powerString := ""
	var power string
	if t.Power < 1 {
		powerString += "⁻"
		power = strconv.Itoa(int(t.Power) * -1)
	} else {
		power = strconv.Itoa(int(t.Power))
	}
	for _, char := range power {
		digit, _ := strconv.Atoi(string(char))
		digitString := superscriptDigits[digit]
		powerString = powerString + digitString
	}

	parity := "+"
	if t.Coefficient < 0 {
		parity = "-"
	}

	return fmt.Sprintf("%s %.2f%s", parity, t.Coefficient, powerString)
}

func (p Polynomial) String() string {
	SortTerms(&p)
	termStrings := make([]string, len(p.Terms))
	for i, term := range p.Terms {
		termStrings[i] = term.String()
	}
	polyString := strings.Join(termStrings, " ")
	if polyString[0] == '+' {
		polyString = polyString[2:]
	}
	return polyString
}

func SortTerms(p *Polynomial) {
	sort.Slice(p.Terms, func(i, j int) bool {
		return p.Terms[i].Power < p.Terms[j].Power
	})
}

func MaxPower(p *Polynomial) int64 {

	var currentMax int64 = 0
	for _, term := range p.Terms {
		if term.Power > currentMax {
			currentMax = term.Power
		}
	}
	return currentMax
}

func Negate(p *Polynomial) Polynomial {
	terms := make([]Term, len(p.Terms))
	for i, term := range p.Terms {
		terms[i] = Term{Coefficient: term.Coefficient * -1, Power: term.Power}
	}
	return Polynomial{Terms:terms}
}


