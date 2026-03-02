tokens = STDIN.read.split.map(&:to_i)
return if tokens.empty?
v, e = tokens[0], tokens[1]
adj = Array.new(v) { [] }
idx = 2
e.times do
  u,nxt,w = tokens[idx], tokens[idx+1], tokens[idx+2]
  adj[u] << [nxt, w]; idx += 3
end
dist = Array.new(v, 1000000000)
dist[0] = 0
# Ruby lacks native Priority Queue
q = [[0, 0]]
while !q.empty?
  q.sort_by! { |d, n| d }
  d, curr = q.shift
  next if d > dist[curr]
  adj[curr].each do |nxt, w|
    if dist[curr] + w < dist[nxt]
      dist[nxt] = dist[curr] + w
      q << [dist[nxt], nxt]
    end
  end
end
