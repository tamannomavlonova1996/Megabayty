package billing

func Calculate1000(useMb float64) float64 {
	abonPlata := 35.0
	oneMb := 0.06
	sum := 0.0
	sum = (useMb * oneMb) + abonPlata
	return sum
}
