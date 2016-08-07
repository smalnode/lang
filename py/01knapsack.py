# File: 01knapsack.py
# Author: wdeqin
# Time: 2015/5/22 1:20:30
# Purpose: dynamic programming - 01 backbag problem

# TODO: v: value, w: weight, c: capacity of backbag
def stole(v, w, c) :
    l = len(v)
    s = [[0 for c in range(c+1)] for r in range(l+1)]
    for i in range(1, l + 1) :
        for j in range(w[i-1], c + 1) :
            s[i][j] = s[i-1][j]
            t = s[i-1][j-w[i-1]] + v[i-1]
            if s[i][j] < t :
                s[i][j] = t
    print([e for e in s])
    return s[l][c]

# TODO: command line entry
def main() :
    v = [60, 100, 120]
    w = [10, 20, 30]
    c = 50
    print(stole(v, w, c))


if __name__ == '__main__' :
    main()
