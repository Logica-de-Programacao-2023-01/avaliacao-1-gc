package q4

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

func ClassifyPrices(prices []int) (int, error) {
	package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual é o valor do primeiro número?")
	fmt.Scan(&x)
	var y int
	fmt.Print("Qual é o valor do segundo número?")
	fmt.Scan(&y)
	var z int
	fmt.Print("Qual é o valor do terceiro número?")
	fmt.Scan(&z)
	var w int
	fmt.Print("Qual é o valor do quarto número?")
	fmt.Scan(&w)
	var m int
	fmt.Print("Qual é o valor do quinto número?")
	fmt.Scan(&m)
	if x<y<z<w<m
		fmt.Println("1")
	if y, z, w, m==0
		fmt.Println("3")
}

	return 0, nil
}
