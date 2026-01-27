// 1- Object.keys

const User = {
  name: "User Name",
  age: "User Age",
  date: "27/01/26"
}

const userKeys = Object.keys(User)

console.log("Object.keys", userKeys)

// 2- Json.parse oposto de JSON.stringify().

//Obs o elemento que vai receber o parse precisa ser uma string no caso entre "" , '' ou ``
const json_string = `{
  "name": "Cássio",
  "working": "Dev", 
  "year": 2026
}`

const parsed = JSON.parse(json_string)

console.log("Json", parsed)

//3 - JSON.stringify

const json_json = {
 name: "Cássio",
  working: "Dev", 
  year: 2026
}

const stringify_ = JSON.stringify(json_json)

console.log("Json para string", stringify_)