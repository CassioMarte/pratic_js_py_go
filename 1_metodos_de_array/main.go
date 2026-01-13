package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4}
	result := mapFunc(data)
	fmt.Println(result)

	parRes := filter(data)
	fmt.Println(parRes)
}


// 1- map go não possui map mais podemos executar assim 
func mapFunc(data []int) []int {
	plusTwo := []int{}

	for _, n := range data {
		plusTwo = append(plusTwo, n*2)
	}

	return plusTwo
}


// 2 go não tem filter

func filter(data []int) []int{
  par := []int{}

	for _, n := range data{
		if n%2 == 0{
			par = append(par, n)
		}
	}

	return par
}
