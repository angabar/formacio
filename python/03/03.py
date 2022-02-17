from statistics import mode


print("hello world")
print(3 % 2)

# This is an exponentian number
print(2 ** 2)

print(0.1 + 0.2 - 0.3)

number = 5

print(number + 10)

number = number + number

print(number)

# type is a function that returns the class of the variable type if we want to
# compare it we need to use "is"
print(type(number))

# type comparsion
print(type(number) is int)

# strings can be encoded by two or single quotes, but be consistent
name = "placa"

# we can grab characters like array
print(name[0])

# and use negative index position
print(name[-2])

# we can slice a string using [ini:end:step]
print(name[1:])
print(name[1:4])
print(name[::2])

# strings in python are immutable you canno't do this name[0] = "t"

# this is a string concatenation
print("aa" + "bb")

print(name.upper())

# len can check the lenght of a string or list
print(len(name))

# string concatenation
# El +, format o el fstring

test_name = "ajajaja"

# Lo que vaya en el {} es lo que se pondrá en el .format, el orden es importante
print("this is a string {}".format(test_name))

# Si queremos alterar el orden tenemos que especificar el índice en los corchetes
print("The {2} {1} {0}".format("fox", "brown", "quick"))

# O tambien podemos especificar una key por cada valor
print("The {q} {b} {f}".format(f="fox", b="brown", q="quick"))

# Podemos especificar la precision de los decimales agregando el . y una f
print("The result was {r:.3f}".format(r=100/777))

# Tambien podemos tener la interpolacion con f
print(f"Hello {test_name}")

# LIST

my_list = [1, 2, 3]

my_list = ['jaja', 2, 3]

# La longitud de una lista es como en los strings
print(len(my_list))

my_list = ['one', 'two', 'three']

print(my_list[0])
print(my_list[1:])

# Las listas se pueden concatenar con +
another_list = ["jaja", "placa"]

# El resultado es la fusion de ambas listas
print(my_list + another_list)

# Las listas a diferencia de los strings son mutables
my_list[0] = "ONE"

print(my_list)

# append agrega un elemento al final de la lista
my_list.append(6)

# pop elimina el ultimo elemento de la lista
my_list.pop()

# Podemos guardar el elemento borrado
removed_element = my_list.pop()
print(removed_element)

# Si lo que queremos es borrar un elemento determinado, pasamos el índice a pop
my_list.pop(0)

# Las listas se guardan por referencia, modificarlas modifican el valor original
print(my_list)

my_unordered_list = ["a", "f", "b"]

# Es muy importante recalcar que sort no devuelve nada, modifica el objeto
# original por lo que poner my_list = list.sort() es una gilipollez
my_unordered_list.sort()
print(my_unordered_list)

# Igual que sort pero a la inversa, tenemos reverse
my_unordered_list.reverse()
print(my_unordered_list)

# DICTIONARIES

my_dict = {
    "asd": "test",
    "aaa": 6
}

print(my_dict["asd"])

prices = {
    "apple": 2.99,
    "oranges": 1.99,
    "milk": 5.8
}

print(prices["apple"])

# Los diccionarios pueden guardar lo que sea
dict2 = {
    "k1": 2,
    "k2": {
        "inside_key": 2,
    },
    "k3": [1, 2, 3]
}
print(dict2["k3"][0])

# De este modo agregamos un elemento al dictionario
dict2["k4"] = "test"
print(dict2)

# Del mismo modo pero con una clave existente si lo que queremos es actualizar
dict2["k1"] = None
print(dict2)

# Para ver todas las keys o los values usamos lo siguientes metodos
print(dict2.keys())
print(dict2.values())
print(dict2.items())

# TUPLAS

# Son iguales que las listas, pero con la diferencia de que son inmutables, no
# pueden modificarse en lugar de braquets se usan parentesis
tupla_example = (1, 2, 3)
print(type(tupla_example))

# Las tuplas pueden redefinirse por completo como los strings
tupla_example = ("aaa", 1)
print(tupla_example)

# Hay dos metodos en las tuplas que son interesantes, count (que cuenta el
# numero de veces que se repite un elemento en una tupla) e index (que nos dice
# el indice de un elemento en una tupla)
print(tupla_example.count("aaa"))
print(tupla_example.index("aaa"))

# Podmeos acceder a los elemento por indice como en las listas
print(tupla_example[1])

# Lo que no podemos hacer con las tuplas es esto: tupla_example[0] = "bbb"

# SETS
# Los sets son listas desordenadas de elementos unicos
my_set = set()

# Usamos add para añadir elementos a un set
my_set.add(1)

print(my_set)

# Tambien podemos definir los sets de esta forma, pero ojo con poner los
# braquets vacio porque python puede considerarlo como un dicionario, si lo que
# queremos es inicializarlo vacio, lo mejor es poner set()
another_set = {1, 2, 3}
print(another_set)

# Un truco para eliminar duplicados es convertir la lista a un set como aqui
my_list_2 = [1, 1, 1, 2, 2, 2, 3, 3]
distinct_my_list_2 = set(my_list_2)
print(distinct_my_list_2)

# Booleans

# La unica consideracion a tener en cuenta es que se ponen con la
# primera letra en mayuscula

# I/O

# Si el archivo no existe open nos devuelve un error
myfile = open("test.txt")

# Con read podemos leer el archivo que cargamos previamente
print(myfile.read())

# Si intentamos volverlo a leer nos dará un string vacío porque el cursor está
# al final del archivo y hay que volverlo a poner al principio para que vuelva a
# leerse, esto se hace con el metodo seek
print(myfile.read())

myfile.seek(0)

# Uno de los metodos mas utiles es el de readlines, el cual nos permite
# introducir cada linea dentro de un array para poder tener mejor acceso a estas
content_file = myfile.readlines()
print(content_file)

# Lo ideal es que cuando terminemos de realizar las operaciones pertinentes con
# los archivos, lo cerremos
myfile.close()

# Existe una forma mas moderna y segura de acceder a los contenidos de un
# archivo, usando with, con esto cerramos el archivo directamente, es como
# abrirlo, realizar las operaciones que queramos y cerrarlo todo en este bloque
with open("test.txt") as my_other_file:
    content_file = my_other_file.read()

print(content_file)

# Si lo que queremos es escribir tenemos que usar el segundo parametro de la
# funcion open y especificar "a" de append el cual nos envia el cursor al final
# del texto para escribir
with open("test.txt", mode="a") as my_other_file:
    content_file = my_other_file.write("\njajajaja")

# Si lo que queremos es crear un archivo y / o sobreescribir uno existente,
# tenemos que usar el modo "w" de write, en este caso, si ponemos un nombre de
# archivo que no existe python lo creara
with open("logs.txt", mode="w") as logs_file:
    logs_file.write("my log jajajaja")
