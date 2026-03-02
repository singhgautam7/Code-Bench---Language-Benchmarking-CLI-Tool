import sys
lines = sys.stdin.read().split()
if not lines: sys.exit(0)
n = int(lines[0])
idx = 1
A = [[0]*n for _ in range(n)]
B = [[0]*n for _ in range(n)]
for i in range(n):
    for j in range(n):
        A[i][j] = int(lines[idx]); idx+=1
for i in range(n):
    for j in range(n):
        B[i][j] = int(lines[idx]); idx+=1
C = [[0]*n for _ in range(n)]
for i in range(n):
    for j in range(n):
        for k in range(n):
            C[i][j] += A[i][k] * B[k][j]
