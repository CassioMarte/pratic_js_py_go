//1- split ele separa os elementos onde a um separador que especificamos aqui split(",") ou split("")

const splitString = "splitString"
const splitStringSemicolun = "s,p,l,i,t,S,t,r,i,n,g"

const splitArrLetters = splitString.split("")
const splitStringSemicolunLetters = splitStringSemicolun.split(",")

console.log("split", splitArrLetters)
console.log("splitStringSemicolun", splitStringSemicolunLetters)

/**
 * | Separador   | Resultado                  |
| ----------- | -------------------------- |
| `","`       | Divide onde houver vírgula |
| `"-"`       | Divide onde houver hífen   |
| `" "`       | Divide por espaço          |
| `""`        | Divide letra por letra     |
| inexistente | Retorna a string inteira   |

 */

// 2- includes pesquisar arrInclu.indexOf("OverLord") !== -1

const arrInclu = "Fullmetal Alchimist"

const includes_t = arrInclu.includes("t")

const overLord = "OverLord é top".includes("Over")

const apartir_index_10 = arrInclu.includes("t", 10)

console.log("includes_t", includes_t)

// 3- replace e replaceAll
/**
 * 📌 replace → só a primeira
 * 📌 replaceAll → todas
 */

const forca = "C_sa"

const replace_a = forca.replace("_", "a") 

const forca_2 = "C_s_"

const replace_two_a = forca_2.replaceAll("_", "a")

console.log("forca", replace_a, "forca_2", replace_two_a)

//toUpperCase() / toLowerCase() 

/**
 * | Situação                                  | Use                         |
| ----------------------------------------- | --------------------------- |
| Comparações simples, IDs, validações      | `toUpperCase()`             |
| Texto exibido ao usuário (idioma importa) | `toLocaleUpperCase()`       |
| Aplicações multilíngues                   | `toLocaleUpperCase(locale)` ex:  `toLocaleUpperCase('pt-br')`  |

 */

const low = "anime"

const uperWord = low.toUpperCase()

const uper = "OVER"

const lowWOrd = uper.toLowerCase()

const prota = "makoto"

const name =  prota.charAt(0).toLocaleUpperCase() + prota.slice(1)

console.log("up", uperWord, "low", lowWOrd, "name", name)

// trim()

const wordSpace = " oii eu tenho espaços no começo e fim  "

const notSpace = wordSpace.trim()

console.log("wor", wordSpace.split(""), "sem espeços", notSpace.split(""))

// substring() substring(inicio, fim) / slice()

const substringPhrase =  "Removi o resto como substring"

const resSub = substringPhrase.substring(5, 12)

console.log("sub", resSub)

const slicePhrase = "Usei slice pq aceita dados negativos"

const resSlice =  slicePhrase.slice(5, 12)

console.log("sli", resSlice)

// indexOf() / includes

const letters =  "abcde"

const c = letters.indexOf('c') //indexOf 2

const c_inclu = letters.includes("c") // includes true

console.log("indexOf", c, "includes", c_inclu)


// startsWith() / endsWith()

const file = "studiesData.doc"

const start = file.startsWith("stud")

const end = file.endsWith("doc")

console.log("start stud", start, "termina com doc", end)

// charAt()

const caract = "Fullmetal".charAt(0)

console.log(caract)

// repeat

const repetido = "ha".repeat(3); // "hahaha"