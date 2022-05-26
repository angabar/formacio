from typing import List, Union
from card import Card


class Player:
    __name: str = ""
    __all_cards: List[Card] = []

    def __init__(self, name) -> None:
        self.__name = name

    def remove_card(self) -> None:
        # Es importante especificar 0 aqui porque queremos borrar una carta
        # desde el principio, no desde el final de la lista
        return self.__all_cards.pop(0)

    def add_cards(self, new_cards: Union[List[Card], Card]) -> None:
        # En caso de ser un listado de cartas usamos extend en caso de ser una,
        # append
        if type(new_cards) == type([]):
            self.__all_cards.extend(new_cards)
        else:
            self.__all_cards.append(new_cards)

    def get_all_cards(self) -> List[Card]:
        return self.__all_cards

    def __str__(self) -> str:
        return f"Player {self.__name} has {len(self.__all_cards)} cards"
