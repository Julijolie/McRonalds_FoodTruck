package operacoes

import (
	"McRonalds/produtos"
	"McRonalds/pedidos"
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
)

var listaProdutos []produtos.Produto
var proximoId int // Certifique-se de que você tem a variável proximoId definida.
var valorTotalGanho float64 = 0
var pedidosEncerrados = 0
var j = 1

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

func BuscaProduto(produtoID string) produtos.Produto{
	var produto produtos.Produto

	//para encontrar o produto pelo id
	encontrado := false
	for _, p := range listaProdutos {
		if produtoID == fmt.Sprint(p.ID) {
			encontrado = true
			produto = p
		}
	}

	if !encontrado {
		fmt.Println("Nenhum produto encontrado com esse ID")
	}

	return produto
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

func AdicionaPedido() {
	var i = 0
	novoPedido := pedidos.Pedido{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Caso o pedido seja delivery selecione 1, caso contrário, 2: ")
	var escolha string
	scanner.Scan()
	escolha = scanner.Text()

	switch escolha {
		case "1":
			novoPedido.Delivery = true
		case "2":
			novoPedido.Delivery = false
		default:
			fmt.Println("Opção inválida!")
	}
	

	for j:= 0; j < 10; j++{
		fmt.Println("Informe o id do produto: ")
		scanner.Scan()
		id := scanner.Text()
		if id == "-1" {
			break
		}
		fmt.Println("Qual a quantidade que você deseja?")
		scanner.Scan()
		quantidadeString := scanner.Text()
		quantidade, err := strconv.Atoi(quantidadeString)
		if err != nil {
			log.Fatal(err)
		}

		
		for _, p := range listaProdutos {
			if id == fmt.Sprint(p.ID) {
				if p.Quantidade == 0 {
					fmt.Println("O produto está sem estoque")
				}else{
					novoPedido.ValorTotal = novoPedido.ValorTotal + p.Valor
					novoPedido.Produtos[i] = p
					i++
					p.Quantidade-= quantidade
				}
			} else {
				break
			}
		}
	}

	
}

func mostrarProdutosPedido(p pedidos.Pedido){
	for _, item := range p.Produtos {
		fmt.Printf("Pedidos:\n ID: %d, Nome: %s\n, Descrição: %s\n, Valor: %.2f\n", item.ID, item.Nome, item.Descricao, item.Valor)
	}
}

func ExpedirPedido() {
	var pedido pedidos.Pedido
	var delivery bool
	var valor float64
	for {
		if j == pedido.ID {
			delivery = pedido.Delivery
			valor = pedido.ValorTotal
			valorTotalGanho += valor
			j++
			break
		}
	}
	fmt.Println("Pedido a ser expedido:")
	fmt.Printf("Id: %d\n Delivery: %t\n Produtos:", j, delivery)
	mostrarProdutosPedido(pedido)
	fmt.Printf("Valor da compra: %.2f", valor)
	pedidosEncerrados ++

}

func ExibirMetricas() {
	fmt.Printf("Produtos cadstrados: %d\n Numero de pedidos encerrados: %d, Faturamento Total: %.2f", len(listaProdutos), pedidosEncerrados, valorTotalGanho)
}
