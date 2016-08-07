from collections import deque
from collections import defaultdict
from collections import OrderedDict
import heapq

def main() :
    q = deque(maxlen=5)
    for i in range(1, 100, 3) :
        q.append(i)

    print(q)

    l = []
    for i in range(10, 1, -3) :
        heapq.heappush(l, (i, i * i, i + i))

    print(l)

    d = defaultdict(list)
    d['a'].append(1)
    d['a'].append(2)
    d['a'].append(3)
    d['b'].append(1)
    d['a'].append(4)
    d['c'].append(3)
    print(d)

    o = OrderedDict()
    o['d'] = 4
    o['a'] = 1
    o['c'] = 3
    o['b'] = 2
    o['f'] = 5

    print(o)

    r = dict()
    r['f'] = 1
    r['e'] = 2
    r['d'] = 3
    r['c'] = 4
    print(r)


if __name__ == "__main__" :
    main()
