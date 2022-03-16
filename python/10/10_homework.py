from cmath import sqrt


try:
    for i in ["a", "b", "c"]:
        print(i**2)
except TypeError:
    print("Type error occurs!")
except:
    print("Undefined error occurs!")

try:
    x = 5
    y = 0
    z = x/y
except ZeroDivisionError:
    print("You cannot divide by 0!")
except:
    print("Undefined error occurs!")
finally:
    print("All Done")

while True:
    try:
        number = int(input("Insert a number:\n"))
        print(sqrt(number))
    except TypeError:
        print("Type error occurs!")
    except:
        print("Undefined error occurs!")
    else:
        break
