package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Alumno struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	City   string `json:"city"`
	Course string `json:"course"`
	Grade  int    `json:"grade"`
	Id     int    `json:"id"`
}

type Data []Alumno

func main() {
	// Corrige el nombre del archivo JSON
	jsonFile, err := os.Open("alumnos.json")

	if err != nil {
		fmt.Println("Hubo un error al abrir el archivo:", err)
		return
	}

	defer jsonFile.Close()

	bytesValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Hubo un error al leer el archivo:", err)
		return
	}

	var data Data

	err = json.Unmarshal(bytesValue, &data)
	if err != nil {
		fmt.Println("Hubo un error al convertir JSON:", err)
		return
	}

	prom_age := 0
	for _, alumno := range data {
		fmt.Println("Nombre:", alumno.Name)
		fmt.Println("Edad:", alumno.Age)
		fmt.Println("Ciudad:", alumno.City)
		fmt.Println("Cursos:", alumno.Course)
		fmt.Println("-----------")
		prom_age += alumno.Age
	}

	if len(data) > 0 {
		prom_age /= len(data) // Calcula el promedio
	}

	fmt.Println("Promedio de edad:", prom_age)
}
