lines = STDIN.read.split.map(&:to_i)
return if lines.empty?
n = lines[0]
idx = 1
a = Array.new(n) { Array.new(n, 0) }
b = Array.new(n) { Array.new(n, 0) }
(0...n).each { |i| (0...n).each { |j| a[i][j] = lines[idx]; idx+=1 } }
(0...n).each { |i| (0...n).each { |j| b[i][j] = lines[idx]; idx+=1 } }
c = Array.new(n) { Array.new(n, 0) }
(0...n).each { |i|
  (0...n).each { |k|
    (0...n).each { |j| c[i][j] += a[i][k] * b[k][j] }
  }
}
