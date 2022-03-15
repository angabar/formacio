def my_function():
    print("func in one.py")


print("TOP LEVEL IN ONE.PY")

# En python cuando tenemos este if quiere decir que es un bloque de codigo que
# solo se ejecutara cuando el archivo sea ejecutado directamente y no a traves
# de una importacion lo cual nos ayuda a tener aqui lo que no queremos que se
# ejecute al importar el archivo entero
if __name__ == "__main__":
    print("ONE:PY is running directly")
else:
    print("ONE:PY has been imported")

# La estructura tipica donde esto se usa es la de tener los metodos en la parte
# superior del programa y las ejecuciones de los mismos en el name == main de
# esa manera tenemos ejecuciones separadas de las exportaciones
