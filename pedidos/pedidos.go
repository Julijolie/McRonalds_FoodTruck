package pedidos

import (
	prd "McRonald/produto"
	"fmt"
	"time"
)


type Pedido struct{
	ID int
	Produtos [10]prd.Produto
	Data time.Time
	ValorTotal float64
	Delivery bool
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

	if resposta == "S"{
		novoPedido.Delivery = true
		novoPedido.ValorTotal += 10.00
	}else{
		novoPedido.Delivery = false
	}

	novaListaProduto := [10]prd.Produto{}
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

		novaListaProduto[i] = produto
		i++

		novoPedido.ValorTotal += float64(quantidade) * novoPedido.Produtos[i].Valor
	}

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
		fmt.Printf("ID: %d, Data: %s, Valor: %.2f, Delivery: %t, Produtos: %v\n", p.ID, p.Data, p.ValorTotal, p.Delivery, p.Produtos)
	}
}

/*

func ExpedirPedido(){
	pedidoExpedido := removePedido() // pedidoExpedido = {1; 2 xburguers, 3 xfrango, 1 xpeixe; 200; 18:00}
	dataExpedido = time.Now // 18:30

	tempoSub = dataExpedido.Sub(pedidoExpedido.Data).Minutes()

	tempoExpedidoTotal = tempoExpedidoTotal + tempoSub
	totalPedExpedidos++
}


*/