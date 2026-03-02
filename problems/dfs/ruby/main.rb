tokens = STDIN.read.split.map(&:to_i)
return if tokens.empty?
v, e = tokens[0], tokens[1]
$adj = Array.new(v) { [] }
idx = 2
e.times do
  u,nxt,w = tokens[idx], tokens[idx+1], tokens[idx+2]
  $adj[u] << nxt
  idx += 3
end
$vis = Array.new(v, false)
def dfs(u)
  $vis[u] = true
  $adj[u].each { |v| dfs(v) unless $vis[v] }
end
dfs(0) rescue nil
