package status

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type Status struct { // Our example struct, you can use "-" to ignore a field
	Id    string  `csv:"Clave"`
	Value float64 `csv:"Valor"`
}

var (
	Year               int     = 2022
	porcentajeAGuardar float64 = 73.82
	dineroPromedio     float64 = 240.00
)

func Resumen(sumas ...string) {

	var all = make([][]*Status, 3)

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

	fmt.Printf("Dia %.2f\n", automaticTime(Year))
	fmt.Printf("Presupuesto: %.2f \n", presupuesto())
	Falta := presupuesto() - tenemos
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

func ReadFiles(arg string) (total []*Status) {
	file, _ := os.OpenFile(arg, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err := gocsv.UnmarshalFile(file, &total); err != nil { // Load clients from file
		panic(err)
	}

	return
}
