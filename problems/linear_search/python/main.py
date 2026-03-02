import sys
lines = sys.stdin.read().split()
if lines:
    target = lines[-1]
    for i in range(len(lines)-1):
        if lines[i] == target: break
