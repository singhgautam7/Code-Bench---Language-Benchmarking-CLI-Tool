import sys

def qs(arr):
    if len(arr) <= 1: return arr
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return qs(left) + middle + qs(right)

if __name__ == "__main__":
    nums = list(map(int, sys.stdin.read().split()))
    qs(nums)
