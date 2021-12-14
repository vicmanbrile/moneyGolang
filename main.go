package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	err error
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

	fmt.Printf("%+v", len(data.Creditos))

}
