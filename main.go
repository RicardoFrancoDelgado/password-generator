package main

import (
	"fmt"
	"math/rand"
)

func generatePassword(length int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"

	allChars := lowerCase + upperCase + numbers + special

	mandatory := []byte{
		upperCase[rand.Intn(len(upperCase))],
		numbers[rand.Intn(len(numbers))],
		special[rand.Intn(len(special))],
	}

	password := make([]byte, length-len(mandatory))

	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}
	password = append(password, mandatory...)

	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

func main() {
	fmt.Println("Bem-vindo(a) ao seu gerador automático de senha:")
	fmt.Print("Informe quantos carácteres você quer na sua senha:")
	defer fmt.Println("Senha gerada com sucesso!")

	var lenPassword int
	fmt.Scan(&lenPassword)
	password := generatePassword(lenPassword)
	fmt.Println(password)
}
