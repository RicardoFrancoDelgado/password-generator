package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
)

func generatePassword(length int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"

	allChars := lowerCase + upperCase + numbers + special

	mandatory := []byte{
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
	fmt.Println("Bem-vindo ao Gerador de Senhas!")
	fmt.Println("Escolha o tamanho da senha (mínimo 6 caracteres):")

	var length int

	for {
		fmt.Print("Tamanho da senha: ")
		_, err := fmt.Scan(&length)
		if err != nil || length < 6 {
			fmt.Println("Por favor, insira um número válido maior ou igual a 6.")
			continue
		}
		break
	}

	fmt.Println("Gerando sua senha, aguarde...")
	time.Sleep(time.Second * 5)

	senha := generatePassword(length)
	fmt.Printf("Senha gerada com sucesso: %s\n", senha)

	err := clipboard.WriteAll(senha)
	if err != nil {
		fmt.Println("Erro ao copiar para área de transferência!")
		return
	}

	fmt.Println("Sua senha foi copiada para a área de transferência")
}
