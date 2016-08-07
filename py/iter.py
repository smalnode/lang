def main():
    with open('./iter.py') as file:
        while True:
            line = next(file, None)
            if line is None:
                break
            print(line)

    numbers = [i for i in range(5)]
    it = iter(numbers)
    try:
        print(next(it))
        print(next(it))
        print(next(it))
        print(next(it))
        print(next(it))
        print(next(it))
        print(next(it))
    except StopIteration:
        pass

if __name__ == '__main__':
    main()
