from curses.ascii import isalpha
import math
import string


def vol(radius):
    return (4 / 3) * math.pi * radius**3


print(vol(2))


def ran_check(num, low, high):
    return num in range(low, high + 1)


print(ran_check(3, 1, 10))


def up_low(text):
    lower_count = 0
    upper_count = 0
    for character in text:
        if character.islower() and character.isalpha():
            lower_count = lower_count + 1
        if character.isupper() and character.isalpha():
            upper_count = upper_count + 1
    print("Lower: ", lower_count)
    print("Upper", upper_count)


s = "Hello Mr. Rogers, how are you this fine Tuesday?"
print(up_low(s))


def unique_list(list):
    unique_list = []
    for element in list:
        if element not in unique_list:
            unique_list.append(element)
    return unique_list


print(unique_list([1, 1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 5]))


def multiply(numbers):
    result = 1
    for number in numbers:
        result = result * number
    return result


print(multiply([1, 2, 3, -4]))


def palindrome(word):
    return word.replace(" ", "") == word[::-1].replace(" ", "")


print(palindrome('helleh'))
print(palindrome('hello'))


def ispangram(word):
    distinct_characters = []
    total_characters = string.ascii_lowercase
    for text_char in word:
        if text_char.lower() not in distinct_characters and text_char.isalpha():
            distinct_characters.append(text_char.lower())
    return len(distinct_characters) == len(total_characters)


print(ispangram('jajaja'))
print(ispangram("The quick brown fox jumps over the lazy dog"))
