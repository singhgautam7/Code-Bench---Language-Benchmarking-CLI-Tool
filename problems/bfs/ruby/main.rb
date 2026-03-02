tokens = STDIN.read.split.map(&:to_i)
return if tokens.empty?
v, e = tokens[0], tokens[1]
adj = Array.new(v) { [] }
idx = 2
e.times do
  u,nxt,w = tokens[idx], tokens[idx+1], tokens[idx+2]
  adj[u] << nxt
  idx += 3
end
q = [0]
vis = Array.new(v, false)
vis[0] = true
head = 0
while head < q.size
  curr = q[head]
  head += 1
  adj[curr].each do |nxt|
    unless vis[nxt]
      vis[nxt] = true
      q << nxt
    end
  end
end
