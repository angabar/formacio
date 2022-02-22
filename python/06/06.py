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
