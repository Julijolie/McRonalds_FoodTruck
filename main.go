package main

import (
	"bufio"
	"fmt"
	"os"
	prd "McRonald/produto"
	ext "McRonald/extra"
	ped "McRonald/pedidos"
)

var totalPedExpedidos = 0

// var tempoExpedidoTotal time.Time

func main() {
	ext.LerProdAdcLista()
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Cadastrar produto")
		fmt.Println("2. Visualizar produtos cadastrados")
		fmt.Println("3. Cadastrar pedido")
		fmt.Println("4. Visualizar pedidos cadastrados")
		fmt.Println("5. Busca do produto pelo id")
		fmt.Println("6. Pesquisar pelo nome do produto")
		fmt.Println("7. Remover 1 produto")
		fmt.Println("8. Remover todos os produtos")
		//fmt.Println("9. Expedir pedidos")
		//fmt.Println("10. Tempo médio")
		//fmt.Println("11. metrícas")
		fmt.Println("12. Sair")

		var escolha int
		fmt.Scanln(&escolha)

		switch escolha {
		case 1: // cadastra produto
			if len(prd.ListaProdutos) < 50 {
				prd.CadastraProduto()
			} else {
				fmt.Println("Limite de 50 produtos atingido!")
			}

		
		case 2: //Produtos cadastrados
			prd.ProdutosCadastrados()

		
		case 3: //Cadastra pedido
			ped.AdicionaPedido()	
			
		
		case 4: //Pedidos cadastrados
			ped.PedidosCadastrados()
		
		
		case 5: //busca pelo id
			fmt.Println("Digite o ID do produto:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			produtoID := scanner.Text()
			prd.BuscaProduto(produtoID)


		case 6:
			prd.BuscarProdutoNome()


		case 7: //Remover um produto
			fmt.Print("Digite o ID do produto que você deseja remover: ")
			var idParaRemover int
			fmt.Scanln(&idParaRemover)
			prd.RemoveProdutoPorID(idParaRemover)


		case 8: //Remover todos os produtos
			prd.ListaProdutos = [50]prd.Produto{}
			fmt.Println("Todos os produtos foram removido")

			/*
				case 9:
				ExpedirPedido()

				case 10:
				MediaTempoExpedicao()
			*/

		case 12:
			os.Exit(0)

		default:
			fmt.Println("Opção inválida!")
		}
		fmt.Println("-------------------------------")
	}
}