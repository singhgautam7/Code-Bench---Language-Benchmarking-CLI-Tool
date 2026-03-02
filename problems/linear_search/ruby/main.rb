tokens = STDIN.read.split
return if tokens.empty?
target = tokens.last
(0...tokens.size-1).each { |i| break if tokens[i] == target }
