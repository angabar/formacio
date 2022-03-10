class Dog():
    # Podemos poner todos los atributos que queramos antes del constructor
    species = "mammal"

    def __init__(self, speed):
        self.speed = speed

    # Con self accedemos a las propiedades de la clase, cuidado con Python, las
    # funciones con argumentos deben recibir dichos argumentos o tendremos un
    # error
    def bark(self, number):
        print(f"My speed is {self.speed}")
        print(f"My number is {number}")


my_dog = Dog(speed=2)
print(my_dog.species)
my_dog.bark(number=1)


class Circle():
    pi = 3.1415

    # Para evitar errores podemos definir valores por defecto en los argumentos
    def __init__(self, radius=1):
        self.radius = radius

        # No hace falta que los atributos tengan que recibir valores externos,
        # pueden inicializarse con valores internos
        self.area = radius * radius * self.pi

    def get_circumference(self):
        return self.radius * self.pi * 2


my_circle = Circle(radius=23)
print(my_circle.area)
print(my_circle.get_circumference())


class Animal():
    def __init__(self):
        print("Animal created")

    def who_am_i(self):
        print("I am an animal")

    def eat(self):
        print("I am eating")


my_animal = Animal()
my_animal.who_am_i()
my_animal.eat()


# Pasando la clase como argumento de la clase hija es como se hace la herencia
# en Python
class Cat(Animal):
    def __init__(self):
        # Si queremos inicializar el init del padre (no es obligatorio para
        # acceder a sus metodos) tenemos que invocarlo asi Padre.__init__(self)
        Animal.__init__(self)
        print("Cat created")

    # Para sobreescribir un metodo de la clase padre tan solo tenemos que usar
    # el mismo nombre
    def who_am_i(self):
        print("I am a cat")

    def burk(self):
        print("miau")


my_cat = Cat()
my_cat.who_am_i()
my_cat.burk()


# POLIMORFISMO


class Horse():
    def __init__(self, name):
        self.name = name

    def speak(self):
        return f"{self.name} says wolf"


class Foca():
    def __init__(self, name):
        self.name = name

    def speak(self):
        return f"{self.name} says oioioi"


niko = Horse(name="niko")
felix = Foca(name="felix")

print(niko.speak())
print(felix.speak())

for pet in [niko, felix]:
    print(type(pet))
    print(type(pet.speak()))


# Las clases abstractas como tal no existen en Python, lo que podemos hacer es
# definir los metodos que queremos establecer como contrato y lanzar un error si
# alguien intenta instanciarla y ejecutar dichos metodos
class Animal():
    def __init__(self, name):
        self.name = name

    def speak(self):
        raise NotImplementedError("Subclass must implement this method")


my_second_animal = Animal(name="placa")
print(my_second_animal.speak())
