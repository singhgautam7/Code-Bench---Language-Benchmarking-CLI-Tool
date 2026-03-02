import sys
def ms(arr):
    if len(arr) <= 1: return arr
    mid = len(arr)//2
    L = ms(arr[:mid])
    R = ms(arr[mid:])
    res, i, j = [], 0, 0
    while i < len(L) and j < len(R):
        if L[i] < R[j]:
            res.append(L[i])
            i+=1
        else:
            res.append(R[j])
            j+=1
    res.extend(L[i:])
    res.extend(R[j:])
    return res

if __name__ == "__main__":
    nums = list(map(int, sys.stdin.read().split()))
    ms(nums)
