from random import shuffle
from random import randint

# IF

hungry = False

if hungry:
    print("jajajaj")
else:
    print("false!")

location = "bank"

if location == "auto shop":
    print("cars!")
elif location == "bank":
    print("money is cool")
else:
    print("i do not know much")

# FOR LOOP
my_list = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

# La variable item puede tener el nombre que queramos
for item in my_list:
    if item % 2 == 0:
        print(item)

acc = 0

for item in my_list:
    acc = acc + item

print(acc)

for word in "Hello world":
    print(word)

# Es una buena practica poner el underscore si no vamos a utilizar la variable
for _ in "Hello world":
    print("jajaja")

my_tuple = (1, 2, 3, 4)

for item in my_tuple:
    print(item)

my_special_list = [(1, 2), (2, 3), (3, 4), (4, 5)]

for item in my_special_list:
    print(item)

# Esto se llama tuple unpacking y nos sirve para obtener los elementos de las
# tuplas de manera individual, ojo porque si creamos una variable unpacking que
# no existe en la tupla tendremos un error
for a, b in my_special_list:
    print(a)
    print(b)

my_dict = {"k1": 1, "k2": 2, "k3": 3}

# Cuando iteramos sobre un diccionario solo nos devuelve las keys
for item in my_dict:
    print(item)

# Si lo que queremos es obtener los pares clave/valor tenemos que invocar el
# metodo items()
for key, value in my_dict.items():
    print(value)

# Si solo queremos los valores podemos invocar al metodo values()
for value in my_dict.values():
    print(value)

# WHILE LOOP
x = 0

# Los while se pueden combinar con else para los casos en los que la condicion
# no se cumpla
while x < 5:
    x = x + 1
    print(x)
else:
    print("juas")

# break, continue, pass
list_2 = [1, 2, 3]

# pass solo sirve para pasar olimpicamente del codigo que haya, no hace nada mas
for item in list_2:
    print(item)
    pass

my_string = "Sammy"

# continue lo que hace es pasar a la siguiente iteración ignorando lo que venga
# despues
for letter in my_string:
    if letter == "a":
        continue
    print(letter)

# break lo que hace es interrumpir el bucle sin que se ejecute lo que haya
# despues
for letter in my_string:
    if letter == "a":
        break
    print(letter)

# USEFUL OPERATORS

# range es un método que nos devuelve los numeros desde la primera posicion que
# especifiquemos, hasta la ultima (sin incluirla) con un step definido en la
# tercera posicion (identico a los string)
for num in range(10, 30, 3):
    print(num)

# Si queremos usarlo como variable, tenemos que castearla a list antes, por el
# mismo no devuelve nada
nums = list(range(20, 50, 5))
print(nums)

index_count = 0

for letter in "abcde":
    print(f"at index {index_count} the letter is {letter}")
    index_count = index_count + 1

# otra funcion es enumerate, la cual nos devuelve una serie de tuplas con el
# indice y el contenido del mismo en cada item
for index, letter in enumerate("jajajaja"):
    print(index)
    print(letter)
    print("----")

mylist1 = [1, 2, 3, 4, 5]
mylist2 = ["a", "b", "c"]

# zip lo que hace es fusionar dos listas en pares de indices con tuplas, solo
# los que pueda, los elementos extra que no pueda unir los ignorara
for item in zip(mylist1, mylist2):
    print(item)

# Podemos usar in para determinar si un elemento esta en un listado
print("z" in [1, 2])

# ... o en un string
print("a" in "abc")

# ... o en un diccionario
print("mykey" in {"mykey": 1})

mylist3 = [2, 3, 4, 5, 6]
print(min(mylist3))
print(max(mylist3))

# shuffle lo que hace es alterar el orden de los elementos en una lista de
# manera aleatoria, no devuelve nada sino que modifica el elemento original
my_ordered_list = [1, 2, 3, 4]
shuffle(my_ordered_list)
print(my_ordered_list)

random_number = randint(0, 100)
print(random_number)

# input siempre devuelve un string, ojo con las operaciones matematicas
# number_user = input("enter a number")
# print("number", number_user)

my_string = "hello"

# Esta forma de construir los arrays se llama list comprehension
my_update_string = [letter for letter in my_string]
print(my_update_string)

# Podmeos manipular la varibale directamente de esta forma
print([x * 2 for x in range(0, 90)])

# En el cado de que necesitemos definir un if, lo podemos poner en ultima
# instacia
print([x for x in range(0, 11) if x % 2 == 0])

# No obstante, en el caso de que tengamos que definir un else tambien, el orden
# cambiara y estos se definiran al principio (no recomendable por legibilidad)
print([x if x % 2 == 0 else 'ODD' for x in range(0, 100)])

# En el caso de tener que deifinir dos bucles anidados, lo que tenemos que hacer
# es definir estos uno al lado del otro
print([x * y for x in [2, 4, 6] for y in [1, 2, 3]])
