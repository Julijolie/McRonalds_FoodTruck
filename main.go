package main

import (
	"McRonalds/operacoes"
	"McRonalds/produtos"
	"fmt"
	"os"
)


var listaDeProdutos []produtos.Produto
var proximoId = 1
var totalPedExpedidos = 0

// var tempoExpedidoTotal time.Time


func main() {
	operacoes.LerProdAdcLista()
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Cadastrar produto")
		fmt.Println("2. Visualizar produtos cadastrados")
		fmt.Println("3. Busca do produto pelo id")
		fmt.Println("4. Remover 1 produto")
		fmt.Println("5. Remover todos os produtos")
		fmt.Println("6. Busca do produto pelo nome")

		var escolha int
		fmt.Scanln(&escolha)

		switch escolha {
			
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

		case 5:	// remover todos os produtos
			listaDeProdutos = nil
			fmt.Println("Todos os produtos foram removido")

		case 6: // busca produto pelo nome
			operacoes.BuscarProdutoNome()

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



//EXTRAS

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
