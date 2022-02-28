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
        if index > 0 and number_list[index - 1] == number_list[index] and number_list[index] == 3:
            condition = True
    return condition


print(has_33([1, 3, 3]))
print(has_33([1, 3, 1, 3]))
print(has_33([3, 1, 3]))


def paper_doll(text):
    new_text = []
    for letter in text:
        for _ in range(0, 3):
            new_text.append(letter)
    return "".join(new_text)


print(paper_doll('Hello'))
print(paper_doll('Mississippi'))


def blackjack(*args):
    if sum(args) <= 21:
        return sum(args)
    if sum(args) > 21 and 11 in args:
        return sum(args) - 10
    else:
        return "BUST"


print(blackjack(5, 6, 7))
print(blackjack(9, 9, 9))
print(blackjack(9, 9, 11))


def summer_69(number_list):
    can_sum = True
    sum = 0
    for num in number_list:
        if num == 6:
            can_sum = False
        if num == 9:
            can_sum = True
            continue
        if can_sum:
            sum = sum + num
    return sum


print(summer_69([1, 3, 5]))
print(summer_69([4, 5, 6, 7, 8, 9]))
print(summer_69([2, 1, 6, 9, 11]))


def spy_game(nums):
    count = 0
    for num in nums:
        if num == 0:
            count = count + 1
        if num == 7 and count == 2:
            count = count + 1
    return count == 3


print(spy_game([1, 2, 4, 0, 0, 7, 5]))
print(spy_game([1, 0, 2, 4, 0, 5, 7]))
print(spy_game([1, 7, 2, 0, 4, 5, 0]))
