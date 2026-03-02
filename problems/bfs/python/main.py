import sys
from collections import deque
lines = sys.stdin.read().split()
if not lines: sys.exit(0)
V, E = int(lines[0]), int(lines[1])
adj = [[] for _ in range(V)]
idx = 2
for _ in range(E):
    u, v = int(lines[idx]), int(lines[idx+1]); idx+=3
    adj[u].append(v)
q = deque([0])
vis = [False]*V
vis[0] = True
count = 0
while q:
    curr = q.popleft()
    count+=1
    for nxt in adj[curr]:
        if not vis[nxt]: vis[nxt]=True; q.append(nxt)
