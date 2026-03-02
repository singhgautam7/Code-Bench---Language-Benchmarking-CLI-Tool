def mergesort(arr)
  return arr if arr.size <= 1
  mid = arr.size / 2
  l = mergesort(arr[0...mid])
  r = mergesort(arr[mid...arr.size])
  res = []
  i, j = 0, 0
  while i < l.size && j < r.size
    if l[i] < r[j]
      res << l[i]; i+=1
    else
      res << r[j]; j+=1
    end
  end
  res + l[i..-1] + r[j..-1]
end
input = STDIN.read.split.map(&:to_i)
mergesort(input) unless input.empty?
