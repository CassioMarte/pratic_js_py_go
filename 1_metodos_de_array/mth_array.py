# 1- map e compression duas formas de fazer a mesmo coisa

arrMap = [1,2,3,4,5,6,7,8,9 ]

arrMapPlusTwo = list(map(lambda n: n * 2, arrMap))

print(arrMapPlusTwo)

arrMapPlusTwoCompre = [n * 2 for n in arrMap]

print(arrMapPlusTwoCompre)



# filter

arrFilter = list(range(5))

arrPar = list(filter(lambda n: n % 2 == 0, arrFilter))

arrParCompres = [n for n in arrFilter if n % 2 == 0]

print(f"arrParFilter, {arrPar} ou arrParCompres {arrParCompres}")

# append e pop 

arrNew = ["item 1", "item 2", "item 3"]

print(f"Array original {arrNew} " )

arrNew.append("Item novo colocado pelo append")

print(f"Array após append {arrNew}")

arrNew.pop()

print(f"Array após pop {arrNew}")

#  find e findIndex não existem em py 
users = [
  { id: 1, "name": "Ana" },
  { id: 2, "name": "João" },
  { id: 3, "name": "Maria" },
]

userName = "Ana"

userFind = next((user for user in users if user["name"] == userName), None)

userFindIndex = next((index for index, user in enumerate(users) if user["name"] == userName), -1)

print(f"{userFind} {userFindIndex}")

# includes 

arrIncludeTen = list(range(20))

resInclud = 10 in arrIncludeTen

print(f"Tem 10 no array {resInclud}")


# join

arrWord = ["Hello", "Tchê"]

hi = " ".join(arrWord)

print(hi)

hi_traco = "-".join(arrWord)

print(hi_traco)

# slice 

arr = ["A", "B", "C", "D"]
part = arr[1:3]


print(part) #['B', 'C']