#include <iostream>
#include <string>
#include <vector>
using namespace std;
int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    vector<string> arr; string s;
    while(cin >> s) arr.push_back(s);
    if(arr.empty()) return 0;
    string target = arr.back();
    for(int i=0; i<arr.size()-1; i++) {
        if(arr[i] == target) break;
    }
    return 0;
}