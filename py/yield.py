def onlyeven(ns) :
    for n in ns :
        if n % 2 == 0 :
            yield n

def main() :
    ns = [i for i in range(10)]
    print(ns)
    ns = [i for i in onlyeven(ns)]
    print(ns)
    ns = [i for i in ns if i > 4]
    print(ns)

if __name__ == '__main__' :
    main()
