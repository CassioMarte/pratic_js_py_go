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