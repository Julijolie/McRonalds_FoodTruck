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
	delivery bool
}

var idPedido = 1

func AdicionaPedido(id int,novoPedido Pedido) {
	novoPedido.ID = idPedido
	idPedido++

	var resposta string
	fmt.Println("Pedido é para Delivery? (S/N)")
	fmt.Scanln(&resposta)

	if resposta == "S"{
		novoPedido.delivery = true
		novoPedido.ValorTotal += 10.00
	}else{
		novoPedido.delivery = false
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
	
	fmt.Println("Pedido adicionado com sucesso!")

}

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