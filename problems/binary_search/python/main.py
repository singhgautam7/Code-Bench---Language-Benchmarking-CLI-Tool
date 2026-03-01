import sys


def binary_search(arr, target):
    """Perform binary search on a sorted array. Returns the index or -1."""
    low, high = 0, len(arr) - 1
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1
    return -1


def main():
    data = sys.stdin.read().split()
    if len(data) < 2:
        print("Error: expected array and target on stdin", file=sys.stderr)
        sys.exit(1)

    # Last element is the target, rest are the array
    *arr_strs, target_str = data
    target_str = data[-1]
    arr = list(map(int, data[:-1]))
    target = int(target_str)

    result = binary_search(arr, target)
    print(result)


if __name__ == "__main__":
    main()
