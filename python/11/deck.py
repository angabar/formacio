from card import suits, ranks, Card
from random import shuffle
from typing import List


class Deck:
    __all_cards: List[Card] = []

    def __init__(self) -> None:
        for suit in suits:
            for rank in ranks:
                created_card = Card(suit=suit, rank=rank)
                self.__all_cards.append(created_card)

    def shuffle(self) -> None:
        shuffle(self.__all_cards)

    def deal_one(self) -> Card:
        return self.__all_cards.pop()

    def get_all_cards(self) -> List[Card]:
        return self.__all_cards
