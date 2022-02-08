package profile

type Registers struct {
	Entries []struct {
		Week  float64 `json:"week"`
		Money float64 `json:"money"`
	} `json:"entries"`
}

func (r *Registers) Budgets() (Bdgt Budget) {
	for _, value := range r.Entries {
		Bdgt.Entries += value.Money
	}

	return
}

type Budget struct {
	Entries float64
}
