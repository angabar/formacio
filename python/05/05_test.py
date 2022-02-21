st = "Print only the words that start with s in this sentence"
list_st = st.split()
for word in list_st:
    if word[0].lower() == "s":
        print(word)

for num in range(0, 11):
    if num % 2 == 0:
        print(num)

print(range(0, 11, 2))

print([x for x in range(1, 51) if x % 3 == 0])

st = "Print every word in this sentence that has an even number of letters"
list_st = st.split()
for word in list_st:
    if len(word) % 2 == 0:
        print("even!")
    else:
        print(word)

for x in range(1, 101):
    if x % 3 == 0 and x % 5 == 0:
        print("FizzBuzz")
    elif x % 3 == 0:
        print("Fizz")
    elif x % 5 == 0:
        print("Buzz")

st = "Create a list of the first letters of every word in this string"
list_st = st.split()

print([word[0] for word in list_st])
