package status

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	Year               int     = 2022
	porcentajeAGuardar float64 = 73.82
	dineroPromedio     float64 = 240.00
)

func Resumen(sumas ...string) {
	var tenemos [][]float64

	tenemos = append(tenemos, ReadFiles(sumas[0])) // Gastados
	tenemos = append(tenemos, ReadFiles(sumas[1])) // Guardados
	tenemos = append(tenemos, ReadFiles(sumas[2])) // Extra

	fmt.Printf("Dia %.2f\n", automaticTime(Year))
	fmt.Printf("Presupuesto: %.2f \n", presupuesto())
	Falta := presupuesto() - sumaTenemos(tenemos)
	fmt.Printf("Falta: %.2f\n", Falta)
	fmt.Printf("Status %.2f\n", (0-Falta)/(porcentajeAGuardar*dineroPromedio))
	fmt.Println("Libres", dineroPromedio*((100-porcentajeAGuardar)/100))

}

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

func ReadFiles(arg string) (total []float64) {
	file, _ := os.Open(arg)
	r := csv.NewReader(bufio.NewReader(file))
	r.Comma = ','

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		num, _ := strconv.Atoi(record[1])

		total = append(total, float64(num))
	}

	return
}
