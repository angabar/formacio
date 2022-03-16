try:
    RESULT = 10 + "10"
except:
    print("Error here!")
# El bloque else se lanza solo si las cosas van bien
else:
    print("Ademas hago esto")
finally:
    print("Me lanzo pase lo que pase")


# Podemos asociar los errores a tipos concretos y usar tantos except como
# queramos
try:
    RESULT = 10 + "10"
except TypeError:
    print("Type error here!")
except:
    print("Error here!")


def ask_for_int():
    while True:
        try:
            RESULT = int(input("please provide a number:\n"))
        except:
            print("Error: That is not a number")
            continue
        else:
            print("Yes thank you")
            break
        finally:
            print("Im always launch")


ask_for_int()
