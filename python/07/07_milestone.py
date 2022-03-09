from random import randrange

player_choices = {}
board = ["", "", "", "", "", "", "", "", ""]
turn: None


def display_board(board):
    print(f"{board[0]} | {board[1]} | {board[2]}")
    print(f"{board[3]} | {board[4]} | {board[5]}")
    print(f"{board[6]} | {board[7]} | {board[8]}")


def player_input(first_player):
    choice = ""
    while choice != "X" and choice != "O":
        choice = input(f"Player {first_player} pick X or O\n")

    player_choices["1"] = "X" if first_player == 1 and choice == "X" else "O"
    player_choices["2"] = "X" if first_player == 2 and choice == "X" else "O"


def place_marker(board, marker, position):
    board[position - 1] = marker


def win_check(board, mark):
    win_flag = False
    win_combinations = [[0, 1, 2], [3, 4, 5], [6, 7, 8], [
        0, 3, 6], [1, 4, 7], [2, 5, 8], [0, 4, 8], [2, 4, 6]]
    for win_combo in win_combinations:
        if board[win_combo[0]] == mark and board[win_combo[1]] == mark and board[win_combo[2]] == mark:
            win_flag = True
            break
    return win_flag


def choose_first():
    return randrange(2) + 1


def space_check(board, position):
    return board[position - 1] == ""


def full_board_check(board):
    return "" in board


def player_choice(board):
    choice = 0
    while choice not in range(1, 10):
        choice = input("Select a position on the board\n")
        if choice.isdigit():
            choice = int(choice)
            if not space_check(board, choice):
                print("La ubicación ya está ocupada, selecciona otro valor\n")
                display_board(board)
                choice = 0
    return choice


def replay():
    choice = ""
    while choice != "Y" and choice != "N":
        choice = input("Play again? (Y or N)\n")
    return choice == "Y"


print('Welcome to Tic Tac Toe!')

while True:
    first_player = choose_first()
    turn = first_player
    player_input(first_player)

    while full_board_check(board):
        player_position = player_choice(board)
        marker = player_choices["1"] if turn == 1 else player_choices["2"]
        place_marker(board, marker, player_position)

        if win_check(board, marker) and marker == player_choices["1"]:
            print("Player 1 wins!")
            break
        if win_check(board, marker) and marker == player_choices["2"]:
            print("Player 2 wins!")
            break

        display_board(board)
        turn = 2 if turn == 1 else 1

    if not replay():
        break
    else:
        player_choices = {}
        board = ["", "", "", "", "", "", "", "", ""]
        turn: None
