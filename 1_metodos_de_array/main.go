package main

import (
	"fmt"
	"slices"
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