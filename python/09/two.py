import one

print("TOP level in two.py")

one.my_function()

if __name__ == "__main__":
    print("TWO is being run directly")
else:
    print("two has been imported")
