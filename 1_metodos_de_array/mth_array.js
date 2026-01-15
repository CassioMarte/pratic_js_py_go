// 1 - MAP 

const arrMap = [1,2,3,4,5,6,7,8,9]

const arrMapPlusTwo = arrMap.map((n) => n * 2) 

console.log(arrMapPlusTwo) // node mth_array.js


// 2- FILTER   filter decide QUEM → map decide O QUE vira

const arrFilter = Array.from({length: 5}, (_, i)=> i++)

const arrPar = arrFilter.filter(n => n%2 ===0)

console.log(arrPar)


// 3 reduce

const arrReduce = Array.from({length: 10}, (_, i)=> i++)
                                    //acc = acumulador e now = atual  
const reduceData = arrReduce.reduce((acc, now)=> acc + now, 0)

console.log(reduceData)


//4-  foreach 

const arrFore = ['Oi', "tchê", "bah", ]

arrFore.forEach((wd)=> console.log(wd))

// pop e push 

const arrNew = ["item 1", "item 2", "item 3"]

console.log("arr original", arrNew)

arrNew.push("item colocado com push")

console.log("após push", arrNew)

arrNew.pop()

console.log('após pop', arrNew)


// find() / findIndex()

const users = [
  { id: 1, name: "Ana" },
  { id: 2, name: "João" },
  { id: 3, name: "Maria" },
]

const userName = "Ana"

const getUser = users.find(user => user.name === userName)

console.log("getUser", getUser)

const getUserPosition = users.findIndex(user => user.name === userName)

console.log("getUserPosition", getUserPosition)

// includes()

const arrIncludeTen = Array.from({length: 20}, (_, i)=> i++)

const resInclud = arrIncludeTen.includes(10)

console.log("tem 10?", resInclud)

// join

const arrWord = ["Hello", "Tchê"]

const hi = arrWord.join(" ")

const hi_traco = arrWord.join("-")

console.log(hi, hi_traco)

// slice 

const sliceArr = ["item 1", "item 2", "item 3", "item 4", "item 5"]

const part = sliceArr.slice(1, 4)

console.log("part 1 a 4", part) //part 1 a 4 [ 'item 2', 'item 3', 'item 4' ]

const part_2 = sliceArr.slice(3)

console.log(part_2) //[ 'item 4', 'item 5' ]



//splice  array.splice(inicio, quantidade, ...novosItens)

const spliceArr = ["position 0", "position 1", "position 2", "position 3", "position 4", "position 5", "position 6"]

console.log("arr original", spliceArr)

const removeItensOne_three = spliceArr.splice(1, 4)

console.log("arr original após remove", spliceArr)
console.log("1 a 3", removeItensOne_three)

const add_itens = spliceArr.splice(3, 0, "item_novo_1", "item_novo_2")

console.log("add", add_itens)

console.log("arr após add a partir da terceira posição", spliceArr)

// sort

const arrSort = [15, 84, 1, 99, 10, 54, 20]

const arrOrd = arrSort.sort((a , b)=> a - b)

console.log("ord", arrOrd)