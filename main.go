package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	file, err := ioutil.ReadFile("filename.json")
	if err != nil {
		fmt.Printf("Error al convertir a JSON: %v", err)
	}

	data := &Perfil{}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Printf("Error al convertir a JSON: %v", err)
	}

	data.PrintTable()

}
