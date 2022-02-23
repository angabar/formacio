from random import shuffle


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
