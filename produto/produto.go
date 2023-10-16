package produto

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
)

type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

var ListaProdutos [50]Produto
var tamanhoListaProdutos = 0
var proximoId = 0


func ProdutosCadastrados() {
	fmt.Println("Produtos cadastrados:")
	for _, p := range ListaProdutos {
		fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Valor: %.2f, Quantidade: %d\n", p.ID, p.Nome, p.Descricao, p.Valor, p.Quantidade)
	}
}

func BuscaProduto(produtoID string) Produto { 
	//para encontrar o produto pelo id
	for _, p := range ListaProdutos {
		if produtoID == fmt.Sprint(p.ID) {
			fmt.Printf("Produto Encontrado:\nID: %d, Nome: %s, Descrição: %s, Valor: %.2f, Quantidade: %d\n", p.ID, p.Nome, p.Descricao, p.Valor, p.Quantidade)
			return p
		} else {
			fmt.Println("Nenhum produto encotrado com esse ID")
		}
	}
	return Produto{}
}

func CadastraProduto() {
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

		AdicionaProdutoLista(novoProduto)
	}
}

func AdicionaProdutoLista(prod Produto) {
	if tamanhoListaProdutos > 50 {
		fmt.Println("Limite de 50 produtos atingido!")
		return
	}
	ListaProdutos[tamanhoListaProdutos] = prod
	tamanhoListaProdutos++
}


func RemoveProdutoPorID(id int) {
	if tamanhoListaProdutos <= 0{
		fmt.Println("Não há produtos cadastrados")
		return
	}
	ListaProdutos[id] = Produto{}
	tamanhoListaProdutos--
}

func BuscarProdutoNome() {
	listaResultado := make([]Produto, 0)

	fmt.Println("Qual o nome produto que deseja encontrar?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	produtoNome := scanner.Text()

	fmt.Println("\n------------------------------")

	// verificar quais produtos da lista produtos começa com as letras do nome pesquisado
	for _, produto := range ListaProdutos {
		if strings.HasPrefix(strings.ToLower(produto.Nome), produtoNome) {
			listaResultado = append(listaResultado, produto)
		}
	}

	// exibir listaresultado
	for _, resultado := range listaResultado {
		fmt.Println("ID", resultado.ID, ":", resultado.Nome, "|", resultado.Descricao, "| Valor:", resultado.Valor, "| Quantidade:", resultado.Quantidade)
	}
}