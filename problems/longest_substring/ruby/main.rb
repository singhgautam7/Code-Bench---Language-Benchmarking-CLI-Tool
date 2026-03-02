s = STDIN.read.strip
chars = require 'set'; Set.new
l = 0; res = 0
(0...s.length).each do |r|
  while chars.include?(s[r])
    chars.delete(s[l]); l+=1
  end
  chars.add(s[r])
  res = [res, r-l+1].max
end
