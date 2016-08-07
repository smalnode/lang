def test_stack():
    s = []
    for i in range(1, 100, 3):
        s.append(i)

    while len(s) > 0:
        print s.pop()

def test_queue():
    s = []
    for i in range(1, 100, 3):
        s.append(i)

    while len(s) > 0:
        print s.pop(0)

def test_set():
    s1 = { 1, 3, 5, 7, 9 }
    s2 = { 1, 5, 7, 11}

    print s1.union(s2)
    print s1 - s2
    print s2 - s1
    print s1.difference(s2)
    print type(type(s1))
    print isinstance(s1, set)

def main():
    test_set()


if __name__ == '__main__':
    main()
