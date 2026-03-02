#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

void quicksort(vector<int>& arr, int left, int right) {
    if (left >= right) return;
    int pivot = arr[left + (right - left) / 2];
    int i = left, j = right;
    while (i <= j) {
        while (arr[i] < pivot) i++;
        while (arr[j] > pivot) j--;
        if (i <= j) {
            swap(arr[i], arr[j]);
            i++; j--;
        }
    }
    quicksort(arr, left, j);
    quicksort(arr, i, right);
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    vector<int> nums;
    int n;
    while (cin >> n) nums.push_back(n);
    if (!nums.empty()) quicksort(nums, 0, nums.size() - 1);
    return 0;
}
