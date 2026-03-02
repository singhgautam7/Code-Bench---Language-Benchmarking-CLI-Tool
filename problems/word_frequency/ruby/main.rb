counts = Hash.new(0)
STDIN.read.split.each { |w| counts[w] += 1 }
