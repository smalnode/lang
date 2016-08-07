import math

def nat(bound):
    x = 1
    y = 2
    while y < bound:
        yield x
        x += y
        y += 1


def main():
    prev = 0.
    for i in nat(10000):
        cur = math.log(i)
        print cur - prev
        prev = cur

if __name__ == '__main__':
    main()
