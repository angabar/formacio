# La importacion de un archivo simple de python se hace de la siguiente manera
from my_module import my_func

# Para importar de un package tenemos que tener los archivos y funciones dentro
# de carpetas con __init__ a lo que especificamos a python que ese es un
# directorio de un repo de python
from MyMainPackage import some_main_script
from MyMainPackage.SubPackage import my_subscript

my_func()

some_main_script.report()
