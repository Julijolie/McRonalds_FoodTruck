package pedidos

import "McRonalds/produtos"

type Pedido struct {
	ID         int
	Delivery   bool
	Produtos   [10]produtos.Produto
	ValorTotal float64
}