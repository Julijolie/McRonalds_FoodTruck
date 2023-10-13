package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

var produtos []Produto
var proximoId = 1

func main() {
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Cadastrar produto")
		fmt.Println("2. Visualizar produtos cadastrados")
		fmt.Println("3. Busca do produto pelo id")
		fmt.Println("4. Remover 1 produto")
		fmt.Println("5. Remover todos os produtos")
		fmt.Println("6. Sair")

		var escolha int
		fmt.Scanln(&escolha)

		switch escolha {
		case 1: // cadastra
			if len(produtos) < 50 {
				cadastraProduto()
			} else {
				fmt.Println("Limite de 50 produtos atingido!")
			}

		case 2: //visualiza produtos
			produtosCadastrados()

		case 3: //busca pelo id
			buscaProduto()

		case 4: // remover um produto
			fmt.Print("Digite o ID do produto que você deseja remover: ")
			var idParaRemover int
			fmt.Scanln(&idParaRemover)
			removeProdutoPorID(idParaRemover)

		case 5:
			// remover todos os produtos
			produtos = nil
			fmt.Println("Todos os produtos foram removido")

		case 6:
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

		fmt.Print("Nome do produto (digite -1 para sair): ")
		scanner.Scan()
		nome := scanner.Text()

		if nome == "-1" {
			break
		}
		novoProduto.Nome = nome

		novoProduto.ID = proximoId
		proximoId++

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

func produtosCadastrados() {
	fmt.Println("Produtos cadastrados:")
	for _, p := range produtos {
		fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Valor: %.2f, Quantidade: %d\n", p.ID, p.Nome, p.Descricao, p.Valor, p.Quantidade)
	}
}

func buscaProduto() {
	fmt.Println("Digite o ID do produto:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	produtoID := scanner.Text()

	//para encontrar o produto pelo id
	for _, p := range produtos {
		if produtoID == fmt.Sprint(p.ID) {
			fmt.Printf("Produto Encontrado:\nID: %d, Nome: %s, Descrição: %s, Valor: %.2f, Quantidade: %d\n", p.ID, p.Nome, p.Descricao, p.Valor, p.Quantidade)
		} else {
			fmt.Println("Nenhum produto encotrado com esse ID")
		}
	}
}

func removeProdutoPorID(id int) {
	for i, produto := range produtos {
		if produto.ID == id {
			produtos = append(produtos[:i], produtos[i+1:]...)
			fmt.Printf("Produto com ID %d removido com sucesso!\n", id)
			return
		}
	}
	fmt.Printf("Produto com ID %d não encontrado.\n", id)
}
