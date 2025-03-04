# Gerador de Senhas

Este é um simples gerador de senhas escrito em Go. Ele gera uma senha aleatória com letras maiúsculas, minúsculas, números e caracteres especiais. A senha gerada é copiada automaticamente para a área de transferência.

## Como usar

1. Clone o repositório:
    ```sh
    git clone <URL_DO_REPOSITORIO>
    cd password-generator
    ```

2. Compile o código:
    ```sh
    go build -o password-generator main.go
    ```

3. Execute o programa:
    ```sh
    ./password-generator
    ```

4. Siga as instruções no terminal para gerar uma senha.

## Requisitos

- Go 1.16 ou superior
- Biblioteca `github.com/atotto/clipboard`

## Instalação da biblioteca

Para instalar a biblioteca `clipboard`, execute:
```sh
go get github.com/atotto/clipboard
```

## Funcionalidades

- Geração de senhas com comprimento mínimo de 6 caracteres.
- Inclusão obrigatória de pelo menos um número e um caractere especial.
- Copia automaticamente a senha gerada para a área de transferência.