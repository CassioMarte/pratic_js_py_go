# 1 split
splitString = "splitString"
splitStringSemicolun = "s,p,l,i,t,S,t,r,i,n,g"

splitArrLetters = splitString.split()

splitSemicolon = splitStringSemicolun.split(",")

print(f"splitArrLetters {splitArrLetters}, splitSemiColon {splitSemicolon}")

# obs para em uma string pura como splitString = "splitString" para ter o mesmo efeito que 
# split("") do js precisamos usar list

split = list(splitString)

print(f"split {split}")

# include 

arrInclu = "Fullmetal Alchimist"

include_t = "t" in arrInclu 

overLord = "OverLord é top"

overIN = "Over" in overLord

print(f"t in {include_t}, over in {overLord}")


# replace
# aqui como não existe o replace all se for só replace troca todos mas se colocar "old", "new", number vai ser torcado o numero de itens solicitado

forca = "C_s_"

resForca = forca.replace("_", "a", 1)

print(resForca)

# upper / lower / title

low = "anime"

up = low.upper()

uper = "OVER"

lw = uper.lower()

prota = "makoto"

name = prota.capitalize() # prota[0].upper + prota[1:]

titleAnime = "dragon ball z" # Dragon Ball Z

tt = titleAnime.title()

print(f"up {up}, low {lw}, name {name}, titulo {tt}")