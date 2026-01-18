package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type Users struct{
	ID int
	Name string
}

func main() {
	data := []int{1, 2, 3, 4}
	result := mapFunc(data)
	fmt.Println(result)

	parRes := filter(data)
	fmt.Println(parRes)
  
	arrStr := []string{"item 1", "item 2"}
	addAndRemoveItem(arrStr)

	users := []Users{{1, "Ana"}, {2, "João"}, {3, "Tina"}}

	user, index, found := findAndFindindex(users)

	if found {
	fmt.Println(user, index)
 }

 arrWord := []string{"Hello", "Tchê"}

 word_1, word_2 := joinsArr(arrWord)

 fmt.Appendln(word_1, word_2)

 arrLetter := []string{"A", "B", "C", "D"}

 res_1, res_2 :=  slicesfunc(arrLetter)

 fmt.Appendln([]byte(res_1[]), res_2)

 arrSort := []int{15, 84, 1, 99, 10, 54, 20}

 sort_1, sort_2 := sortFunc(arrSort)

 fmt.Appendln(sort_1, sort_2)
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

/*
 arr := []int{0,1,2,3,4,5,6,7,8,9}
  total := 0

for _, n := range arr {
    total += n
}

*/


// append e algo similar a pop

func addAndRemoveItem(data []string) {
	fmt.Println(data)

	data = append(data, "Item colocado pelo append")

	fmt.Println(data)

	data = data[:len(data)-1]
  
	fmt.Println(data)
}

// find e findindex

func findAndFindindex (usersData []Users) (Users, int, bool) {
	index := -1
	var user Users

	for ind, us := range usersData {
		if us.Name == "Ana"{
			user = us
			index = ind

			return user, index , true
		}
	}

	return user, index, false
}

// includes loop ou slice Contains

func includeTen (data []int) bool {
	for _, num := range data{
		if num == 10{
			return   true
		}
	}

	return false
} 

func includeTenSlice(data []int) bool{
	return slices.Contains(data, 10)
}


// join

func joinsArr (arrWord []string) (string, string){
   resJoin :=""

	 resJoinStringMethod :=""

	 for i, v := range arrWord{
		if i > 0 {
			resJoin += " "
		}
		 resJoin += v
	 }

	 resJoinStringMethod = strings.Join(arrWord, " ")

	 return resJoin,  resJoinStringMethod
}

// slice 

func slicesfunc (data []string) ([]string, []string) {
  copy := append([]string{}, data[1:3]...)

	part := data[1:3]

	return  copy, part
}

// sort

//obs: copy 
/**
 Para criar uma copiar preciso:
 1- criar um novo slice para alocar espaço em memoria
 arrOrd = make([]int, len(data))
 2- fazer o copy
 copy(arrOrd, data)
*/

func sortFunc (data []int) ([]int, []int){
	arrOrd := make([]int, len(data))
	copy(arrOrd, data)

	for i := 0; i < len(arrOrd); i++ {
		for j := i + 1; j < len(arrOrd); j++{
			if arrOrd[i] > arrOrd[j]{
				arrOrd[i], arrOrd[j] = arrOrd[j], arrOrd[i]
			}
		}
	}

	sort.Ints(data)
  
	//reverse inverte para a ordem ficar do maior para menor
	//sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	return  arrOrd, data
}