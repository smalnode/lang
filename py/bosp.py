# File: bosp.py
# Author: wdeqin
# Time: 2015/5/22 1:24:33
# Purpose: bitonic Eucliden traversal-saleman problem(CLRs 15-4)

# X = x1, x2, ..., xn (order by x-aix value) 
X = [(1, 5), (2, 1), (3, 9), (4, 1), (5, 9), (6, 1), (7, 5)]
n = len(X)
# d[i, j] = distance(x[i], x[j]), i <= j
d = [0 for i in range(n * n)]

from math import sqrt
def dist(x1, x2):
    x = x1[0] - x2[0]
    y = x1[1] - x2[1]
    return sqrt(x * x + y * y)

def show(arr, n, m):
    for i in range(n):
        for j in range(m):
            print("%.2f" % arr[i * m + j], end='\t')
        print()

for i in range(n):
    for j in range(i + 1, n):
        d[i * n + j] = dist(X[i], X[j])

print("X = ", X)
print("d[n][n] = ")
show(d, n, n)

# B[i][j] the optimal path length from x[j] to x[0] then to x[i], i <= j 
# B[i][j] = B[j][i]
B = [0 for i in range(n * n)]

# P[i][j] the presuccessor of x[j] it the optimal path of B[i][j]
P = [0 for i in range(n * n)]

for i in range(n):
    for j in range(i, n):
        if i == j == 0:
            B[i * n + j] =0
        elif i < j - 1:
            B[i * n + j] = B[i * n + j - 1] + d[(j-1) * n + j]
            P[i * n + j] = j - 1
        else:
            B[i * n + j] = B[0 * n + i] + d[0 * n + j]
            P[i * n + j] = 0
            for k in range (1, j - 1):
                t = B[k * n + i] + d[k * n + j]
                if B[i * n + j] > t:
                    B[i * n + j] = t
                    P[i * n + j] = k

print("B[n][n] = ")
show(B, n, n)
print("P[n][n] = ")
show(P, n, n)

def path(P, n):
    p = []
    #V[i] = True -> x[i] in one half-path of two
    #V[i] = Fale -> x[i] in the other path
    V = [False for i in range(n)]
    V[n - 1] = True
    v = n - 1
    p.append(n - 1)
    while v != 0:
        v = P[v * n + v]
        V[v] = True
        p.append(v)
    for i in range(n):
        if not V[i]:
           p.append(i)
    
    print("Path = ", p)

path(P, n)
