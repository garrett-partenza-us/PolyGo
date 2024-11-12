package polynomial

func (p Polynomial) Add(other Polynomial) Polynomial {
	sum := Polynomial{Terms: append(p.Terms, other.Terms...)}
	SortTerms(&sum)
	coefficientSums := make(map[int64]float64)
	for _, term := range sum.Terms {
		coefficientSums[int64(term.Power)] += term.Coefficient
	}
	terms := make([]Term, 0, len(coefficientSums))
	for power, coefficient := range coefficientSums {
		term := Term{Power:power, Coefficient:coefficient}
		terms = append(terms, term)
	}
	return Polynomial{Terms:terms}
}

func (p Polynomial) Sub(other Polynomial) Polynomial {
	inverse := Negate(&other)
	return p.Add(inverse)
}
