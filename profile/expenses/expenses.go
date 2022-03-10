package expenses

import (
	"math"
	"time"
)

var (
	DAY          float64 = 1
	DAYS_WEEK    float64 = DAY * 10
	DAYS_MOUNTH  float64 = DAYS_WEEK * 3
	MOUNT        float64 = 1
	MOUNTHS_YEAR float64 = 12
	DAYS_YEAR    float64 = DAYS_MOUNTH * MOUNTHS_YEAR
)

func ToPriceInDays(Money float64, Average float64) PriceInDays {
	return PriceInDays(Money / Average)
}

var (
	Today DayOfYear = DayOfYear(time.Now().YearDay())
)

type DayOfYear float64

func (dfy DayOfYear) Mounth() float64 {
	return math.Ceil(float64(dfy) / DAYS_MOUNTH)
}

func (dfy DayOfYear) Week() float64 {
	return math.Ceil(float64(dfy) / DAYS_WEEK)
}

type Credits struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Date struct {
		Mount int `json:"mount"`
		Year  int `json:"year"`
	} `json:"date"`
	Datails struct {
		Interes   float64 `json:"interes"`
		Precing   float64 `json:"precing"`
		Mensualy  int     `json:"mensualy"`
		Optionals struct {
			Percentage  float64 `json:"porcentile"`
			Suscription string  `json:"suscription"`
		} `json:"optionals"`
	} `json:"datails"`
	Spent float64 `json:"spent"`
}

func (c *Credits) Calculator(Average float64) (r Resumen) {
	r = Resumen{
		Name: c.Name,
		Type: c.Type,
	}

	{ /* Establecer los tiempos de pago */
		switch c.Type {
		case "Credit":
			{
				r.MonthFinish = float64(c.Datails.Mensualy)
			}
		case "Debt":
			{
				r.MonthFinish = Today.Mounth()
			}
		case "Percentile":
			{
				r.MonthFinish = MOUNTHS_YEAR

				{
					var Procintile = c.Datails.Optionals.Percentage * DAYS_YEAR

					c.Datails.Precing = Procintile * Average
				}
			}
		case "Suscription":
			{
				r.MonthFinish = MOUNTHS_YEAR
				switch c.Datails.Optionals.Suscription {
				case "yearly":
					{
						return
					}
				case "monthly":
					{
						c.Datails.Precing *= MOUNTHS_YEAR
					}
				}
			}
		}
	}

	/* Definimos el precio por dias de los creditos */
	if c.Datails.Interes > 0 {
		c.Datails.Precing *= (c.Datails.Interes + 1)
	}

	r.Price = ToPriceInDays(c.Datails.Precing, Average)

	if c.Spent == 0 {
		r.Paid = 0
	} else {
		r.Paid = ToPriceInDays(c.Spent, Average)
	}

	return
}

type Expenses struct {
	Creditos []Credits `json:"credit"`
}

func (e *Expenses) CalcPerfil(Average float64) AllExpenses {
	var AE = AllExpenses{}

	for _, value := range e.Creditos {
		AE.ToDoExpenses = append(AE.ToDoExpenses, value.Calculator(Average))
	}

	return AE
}

type AllExpenses struct {
	ToDoExpenses []Resumen
}

func (a *AllExpenses) PorcentileForMouthsP() float64 {
	var total float64

	for _, value := range a.ToDoExpenses {
		total += value.PorcentileForMouthsP()
	}

	return total
}

func (a *AllExpenses) PorcentileForMouthsF() float64 {
	var total float64

	for _, value := range a.ToDoExpenses {
		total += value.PorcentileForMouthsF()
	}

	return total
}

func (a *AllExpenses) PorcentileMoutnsPastFuture() float64 {
	var total float64

	for _, value := range a.ToDoExpenses {
		total += value.PorcentileMoutnsPastFuture()
	}

	return total
}

type PriceInDays float64

type Resumen struct {
	Name        string
	Type        string
	Price       PriceInDays
	Paid        PriceInDays
	MonthFinish float64
}

func (r *Resumen) PorcentileForYear() float64 {
	return float64(r.Price) / DAYS_YEAR
}

func (r *Resumen) MountsPast() float64 {
	return Today.Mounth() - 1.00
}

func (r *Resumen) MountsFuture() float64 {
	mounts := r.MonthFinish - Today.Mounth()
	if mounts == 0 {
		mounts = 1
	}

	return mounts
}

func (r *Resumen) PorcentileForMouthsP() (past float64) {
	DaysPast := DAYS_MOUNTH * r.MountsPast()

	past = float64(r.Paid) / DaysPast

	return
}

func (r *Resumen) PorcentileForMouthsF() (future float64) {
	DaysFuture := DAYS_MOUNTH * r.MountsFuture()
	future = float64(r.Price-r.Paid) / DaysFuture

	return
}

func (r *Resumen) PorcentileMoutnsPastFuture() float64 {
	Past := r.PorcentileForMouthsP()
	Future := r.PorcentileForMouthsF()

	MPast := r.MountsPast()
	MFuture := 1.00

	return ((Past * MPast) + (Future * MFuture)) / (MPast + MFuture)
}
