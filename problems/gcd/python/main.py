import sys
n = int(sys.stdin.read().strip())
def gcd(a, b):
    while b: a, b = b, a % b
    return a
for i in range(1, n+1): gcd(1000000007, i)
