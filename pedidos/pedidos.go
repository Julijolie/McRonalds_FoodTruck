package pedidos

import (
	prd "McRonald/produto"
	"fmt"
	"time"
)


type Pedido struct{
	ID int
	Produtos [10]PedidoProduto
	Data time.Time
	ValorTotal float64
	Delivery bool
}

type PedidoProduto struct{
	Produto prd.Produto
	Quantidade int
}

var idPedido = 1
var ListaPedidos [1000]Pedido
var tamanhoListaPedidos = 0

func AdicionaPedido() {
	novoPedido := Pedido{}
	
	novoPedido.ID = idPedido
	idPedido++

	var resposta string
	fmt.Println("Pedido é para Delivery? (S/N)")
	fmt.Scanln(&resposta)

	switch resposta {
		case "S", "s":
			novoPedido.Delivery = true
			novoPedido.ValorTotal += 10.00
		case "N", "n":
			novoPedido.Delivery = false
		default:
			fmt.Println("Resposta inválida!")
			return
	}

	novaListaProduto := [10]PedidoProduto{}
	i := 0

	for{
		if i == 10{
			break
		}

		prd.ProdutosCadastrados()

		fmt.Println("Qual produto quer adicionar ao pedido? (Digite o ID)")
		var idProduto string
		fmt.Scanln(&idProduto)

		produto := prd.BuscaProduto(idProduto)

		fmt.Println("Quantos produtos quer adicionar ao pedido?")
		var quantidade int
		fmt.Scanln(&quantidade)

		novoPedidoProduto := PedidoProduto{}
		novoPedidoProduto.Produto = produto
		novoPedidoProduto.Quantidade = quantidade

		novaListaProduto[i] = novoPedidoProduto
		i++

		novoPedido.ValorTotal += float64(novoPedidoProduto.Quantidade) * novoPedido.ValorTotal

		fmt.Println("Deseja adicionar mais produtos? (S/N)")
		fmt.Scanln(&resposta)

		if resposta == "N" || resposta == "n"{
			break
		}
	}

	novoPedido.Produtos = novaListaProduto

	novoPedido.Data = time.Now()

	AdicionaPedidoLista(novoPedido)
	
	fmt.Println("Pedido adicionado com sucesso!")

}

func AdicionaPedidoLista(ped Pedido) {
	if tamanhoListaPedidos > 1000 {
		fmt.Println("Limite de 1000 pedidos atingido!")
		return
	}
	ListaPedidos[tamanhoListaPedidos] = ped
	tamanhoListaPedidos++  
}

func PedidosCadastrados(){
	fmt.Println("Pedidos casdastrados:")
	for _, p := range ListaPedidos {
		if p == (Pedido{}) {
			continue
		}
		fmt.Printf(" ID: %d \n Data: %s \n Valor: %.2f \n Delivery: %t \n Produtos: ", p.ID, p.Data.Format("02/01/2006 15:04:05"), p.ValorTotal, p.Delivery, )
		for _, prod := range p.Produtos{
			if prod == (PedidoProduto{}){
				continue
			}
			fmt.Printf("%dx de %s ;", prod.Quantidade, prod.Produto.Nome)
		}
		fmt.Printf("\n")
	}
}


func ExpedirPedido(){
	if tamanhoListaPedidos <= 0{
		fmt.Println("Não há pedidos cadastrados")
		return
	}

	for i, p := range ListaPedidos{
		if p == (Pedido{}){
			continue
		}
		ListaPedidos[i] = Pedido{}
		fmt.Println("Pedido expedido com sucesso!")
		return 
	}
}