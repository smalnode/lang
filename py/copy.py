def changeArray(a):
    a[0] = a[0] + 1
def main():
    a = [4, 5, 6]
    b = a[1:]
    b[0] = 0
    if a[1] == b[0]:
        print("b = a[1:] is a reference")
    else:
        print("b = a[1:] is a copy")

    before = a[0]
    changeArray(a)
    after = a[0]
    if before == after:
        print("function pass array as copy")
    else:
        print("function pass array as reference")

if __name__ == "__main__":
    main()
