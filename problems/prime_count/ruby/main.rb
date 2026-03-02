n = STDIN.read.to_i
return if n < 2
sieve = Array.new(n+1, true)
count = 0
(2..n).each do |p|
  if sieve[p]
    count += 1
    (p*p).step(n, p) { |i| sieve[i] = false }
  end
end
