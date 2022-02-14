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
