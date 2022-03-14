from cmath import pi, sqrt


class Line():
    __coord1 = 0
    __coord2 = 0

    def __init__(self, coord1, coord2):
        self.__coord1 = coord1
        self.__coord2 = coord2

    def get_distance(self):
        return sqrt(pow((self.__coord2[1] - self.__coord1[1]), 2) + pow((self.__coord2[0] - self.__coord1[0]), 2))

    def get_slope(self):
        return (self.__coord2[1] - self.__coord1[1]) / (self.__coord2[0] - self.__coord1[0])


coordinate1 = (3, 2)
coordinate2 = (8, 10)

li = Line(coord1=coordinate1, coord2=coordinate2)

print(li.get_distance())
print(li.get_slope())


class Cylinder():
    __height = 1
    __radius = 1

    def __init__(self, height, radius):
        self.__height = height
        self.__radius = radius

    def get_volume(self):
        return pi * pow(self.__radius, 2) * self.__height

    def surface_area(self):
        return 2 * pi * self.__radius * (self.__radius + self.__height)


c = Cylinder(2, 3)

print(c.get_volume())
print(c.surface_area())

# Una de las cualidades mas importantes de las tuplas es que pueden destruirse y
# guardarse por separado
x1, x2 = (200, 100)
print(x1)
print(x2)
