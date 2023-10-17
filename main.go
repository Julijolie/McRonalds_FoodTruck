package main

import (
	"McRonalds/operacoes"
	"McRonalds/produtos"
	"bufio"
	"fmt"
	"os"
)


type Pedido struct {
	ID         int
	Delivery   bool
	Produtos   [10]produtos.Produto
	ValorTotal float64
	Status     bool
}

var listaDeProdutos []produtos.Produto

func main() {
	operacoes.LerProdAdcLista()
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Cadastrar produto")
		fmt.Println("2. Visualizar produtos cadastrados")
		fmt.Println("3. Busca do produto pelo id")
		fmt.Println("4. Remover 1 produto")
		fmt.Println("5. Remover todos os produtos")
		fmt.Println("6. Buscar produto pelo nome")
		fmt.Println("7. Adicionar pedidos")
		fmt.Println("8. Expedir pedidos")
		fmt.Println("9. Exibir métricas")
		fmt.Println("10. Sair")

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
			fmt.Println("Digite o ID do produto:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			produtoID := scanner.Text()
			fmt.Printf("Produto Encontrado:\nID: %d, Nome: %s, Descrição: %s, Valor: %.2f, Quantidade: %d\n", operacoes.BuscaProduto(produtoID).ID, operacoes.BuscaProduto(produtoID).Nome, operacoes.BuscaProduto(produtoID).Descricao, operacoes.BuscaProduto(produtoID).Valor, operacoes.BuscaProduto(produtoID).Quantidade)

		case 4: // remover um produto
			fmt.Print("Digite o ID do produto que você deseja remover: ")
			var idParaRemover int
			fmt.Scanln(&idParaRemover)
			operacoes.RemoveProdutoPorID(idParaRemover)

		case 5:	// remover todos os produtos
			listaDeProdutos = nil
			fmt.Println("Todos os produtos foram removido")

		case 6:
			operacoes.BuscarProdutoNome()
		case 7:
			operacoes.AdicionaPedido()
		case 8:
			operacoes.ExpedirPedido()
		case 9:
			operacoes.ExibirMetricas()
		case 10:
			os.Exit(0)

			default:
				fmt.Println("Opção inválida!")
			}
			fmt.Println("-------------------------------")
		}
	}


