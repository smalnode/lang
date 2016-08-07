# class property and prevent deleting property
class Person:
    def __init__(self, first_name):
        self._first_name = first_name

    @property
    def first_name(self):
        return self._first_name

    @first_name.setter
    def first_name(self, value):
        if not isinstance(value, str):
            raise TypeError('First name must be string')
        self._first_name = value

    @first_name.deleter
    def first_name(self):
        raise AttributeError('Cannot delete property')

def main():
    p = Person('Wang')
    print(p.first_name)
    print(p._first_name)
    del p._first_name
    # print(p._first_name) # error
    # del p.first_name # error 

if __name__ == '__main__':
    main()
