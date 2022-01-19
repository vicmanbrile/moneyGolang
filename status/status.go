package status

import (
	"fmt"
	"time"

	"github.com/vicmanbrile/moneyGolang/profile"
)

type Registers struct {
	Spent []struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	} `json:"spent"`
	Saved []struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	} `json:"saved"`
	Extras []struct {
		Week  float64 `json:"week"`
		Extra float64 `json:"extra"`
		Days  float64 `json:"days"`
	} `json:"extras"`
}

// Presupuesto por ganados | Falta | Libres | Estatus | Libres
// Presupuesto por ahora   | Falta | Libres | Estatus | Libres

func (r *Registers) BudgetsNow(w profile.Wallet) float64 {
	days := automaticTime(Year)
	return days * w.Average
}

func (r *Registers) BudgetsWon(w profile.Wallet) {}

var (
	Year               int     = 2022
	porcentajeAGuardar float64 = 73.82
	dineroPromedio     float64 = 240.00
)

func Resumen(sumas ...string) {
	/*
		var all = make([][]*Valores, 3)

		for index, value := range sumas {
			result := ReadFiles(value)
			all[index] = result
		}

		var tenemos float64
		for _, v := range all[0] {
			tenemos += v.Value
		}
		for _, v := range all[1] {
			tenemos += v.Value
		}
		for _, v := range all[2] {
			tenemos -= v.Value
		}
	*/

	fmt.Printf("Dia %.2f\n", automaticTime(Year))
	fmt.Printf("Presupuesto: %.2f \n", presupuesto())
	Falta := presupuesto() // - tenemos
	fmt.Printf("Falta: %.2f\n", Falta)
	fmt.Printf("Status %.2f\n", (0-Falta)/((porcentajeAGuardar/100)*dineroPromedio))
	fmt.Printf("Libres: %.2f\n", ((dineroPromedio * ((100 - porcentajeAGuardar) / 100)) * automaticTime(Year)))

}

func automaticTime(year int) float64 {
	inicio := time.Date(year, time.January, 0, 0, 0, 0, 0, time.UTC)
	calculo := time.Now().AddDate(0, 0, -inicio.Day()).Day()

	return float64(calculo)
}

func presupuesto() float64 {
	return dineroPromedio * (porcentajeAGuardar / 100) * float64(automaticTime(Year))
}
