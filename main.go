package main

import "fmt"

// Converter o ponto de ebulição da água de Kelvin para celsius.

func main() {

	// Definindo a temperatura em Kelvin
	var temperaturaKelvin float64 = 373.15

	// Convertendo para Celsius
	var temperaturaCelsius float64 = temperaturaKelvin - 273.15

	// Exibindo o resultado
	fmt.Printf("A temperatura em Celsius é: %.2f\n", temperaturaCelsius)

}
