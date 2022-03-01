from random import shuffle
from re import X


# FUNCTIONS
def say_hello(name="Default"):
    print(f"jajajaj {name}")


say_hello("perico")
say_hello()


def add_num(num1, num2):
    return num1 + num2


result = add_num(2, 3)
print(result)


def is_even(num):
    return num % 2 == 0


print(is_even(2))
print(is_even(1))


def list_with_even_number(list_number):
    is_even = False
    even_numbers = []
    for number in list_number:
        if number % 2 == 0:
            is_even = True
            even_numbers.append(number)
    return {
        "is_even": is_even,
        "even_numbers": even_numbers
    }


print(list_with_even_number([1, 2, 4]))


# Tuple unpacking con funciones
stock_prices = [("apple", 200), ("google", 800)]
for name, price in stock_prices:
    print(name)


def price_check(stock_prices):
    current_max = 0
    company_name = ""

    for name, price in stock_prices:
        if price > current_max:
            current_max = price
            company_name = name

    return (company_name, current_max)


print(price_check(stock_prices))

# Podmeos guardar el resultado de una destruccion en todas las variables que
# queramos, pero cuidado con acceder a elementos que no existen
name_max, price_max = price_check(stock_prices)
print(name_max)

example_list = [1, 2, 3, 4, 5, 6, 7]


# shuffle no devuelve nada, pero podemos hacer una funcion para devolver una
# alteracion de la lista que entre
def shuffle_list(my_list):
    shuffle(my_list)
    return my_list


result = shuffle_list(example_list)

my_game_list = [" ", "0", " "]


def player_guess():
    guess = ""
    while guess not in ["0", "1", "2"]:
        guess = input("pick a number 0, 1 or 2\n")

    return int(guess)


def check_guess(my_list, user_guess):
    if my_list[user_guess] == "0":
        print("Correct")
    else:
        print("Wrong guess")
        print(my_list)


# Creamos la lista y la alteramos, creamos la eleccion del usuario y por ultimo
# lo combinamos todo
shuffled_balls = shuffle_list(my_game_list)
player_guess = player_guess()

check_guess(shuffled_balls, player_guess)

# args and kwargs (arguments and keyword arguments)


def myfunc(num1, num2):
    return (num1 + num2) * 0.05


print(myfunc(40, 60))


# Con args tenemos acceso a todos los parametros de una funcion metidos en un
# tupla
def myargsfunc(*args):
    print(args)
    return sum(args) * 0.05


print(myargsfunc(40, 60))


def mykwargsfunc(**kwargs):
    print(kwargs)


mykwargsfunc(test="placa")


# En python podemos usar dos tipos de argumentos, como en Dart, los
# posicionales, los tipicos, y los asociados a une key, los ultimos se definen
# con el mismo nombre por el cual sera utilizado en la variable
def testfunc(myvalue):
    print(myvalue)


# Cuidado, acceder a una propiedad asociativa tenemos que asegurarnos de que la
# key existe en la funcion o tendremos un error
testfunc(myvalue="test")


# MAP AND FILTER
def add_two(number):
    return number + 2


# Las funciones map o filter permiten actuar sobre elementos de manera
# individual, reciben la función y el listado sobre el que tienen que actuar. Es
# importante tener en cuenta que para poder iterar sobre el resultado debemos
# convertirlo en una lista, ya que por defecto, map o filter devuelven
# direcciones de memoria
print(list(map(add_two, [1, 2, 3, 4])))


# LAMBDA EXPRESION
print(list(map(lambda num: num + 2, [1, 2, 3, 4])))


# El scope de las variables sigue la regla LEGB, local, variables de afuera de
# las funciones, globales, built in
x = 50


def my_func(x):
    # La sentencia global nos permite tener acceso a la variable global dentro
    # de la funcion, a partir de esta linea todas las reasignaciones de x se
    # haran sobre la variable global, no se recomienda
    global x

    x = "abc"
    print(x)
