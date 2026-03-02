import sys
# Increase recursion limit
sys.setrecursionlimit(500000)
lines = sys.stdin.read().split()
if not lines: sys.exit(0)
V, E = int(lines[0]), int(lines[1])
adj = [[] for _ in range(V)]
idx = 2
for _ in range(E):
    u, v = int(lines[idx]), int(lines[idx+1]); idx+=3
    adj[u].append(v)
vis = [False]*V
def dfs(u):
    vis[u] = True
    for v in adj[u]:
        if not vis[v]: dfs(v)
dfs(0)
