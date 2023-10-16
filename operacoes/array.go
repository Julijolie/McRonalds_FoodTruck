package operacoes

import (
	"McRonalds/produtos"
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var listaProdutos []produtos.Produto
var proximoId int // Certifique-se de que você tem a variável proximoId definida.

func CadastraProduto() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		novoProduto := produtos.Produto{}

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

		listaProdutos = append(listaProdutos, novoProduto)
	}
}

func ProdutosCadastrados() {
	fmt.Println("Produtos cadastrados:")
	for _, p := range listaProdutos {
		fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Valor: %.2f, Quantidade: %d\n", p.ID, p.Nome, p.Descricao, p.Valor, p.Quantidade)
	}
}

func BuscaProduto() {
	fmt.Println("Digite o ID do produto:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	produtoID := scanner.Text()

	//para encontrar o produto pelo id
	encontrado := false
	for _, p := range listaProdutos {
		if produtoID == fmt.Sprint(p.ID) {
			fmt.Printf("Produto Encontrado:\nID: %d, Nome: %s, Descrição: %s, Valor: %.2f, Quantidade: %d\n", p.ID, p.Nome, p.Descricao, p.Valor, p.Quantidade)
			encontrado = true
			break
		}
	}

	if !encontrado {
		fmt.Println("Nenhum produto encontrado com esse ID")
	}
}

func RemoveProdutoPorID(id int) {
	for i, produto := range listaProdutos {
		if produto.ID == id {
			listaProdutos = append(listaProdutos[:i], listaProdutos[i+1:]...)
			fmt.Printf("Produto com ID %d removido com sucesso!\n", id)
			return
		}
	}
	fmt.Printf("Produto com ID %d não encontrado.\n", id)
}

func BuscarProdutoNome() {
	listaResultado := make([]produtos.Produto, 0)

	fmt.Println("Qual o nome produto que deseja encontrar?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	produtoNome := scanner.Text()

	fmt.Println("\n------------------------------")

	// verificar quais produtos da lista produtos começa com as letras do nome pesquisado
	for _, produto := range listaProdutos {
		if strings.HasPrefix(strings.ToLower(produto.Nome), strings.ToLower(produtoNome)) {
			listaResultado = append(listaResultado, produto)
		}
	}

	// exibir listaresultado
	for _, resultado := range listaResultado {
		fmt.Println("ID", resultado.ID, ":", resultado.Nome, "|", resultado.Descricao, "| Valor:", resultado.Valor, "| Quantidade:", resultado.Quantidade)
	}
}

func LerProdAdcLista() {
	// Abrir o arquivo texto
	file, err := os.Open("ProdAdcLista.csv")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo ProdAdcLista.csv")
		return
	}
	defer file.Close()

	// Criar um leitor para o arquivo
	reader := csv.NewReader(file)
	for {
		linha, err := reader.Read()
		if err != nil {
			break
		}

		// Cria novo produto
		novoProduto := produtos.Produto{}

		novoProduto.ID = proximoId
		proximoId++

		novoProduto.Nome = linha[0]
		novoProduto.Descricao = linha[1]

		valor64, err1 := strconv.ParseFloat(linha[2], 64)
		if err1 != nil {
			fmt.Println("Erro ao converter valor para float64")
			return
		}
		novoProduto.Valor = valor64

		qtdInt, err2 := strconv.Atoi(linha[3])
		if err2 != nil {
			fmt.Println("Erro ao converter valor para inteiro")
			return
		}
		novoProduto.Quantidade = qtdInt

		// Adiciona o novo produto à lista
		listaProdutos = append(listaProdutos, novoProduto)
	}
}
