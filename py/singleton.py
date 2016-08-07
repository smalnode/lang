def get(refresh=False):
    if refresh :
        get.data = 0
    if get.data != 0:
        return get.data
    get.data +=1
    return get.data

get.data = 0

def main():
    print("get().data = %d" % (get(True)))
    print("get().data = %d" % (get(True)))
    print("get().data = %d" % (get(True)))
    print("get().data = %d" % (get(True)))

if __name__ == '__main__':
    main()
