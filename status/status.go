package status

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Status struct { // Our example struct, you can use "-" to ignore a field
	Id    string  `csv:"Clave"`
	Value float64 `csv:"Valor"`
}

var (
// Year int = 2022
// porcentajeAGuardar float64 = 73.82
// dineroPromedio     float64 = 240.00
)

func Resumen(sumas ...string) {

	var all = make([][]*Status, 3)

	for index, value := range sumas {
		result := ReadFiles(value)
		all[index] = result
	}

	for _, value := range all {
		for _, value := range value {
			fmt.Println(*value)
		}
	}

	// fmt.Printf("Dia %.2f\n", automaticTime(Year))
	// fmt.Printf("Presupuesto: %.2f \n", presupuesto())
	// Falta := presupuesto() - sumaTenemos(tenemos)
	// fmt.Printf("Falta: %.2f\n", Falta)
	// fmt.Printf("Status %.2f\n", (0-Falta)/(porcentajeAGuardar*dineroPromedio))
	// fmt.Println("Libres", ((dineroPromedio * ((100 - porcentajeAGuardar) / 100)) * automaticTime(Year)))

}

/*
func automaticTime(year int) float64 {
	inicio := time.Date(year, time.January, 0, 0, 0, 0, 0, time.UTC)
	calculo := time.Now().AddDate(0, 0, -inicio.Day()).Day()

	return float64(calculo)
}

func presupuesto() float64 {
	return dineroPromedio * (porcentajeAGuardar / 100) * float64(automaticTime(Year))
}

func sumaTenemos(total [][]float64) float64 {
	var result float64
	for _, value := range total[0] {
		result += value
	}
	for _, value := range total[1] {
		result += value
	}
	for _, value := range total[2] {
		result -= value
	}

	return result
}
*/

func ReadFiles(arg string) (total []*Status) {
	file, _ := os.OpenFile(arg, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err := gocsv.UnmarshalFile(file, &total); err != nil { // Load clients from file
		panic(err)
	}

	return
}
