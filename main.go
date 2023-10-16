package main

import (
	"bufio"
	"fmt"
	"os"
	prd "McRonald/produto"
	ext "McRonald/extra"
	//ped "McRonald/pedidos"
)

var totalPedExpedidos = 0

// var tempoExpedidoTotal time.Time

func main() {
	ext.LerProdAdcLista()
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Cadastrar produto")
		fmt.Println("2. Visualizar produtos cadastrados")
		fmt.Println("3. Busca do produto pelo id")
		fmt.Println("4. Remover 1 produto")
		fmt.Println("5. Remover todos os produtos")
		fmt.Println("6. Pesquisar pelo nome do produto")
		//fmt.Println("7. Expedir pedidos")
		//fmt.Println("8. Tempo médio")
		fmt.Println("9. Sair")

		var escolha int
		fmt.Scanln(&escolha)

		switch escolha {
		case 1: // cadastra
			if len(prd.ListaProdutos) < 50 {
				prd.CadastraProduto()
			} else {
				fmt.Println("Limite de 50 produtos atingido!")
			}

		case 2: //visualiza produtos
			prd.ProdutosCadastrados()

		case 3: //busca pelo id
			fmt.Println("Digite o ID do produto:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			produtoID := scanner.Text()
			prd.BuscaProduto(produtoID)

		case 4: // remover um produto
			fmt.Print("Digite o ID do produto que você deseja remover: ")
			var idParaRemover int
			fmt.Scanln(&idParaRemover)
			prd.RemoveProdutoPorID(idParaRemover)

		case 5:
			// remover todos os produtos
			prd.ListaProdutos = [50]prd.Produto{}
			fmt.Println("Todos os produtos foram removido")

		case 6:
			prd.BuscarProdutoNome()

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