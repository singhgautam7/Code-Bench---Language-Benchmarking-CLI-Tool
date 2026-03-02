n = STDIN.read.to_i
def gcd(a, b)
  while b != 0; a, b = b, a % b end; a
end
(1..n).each { |i| gcd(1000000007, i) }
