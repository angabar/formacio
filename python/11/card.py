from typing import Union


suits = ("Hearts", "Diamonds", "Spades", "Clubs")
ranks = ("Two", "Three", "Four", "Five", "Six", "Seven",
         "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace")
values = {"Two": 2, "Three": 3, "Four": 4, "Five": 5, "Six": 6, "Seven": 7, "Eight": 8,
          "Nine": 9, "Ten": 10, "Jack": 11, "Queen": 12, "King": 13, "Ace": 14, }

# Los parentesis no son necesarios si no vamos a usar herencia


class Card:
    __suit: Union[str, None] = None
    __rank: Union[str, None] = None

    def __init__(self, suit, rank) -> None:
        self.__suit = suit
        self.__rank = rank
        self.__value = values[rank]

    def get_card_value(self) -> str:
        return self.__value

    def __str__(self) -> str:
        return f"{self.__rank} of {self.__suit}"
