package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/vicmanbrile/moneyGolang/profile"
	"github.com/vicmanbrile/moneyGolang/status"
)

func main() {

	Filename := flag.String("file", "", "string")
	flag.Parse()
	Registro := flag.Args()

	file, err := ioutil.ReadFile(*Filename)
	if err != nil {
		fmt.Printf("Error al convertir a JSON: %v", err)
	}

	data := &profile.Perfil{}

	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Printf("Error al convertir a JSON: %v", err)
	}

	data.PrintTable()

	status.Resumen(Registro...)
}
