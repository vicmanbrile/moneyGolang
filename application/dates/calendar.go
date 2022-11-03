package dates

var (
	Calendar = make([]byte, int(DAYS_YEAR))
)

type RangeCalendar struct {
	count  float64
	range1 int
	range2 int
}

func addDays(Calendar *[]float64, rc RangeCalendar) {
	for i := range (*Calendar)[rc.range1:rc.range2] {
		ls := (*Calendar)[rc.range1:rc.range2]

		ls[i] += rc.count
	}
}
