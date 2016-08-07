class Person:
    def __init__(self, name):
        self._name = name

def changeName(person):
    person._name = 'New' + person._name

p = Person('Wang')
changeName(p)
print(p._name)
