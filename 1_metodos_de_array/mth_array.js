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
