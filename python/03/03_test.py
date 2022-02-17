print(type(3 + 1.5 + 4))

s = "hello"

print(s[1])
print(s[::-1])

print(s[4])
print(s[-1])

list_1 = [0, 0, 0]
list_2 = [0]*3

list_3 = [1, 2, [3, 4, "hello"]]
list_3[2][2] = "Goodbye"

list_4 = [5, 3, 4, 6, 1]

# sort difiere de sorted en que el primero devuelve None, por lo que no podemos
# usar el resultado en una variable, el segundo devuelve el resultado de la
# lista ordenada, no modifica el original
list_4.sort()
sorted(list_4)

d = {"simple_key": "hello"}
print(d["simple_key"])

d = {"k1": {"k2": "hello"}}
print(d["k1"]["k2"])

d = {"k1": [{"nest_key": ["this is deep", ["hello"]]}]}
print(d["k1"][0]["nest_key"][1][0])

d = {"k1": [1, 2, {"k2": ["this is tricky", {"tough": [1, 2, ["hello"]]}]}]}
print(d["k1"][2]["k2"][1]["tough"][2][0])

# Tupples are inmutable
tuple_example = (1, 2, 3)

# Elements in set are unique and unordered
list_5 = [1, 2, 2, 33, 4, 4, 11, 22, 3, 3, 2]
unique_list_5 = set(list_5)
print(unique_list_5)

# Booleans
# False
# False
# False
# True
# False
# False
