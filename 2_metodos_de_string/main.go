package main

import (
	"fmt"
	"strings"
)


func main() {

	split_1, split_2 := splitFunc("s,p,l,i,t,S,t,r,i,n,g", "splitString")

	fmt.Appendln([]byte(split_1[]), split_2)

	overLord := "Over Lord"
  
	conts_1, conts_2 := containFunc(overLord)

	fmt.println(conts_1, conts_2)

	wordForc := "C_s_"

	replace, replaceAll := replaceFunc(wordForc)

	fmt.Println(replace, replaceAll)

	trim := trimFunc("  com espaço para fazer trim  ")

	fmt.Println(trim)

	substring := substringFunc("Removi o resto como substring")

	fmt.Println(substring)

	indexOf := indexOfFunc("abcde")

	fmt.Println(indexOf)

}

func splitFunc(wordSemicol,word string) ([]string, []string) {
  wordSem := strings.Split(wordSemicol, ",") //[]string{"s","p","l","i","t","S","t","r","i","n","g"}
   
	wordsp := strings.Split(word, "") //[]string{"s","p","l","i","t","S","t","r","i","n","g"}
  
	return  wordSem, wordsp
}

func containFunc(word string) (bool, bool) {
	include_d := strings.Contains(word, "o")
	include_over := strings.Contains(word, "Over") 
	
	return include_d, include_over
}

func replaceFunc (word string) (string, string){
	forca := strings.Replace(word, "_", "a", 1)

	forca_all := strings.ReplaceAll(word, "_", "a")

	return forca, forca_all
}

// upper/ lower 

func upperLowerfunc (upperToLower, lowerToUpper, formName string) (string, string, string){
	up := strings.ToLower(upperToLower)

	lw := strings.ToUpper(lowerToUpper)

	name := strings.ToUpper(formName[:1]) + formName[1:] //formName = "makoto" retorna Makoto

	return up, lw, 
}


// trim em go strings.TrimSpace

func trimFunc (data string) string {
	return strings.TrimSpace(data)
}

// go não tem substring mas posso usar slice [:]

/**
  Em Go, string é uma sequência de bytes, não de letras.
  Caracteres acentuados usam mais de 1 byte.
  rune representa um caractere Unicode.
  Converter para []rune evita cortar letras no meio.
*/

/*
  Explicando rune:

	 go enxerga strings como bytes ex: string ação = [97 195 167 195 163 111]

	 quando uso []rune converto de bytes para Unicode ação fica [97 231 227 111]

	 dai string(com a variavel de rune ) = ação

	 s := "ação"

  fmt.Println([]byte(s)) // [97 195 167 195 163 111]
  fmt.Println([]rune(s)) // [97 231 227 111]

  fmt.Println(string([]rune(s))) // "ação"
 */

func substringFunc(data string) (string, string) {
	sub := data[5:12] // apenas seguro para ASCII

	runeData := []rune(data)
	subRune := string(runeData[5:12]) // seguro para Unicode

	return sub, subRune
}


// obs: se data for menos que 12 da erro ok 

//indexOf 

func indexOfFunc(data string) int {
	resIndex := strings.Index(data, "c")

	return resIndex
}

// startsWith() / endsWith()

func startAndEnd(data string) (bool, bool) {
	start := strings.HasPrefix(data, "Full")
	end := strings.HasSuffix(data, "mist")

	return  start, end
}

// charAt()
// text := "makoto"

// fmt.Println(text[0])        // 109 (byte)
// fmt.Println(string(text[0])) // "m"

// r := []rune(text)
// fmt.Println(string(r[0])) // "m"


//repeat

// import "strings"

// repetido := strings.Repeat("ha", 3)
// fmt.Println(repetido) // hahaha
