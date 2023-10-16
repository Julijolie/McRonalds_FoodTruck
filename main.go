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

type Pedido struct {
	ID         int
	Delivery   bool
	Produtos   [10]Produto
	ValorTotal float64
	Status     bool
}

var produtos []Produto
var proximoId = 1
var Pedidos []Pedido
var proximoIdPedido = 1
var i = 0
var parada = true
const valorDelivery = 10

func main() {
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Cadastrar produto")
		fmt.Println("2. Visualizar produtos cadastrados")
		fmt.Println("3. Busca do produto pelo id")
		fmt.Println("4. Remover 1 produto")
		fmt.Println("5. Remover todos os produtos")
		fmt.Println("6. Adicionar pedido")
		fmt.Println("7. Expedir pedido")
		fmt.Println("8. Exibir métricas")
		fmt.Println("9. Sair")

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

			case 5:
				// remover todos os produtos
				produtos = nil
				fmt.Println("Todos os produtos foram removido")
			
			case 6:
				novoPedido := Pedido{}
				novoPedido.ID = proximoIdPedido
				proximoIdPedido++
				var id int
				for{
					fmt.Print("Digite o ID do produto que deseja adicionar ao seu pedido, caso queira parar digite -1 : ")
					fmt.Scanln(&id)
					novoPedido.Produtos.(buscaProduto(&id))
					novoPedido.ValorTotal += buscaProduto(&id).ValorTotal
					if id == -1 {
						if novoPedido.Produtos != nil {
							fmt.Print("Digite True se o seu pedido for delivery e false caso contrário : ")
							var delivery bool
							fmt.Scanln(delivery)
							novoPedido.Delivery = delivery
							AdicionaPedido(id, novoPedido)
							Pedidos.append(Pedidos, novoPedido)
							if novoPedido.Delivery == true {
								novoPedido.ValorTotal = novoPedido.ValorTotal + valorDelivery
							} 
							break
						}else{
							break
						}
					}
				}
			//case 7:
			
			case 9:
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
