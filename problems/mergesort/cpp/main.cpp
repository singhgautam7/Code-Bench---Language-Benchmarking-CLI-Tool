#include <iostream>
#include <vector>
using namespace std;
vector<int> mergesort(vector<int> arr) {
    if(arr.size() <= 1) return arr;
    int mid = arr.size() / 2;
    vector<int> L = mergesort(vector<int>(arr.begin(), arr.begin()+mid));
    vector<int> R = mergesort(vector<int>(arr.begin()+mid, arr.end()));
    vector<int> res; res.reserve(arr.size());
    int i=0, j=0;
    while(i < L.size() && j < R.size()) {
        if(L[i] < R[j]) res.push_back(L[i++]);
        else res.push_back(R[j++]);
    }
    while(i < L.size()) res.push_back(L[i++]);
    while(j < R.size()) res.push_back(R[j++]);
    return res;
}
int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    vector<int> nums;
    int n;
    while (cin >> n) nums.push_back(n);
    mergesort(nums);
    return 0;
}