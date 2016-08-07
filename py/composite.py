import itertools
import sys

class Item:
    def __init__(self, name, *item, price=0.0):
        self.__name = name
        self.__price = price
        self.__children = []
        if item:
            self.add(item)

    @property
    def price(self):
        return sum(item.price for item in self) if self.__children else self.__price

    @price.setter
    def price(self, price):
        self.__price = price

    def add(self, first, *rest):
        self.__children.extend(itertools.chain(first, rest))

    def remove(self, item):
        self.__children.remove(item)

    def __iter__(self):
        return iter(self.__children)

    def composite(self):
        return bool(self.__children)

    def print(self, indent="", file=sys.stdout):
        print("%s$%.1f %s" % (indent, self.price, self.__name), file=file)
        if self.__children:
            for item in self.__children:
                item.print(indent=indent + "    ")

def make_item(name, price):
    return Item(name, price=price)

def make_composite(name, *items):
    return Item(name, *items)

def main():
    pencil = Item("Pencil", price=2.0)
    eraser = Item("Eraser", price=1.0)
    ruler = Item("Ruler", price=1.5)
    pencilSuit = Item("Pencil Suit", pencil, eraser, ruler)
    box = Item("Box", price=2.5)
    pencilbox = Item("Pencil Box", pencilSuit, box)
    enbook = Item("English Book", price=12.5)
    backbag = Item("Back Bag", price=27)
    tomsbag = Item("Tom's Bag", backbag, enbook, pencilbox)
    tomsbag.print()

if __name__ == '__main__':
    main()
