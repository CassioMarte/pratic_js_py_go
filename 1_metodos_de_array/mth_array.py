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