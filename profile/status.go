package profile

type Registers struct {
	Entries []struct {
		Week  float64 `bson:"week"`
		Money float64 `bson:"money"`
	} `bson:"entries"`
}

func (r *Registers) Budgets() (Bdgt float64) {
	for _, value := range r.Entries {
		Bdgt += value.Money
	}

	return
}
