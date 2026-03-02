import sys
n = int(sys.stdin.read().strip())
sieve = [True] * (n + 1)
count = 0
for p in range(2, n + 1):
    if sieve[p]:
        count += 1
        for i in range(p * p, n + 1, p): sieve[i] = False
