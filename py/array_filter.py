def main():
    x = { 1: 1, '2': 2, '3': '3', 4: 4 }
    y = [ x[e] for e in x.keys() if isinstance(e, int) and isinstance(x[e], int) ]
    print y
    print y - x.keys()

if __name__ == '__main__':
    main()
