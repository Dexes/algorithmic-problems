from threading import Event
from typing import Callable


class Foo:
    def __init__(self):
        self.done_first = Event()
        self.done_second = Event()

    def first(self, print_first: 'Callable[[], None]') -> None:
        print_first()
        self.done_first.set()

    def second(self, print_second: 'Callable[[], None]') -> None:
        self.done_first.wait()
        print_second()
        self.done_second.set()

    def third(self, print_third: 'Callable[[], None]') -> None:
        self.done_second.wait()
        print_third()
