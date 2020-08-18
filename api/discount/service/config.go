package service

type uniqueHarryPotterRule struct {
	Qty             int
	PercentDiscount int
	Msg             string
}

var (
	uniqueHarryPotterRules = []uniqueHarryPotterRule{{
		Qty:             2,
		PercentDiscount: 10,
		Msg:             "buy 2 unique Harry Potter series books discount 10%",
	}, {
		Qty:             3,
		PercentDiscount: 11,
		Msg:             "buy 3 unique Harry Potter series books discount 11%",
	}, {
		Qty:             4,
		PercentDiscount: 12,
		Msg:             "buy 4 unique Harry Potter series books discount 12%",
	}, {
		Qty:             5,
		PercentDiscount: 13,
		Msg:             "buy 5 unique Harry Potter series books discount 13%",
	}, {
		Qty:             6,
		PercentDiscount: 14,
		Msg:             "buy 6 unique Harry Potter series books discount 14%",
	}, {
		Qty:             7,
		PercentDiscount: 15,
		Msg:             "buy 7 unique Harry Potter series books discount 15%",
	}}
	harryPotterSeries = []string{
		"Harry Potter and the Philosopher's Stone (I)",
		"Harry Potter and the Chamber of Secrets (II)",
		"Harry Potter and the Prisoner of Azkaban (III)",
		"Harry Potter and the Goblet of Fire (IV)",
		"Harry Potter and the Order of the Phoenix (V)",
		"Harry Potter and the Half-Blood Prince (VI)",
		"Harry Potter and the Deathly Hallows (VII)",
	}
)
