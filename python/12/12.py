from typing import Callable

# Decorators, un decorator tiene la siguiente estructura


def my_decorator(func) -> Callable:
    def wrapper():
        print("Some code before")
        func()
        print("Some code after")
    return wrapper


@my_decorator
def sum() -> int:
    return 2


sum()
