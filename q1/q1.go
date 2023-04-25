package q1

//Em um dia quente de verão, Pete e seu amigo Billy decidiram comprar uma melancia. Eles escolheram a maior e mais
//saborosa, na opinião deles, e, em seguida, pesaram a fruta nas balanças, obtendo seu peso em quilos. Morrendo de sede,
//correram para casa com a melancia e decidiram dividi-la. No entanto, enfrentaram um problema difícil.
//
//Como grandes fãs de números pares, Pete e Billy querem dividir a melancia de tal forma que cada uma das duas partes pese
//um número par de quilos, sem que as partes necessariamente tenham pesos iguais. Mas os meninos estão extremamente
//cansados e querem começar a refeição o mais rápido possível. Portanto, precisam de ajuda para descobrir se é possível
//dividir a melancia da maneira que desejam. É importante destacar que cada um deles deve receber uma parte de peso
//positivo.
//
//A função deve retornar um valor booleano, indicando se é possível ou não dividir a melancia da forma desejada. Se o peso
//da melancia for menor ou igual a 0, a função deve retornar um erro.

func DivideWatermelon(weight int) (bool, error) {
	// Seu código aqui
	return false, nil
}

package main

import "fmt"

func encontrarSomas(x int) {
	for i := 0; i <= x; i += 2 {
		for j := i; j <= x; j += 2 {
			if i+j == x {
				fmt.Println(i, "+", j, "=", x)
			}
		}
	}
}

func main() {
	var numero int
	fmt.Print("Qual é o valor do peso da melancia?")
	fmt.Scan(&numero)
	if numero <= 0 {
		fmt.Println("Peso inválido.")
	}
	if numero/2 == 1 {
		fmt.Println("Não é possível dividir a melancia da forma desejada")
	}
	for numero/2 == 0 {
		for i := 0; i <= numero; i += 2 {
			for j := i; j <= numero; j += 2 {
				if i+j == numero {
					fmt.Println(i, "+", j, "=", numero)
				}
			}
		}
	}
}
