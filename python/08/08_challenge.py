class Account():
    __owner = ""
    __balance = 0

    def __init__(self, owner, balance):
        self.__owner = owner
        self.__balance = balance

    def get_owner(self):
        return self.__owner

    def get_balance(self):
        return self.__balance

    def deposit(self, quantity):
        self.__balance = self.__balance + quantity
        print("Deposit accepted")

    def withdraw(self, quantity):
        if quantity > self.__balance:
            return print("Funds Unavailable!")
        self.__balance = self.__balance - quantity
        return print("Withdrawal Accepted")

    def __str__(self):
        return f"Account owner: {self.__owner}\nAccount balance: {self.__balance}"


acct1 = Account("Jose", 100)
print(acct1)
acct1.deposit(50)
acct1.withdraw(75)
acct1.withdraw(500)
print(acct1.get_balance())
