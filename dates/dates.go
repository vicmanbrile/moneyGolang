package dates

import (
	"math"
	"time"
)

var (
	DAY          float64 = 1
	DAYS_WEEK    float64 = DAY * 10
	DAYS_MOUNTH  float64 = DAYS_WEEK * 3
	DAYS_YEAR    float64 = DAYS_MOUNTH * MOUNTHS_YEAR
	MOUNTHS_YEAR float64 = 1 * 12
)

type PriceInDays float64

func ToPriceInDays(Money float64, Average float64) PriceInDays {
	return PriceInDays(Money / Average)
}

var (
	Today DayOfYear = DayOfYear(time.Now().YearDay())
)

type DayOfYear int

func (dfy DayOfYear) Mounth() float64 {
	return math.Ceil(float64(dfy) / DAYS_MOUNTH)
}

func (dfy DayOfYear) Week() float64 {
	return math.Ceil(float64(dfy) / DAYS_WEEK)
}

func (dfy DayOfYear) ToFraction() (TF TimeFraction) {

	{
		reduce := float64(dfy)

		{ // Se calculan los años
			TF.Years = math.Floor(reduce / DAYS_YEAR)
			reduce -= TF.Years * DAYS_YEAR
		}

		{
			TF.Mounts = math.Floor(reduce / DAYS_MOUNTH)
			reduce -= TF.Mounts * DAYS_MOUNTH
		}

		{
			TF.Weeks = math.Floor(reduce / DAYS_WEEK)
			reduce -= TF.Weeks * DAYS_WEEK
		}

		{
			TF.Days = math.Floor(reduce / DAY)
			reduce -= TF.Weeks * DAY
		}

	}

	return
}

type TimeFraction struct {
	Years  float64
	Mounts float64
	Weeks  float64
	Days   float64
}

func (tf TimeFraction) Total() float64 {

	var total float64

	total += tf.Years * DAYS_YEAR
	total += tf.Mounts * DAYS_MOUNTH
	total += tf.Weeks * DAYS_WEEK
	total += tf.Days * DAY

	return total
}

func (tf TimeFraction) Diff(tfp TimeFraction) (t DayOfYear) {

	t = DayOfYear(tf.Total() - tfp.Total())

	return
}
