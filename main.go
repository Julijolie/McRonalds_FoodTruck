package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Produto struct {
	ID        int
	Nome      string
	Descricao string
	Valor     float64
	Quantidade int
}

var produtos []Produto

func main() {
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Cadastrar produto")
		fmt.Println("2. Remover 1 produto")
		fmt.Println("3. Remover todos os produtos")
		fmt.Println("4. Busca do produto")
		fmt.Println("5. Sair")

		var escolha int
		fmt.Scanln(&escolha)

		switch escolha {
		case 1: // cadastra
			if len(produtos) < 50 {
				cadastraProduto()
			} else {
				fmt.Println("Limite de 50 produtos atingido!")
			}

		case 2:
			// remover um produto
			fmt.Println("Remover Produto")

		case 3:
			// remover todos os produtos
			fmt.Println("Busca do produto")

		case 4:
			fmt.Println("Produtos cadastrados:")
			for _, p := range produtos {
				fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Valor: %.2f, Quantidade: %d\n", p.ID, p.Nome, p.Descricao, p.Valor, p.Quantidade)
			}

		case 5:
			os.Exit(0)

		default:
			fmt.Println("Opção inválida!")
		}
		fmt.Println("-------------------------------")
	}
}

func cadastraProduto() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		novoProduto := Produto{}
		novoProduto.ID = len(produtos) + 1

		fmt.Print("Nome do produto (digite -1 para sair): ")
		scanner.Scan()
		nome := scanner.Text()

		if nome == "-1" {
			break
		}
		novoProduto.Nome = nome

		fmt.Print("Descrição do produto: ")
		scanner.Scan()
		novoProduto.Descricao = scanner.Text()

		fmt.Print("Valor do produto: ")
		scanner.Scan()
		valorStr := scanner.Text()
		valor, err := strconv.ParseFloat(valorStr, 64) // tratamento de exceção - valor
		if err != nil {
			fmt.Println("Erro ao ler valor. Certifique-se de digitar um número válido.")
			fmt.Println("-------------------------------------------------------------")
			continue
		}
		novoProduto.Valor = valor

		fmt.Print("Quantidade do produto: ")
		scanner.Scan()
		quantidadeStr := scanner.Text()
		quantidade, err := strconv.Atoi(quantidadeStr) // tratamento de exceção - quantidade
		if err != nil {
			fmt.Println("Erro ao ler quantidade. Certifique-se de digitar um número válido.")
			fmt.Println("-------------------------------------------------------------")
			continue
		}
		novoProduto.Quantidade = quantidade

		fmt.Println("---------------------------------")

		produtos = append(produtos, novoProduto)
	}
}
