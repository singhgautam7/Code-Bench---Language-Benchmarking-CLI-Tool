def quicksort(arr)
  return arr if arr.length <= 1
  pivot = arr[arr.length / 2]
  left, right, middle = [], [], []
  arr.each do |x|
    if x < pivot
      left << x
    elsif x > pivot
      right << x
    else
      middle << x
    end
  end
  quicksort(left) + middle + quicksort(right)
end

input = STDIN.read.split.map(&:to_i)
quicksort(input)
