package extra

import (
	"fmt"
	"os"
	"strconv"
	"encoding/csv"
	prd "McRonald/produto"
)

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
		novoProduto := prd.Produto{}

		novoProduto.ID = prd.ProximoId
		prd.ProximoId++

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
		prd.AdicionaProdutoLista(novoProduto)
	}
}


/*
func MediaTempoExpedicao(){
	fmt.Print(float64(tempoExpedidoTotal)/float64(totalPedExpedidos))
}
*/
