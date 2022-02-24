def lesser_of_two_evens(a, b):
    if a % 2 == 0 and b % 2 == 0:
        return min([a, b])
    else:
        return max([a, b])


print(lesser_of_two_evens(2, 4))
print(lesser_of_two_evens(2, 5))


def animal_crackers(text):
    splitted_text = text.split()
    return splitted_text[0][0] == splitted_text[1][0]


print(animal_crackers('Levelheaded Llama'))
print(animal_crackers('Crazy Kangaroo'))


def makes_twenty(num1, num2):
    return sum([num1, num2]) == 20 or num1 == 20 or num2 == 20


print(makes_twenty(20, 10))
print(makes_twenty(12, 8))
print(makes_twenty(2, 3))


def old_macdonald(text):
    converted_text = []
    for index in range(len(text)):
        if index == 0 or index == 3:
            converted_text.append(text[index].upper())
        else:
            converted_text.append(text[index])
    return "".join(converted_text)


print(old_macdonald('macdonald'))


def master_yoda(text):
    reversed_sentences = text.split()
    reversed_sentences.reverse()
    return " ".join(reversed_sentences)


print(master_yoda('I am home'))
print(master_yoda('We are ready'))


def almost_there(number):
    # Como se puede ver, el in sirve tambien para ver si un elemento forma parte
    # de un listado de elementos
    return number in range(90, 110) or number in range(190, 210)


print(almost_there(104))
print(almost_there(150))
print(almost_there(209))


def has_33(number_list):
    condition = False
    for index in range(len(number_list)):
        if index > 0 and number_list[index - 1] == number_list[index]:
            condition = True
    return condition


print(has_33([1, 3, 3]))
print(has_33([1, 3, 1, 3]))
print(has_33([3, 1, 3]))
