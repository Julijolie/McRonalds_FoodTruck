package main

import (
	"McRonalds/operacoes"
	"McRonalds/produtos"
	"fmt"
	"os"
)

<<<<<<< HEAD
var listaDeProdutos []produtos.Produto
var proximoId = 0
var totalPedExpedidos = 0

// var tempoExpedidoTotal time.Time
=======
type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

type Pedido struct {
	ID         int
	Delivery   bool
	Produtos   [10]Produto
	ValorTotal float64
	Status     bool
}

var produtos []Produto
var proximoId = 1
>>>>>>> aec3b74d5cdde5415d8e0746f99ba592e47eb5d8

func main() {
	operacoes.LerProdAdcLista()
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
<<<<<<< HEAD
		case 1: // cadastra
			if len(listaDeProdutos) < 50 {
				operacoes.CadastraProduto()
			} else {
				fmt.Println("Limite de 50 produtos atingido!")
			}

		case 2: //visualiza produtos
			operacoes.ProdutosCadastrados()

		case 3: //busca pelo id
			operacoes.BuscaProduto()

		case 4: // remover um produto
			fmt.Print("Digite o ID do produto que você deseja remover: ")
			var idParaRemover int
			fmt.Scanln(&idParaRemover)
			operacoes.RemoveProdutoPorID(idParaRemover)

		case 5:
=======
			case 1: // cadastra
				if len(produtos) < 50 {
					cadastraProduto()
				} else {
					fmt.Println("Limite de 50 produtos atingido!")
				}

			case 2: //visualiza produtos
				produtosCadastrados()

			case 3: //busca pelo id
				fmt.Println("Digite o ID do produto:")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				produtoID, err := strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println("O id necessita ser um inteiro!")
				}else{
					fmt.Printf("Produto Encontrado:\nID: %d, Nome: %s, Descrição: %s, Valor: %.2f, Quantidade: %d\n", buscaProduto(produtoID).ID, buscaProduto(produtoID).Nome, buscaProduto(produtoID).Descricao, buscaProduto(produtoID).Valor, buscaProduto(produtoID).Quantidade)
				}
			case 4: // remover um produto
				fmt.Print("Digite o ID do produto que você deseja remover: ")
				var idParaRemover int
				fmt.Scanln(&idParaRemover)
				removeProdutoPorID(idParaRemover)

		case 6:
>>>>>>> aec3b74d5cdde5415d8e0746f99ba592e47eb5d8
			// remover todos os produtos
			listaDeProdutos = nil
			fmt.Println("Todos os produtos foram removido")

<<<<<<< HEAD
		case 6:
			operacoes.BuscarProdutoNome()
=======
		
>>>>>>> aec3b74d5cdde5415d8e0746f99ba592e47eb5d8

			/*
				case 7:
				ExpedirPedido()

				case 8:
				MediaTempoExpedicao()
			*/

		case 9:
			os.Exit(0)

			default:
				fmt.Println("Opção inválida!")
			}
			fmt.Println("-------------------------------")
		}
	}


<<<<<<< HEAD
=======
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

func buscaProduto(produtoId int)Produto{
	//para encontrar o produto pelo id
	for _, p := range produtos {
		if produtoId == p.ID{
			return p
		} else {
			fmt.Println("Não achamos o produto!")
			break
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

func AdicionaPedido(id int,novoPedido Pedido) {
	for _, p := range produtos {
		if id == p.ID {
			if p.Quantidade == 0 {
				fmt.Println("O produto está sem estoque")
				proximoIdPedido--
			}else{
				novoPedido.ValorTotal = novoPedido.ValorTotal + p.Valor
				novoPedido.Produtos[i] = p
				i++
				p.Quantidade--
			}
		} else {
			fmt.Println("Nenhum produto encotrado com esse ID")
		}
	}
}



func expedirPedidos(){
	
}

//EXTRAS

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

		//cria novo produto
		novoProduto := Produto{}

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

		//Adiciona os produtos na lista
		produtos = append(produtos, novoProduto)
	}
}

>>>>>>> aec3b74d5cdde5415d8e0746f99ba592e47eb5d8
/*

func ExpedirPedido(){
	pedidoExpedido := removePedido() // pedidoExpedido = {1; 2 xburguers, 3 xfrango, 1 xpeixe; 200; 18:00}
	dataExpedido = time.Now // 18:30

	tempoSub = dataExpedido.Sub(pedidoExpedido.Data).Minutes()

	tempoExpedidoTotal = tempoExpedidoTotal + tempoSub
	totalPedExpedidos++
}

func pedidosCadastrados()
	fmt.Println("Pedidos casdastrados:")
	for _, p := range pedidos {
		fmt.Printf("ID: %d, Nome: %s, Descrição: %s, Valor: %.2f, Quantidade: %d\n", p.ID, p.Nome, p.Descricao, p.Valor, p.Quantidade)
	}
}
*/

/*
func MediaTempoExpedicao(){
	fmt.Print(float64(tempoExpedidoTotal)/float64(totalPedExpedidos))
}
*/
