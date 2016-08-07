# File: class_iter.py
# Author: wdeqin
# Time: 2015/5/22 1:25:33
# Purpose: python iterater of class
class People():
    def __init__(self, name, *names):
        self._names = [ name ]
        for n in names:
            self._names.append(n)

    def __iter__(self):
        return iter(self._names)

    def __reversed__(self):
        return reversed(self._names)

def main():
    pp = People('Zhao', 'Qian', 'Sun', 'Li')
    for n in pp:
        print(n, end=' ')
    print()

    for n in reversed(pp):
        print(n, end=' ')
    print()

if __name__ == '__main__':
    main()
