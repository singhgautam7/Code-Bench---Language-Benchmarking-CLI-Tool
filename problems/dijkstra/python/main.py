import sys, heapq
lines = sys.stdin.read().split()
if not lines: sys.exit(0)
V, E = int(lines[0]), int(lines[1])
adj = [[] for _ in range(V)]
idx = 2
for _ in range(E):
    u, v, w = int(lines[idx]), int(lines[idx+1]), int(lines[idx+2]); idx+=3
    adj[u].append((v, w))
dist = [float('inf')]*V
dist[0] = 0
pq = [(0, 0)]
while pq:
    d, u = heapq.heappop(pq)
    if d > dist[u]: continue
    for v, w in adj[u]:
        if dist[u] + w < dist[v]:
            dist[v] = dist[u] + w
            heapq.heappush(pq, (dist[v], v))
