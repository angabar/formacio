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
