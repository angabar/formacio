from player import Player
from deck import Deck

player_one = Player("One")
player_two = Player("Two")

new_deck = Deck()
new_deck.shuffle()

for x in range(26):
    player_one.add_cards(new_cards=new_deck.deal_one())
    player_two.add_cards(new_cards=new_deck.deal_one())

game_on = True
round_number = 0

while game_on:
    round_number = round_number + 1
    print(f"Current round: {round_number}")

    if len(player_one.get_all_cards()) == 0:
        print("Player one out of cards, Player 2 wins")
        game_on = False
        break
    if len(player_two.get_all_cards()) == 0:
        print("Player one out of cards, Player 2 wins")
        game_on = False
        break

    player_one_cards = []
    player_two_cards = []

    player_one_cards.append(player_one.remove_card())
    player_two_cards.append(player_two.remove_card())

    at_war = True

    while at_war:
        if player_one_cards.get_all_cards()[-1].get_card_value() > player_two_cards.get_all_cards()[-1].get_card_value():
            print("jajaja")
