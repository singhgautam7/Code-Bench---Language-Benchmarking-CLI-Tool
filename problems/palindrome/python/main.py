import sys
print(str(sys.stdin.read().strip() == sys.stdin.read().strip()[::-1]))
