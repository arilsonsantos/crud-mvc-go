package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	//
	// router := gin.Default()
	// routes.InitRoutes(&router.RouterGroup)
	//
	// if err := router.Run(":8585"); err != nil {
	// 	log.Fatal(err)
	// }

	texto := "acidente de carro OR moto AND praia OR montanha fria OR neve"
	println(texto)

	e := regexp.MustCompile(`\b(OR|AND)\b`)
	resultados := e.FindAllString(texto, -1)

	// Imprima os resultados encontrados
	fmt.Println("Palavras encontradas:", resultados)

	// Dividir o texto usando "!ou!" como delimitador
	parts := strings.Split(texto, "!ou!")

	for _, valor := range parts {
		fmt.Println(valor)
	}

}
