# Perform binary search on a sorted array.
# Returns the index of target, or -1 if not found.
def binary_search(arr, target)
  low = 0
  high = arr.length - 1

  while low <= high
    mid = low + (high - low) / 2
    if arr[mid] == target
      return mid
    elsif arr[mid] < target
      low = mid + 1
    else
      high = mid - 1
    end
  end

  -1
end

data = $stdin.read.split
if data.length < 2
  $stderr.puts "Error: expected array and target on stdin"
  exit 1
end

# Last element is the target, rest are the array
target = data.pop.to_i
arr = data.map(&:to_i)

result = binary_search(arr, target)
puts result
