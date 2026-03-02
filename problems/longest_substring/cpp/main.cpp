#include <iostream>
#include <string>
#include <unordered_set>
#include <algorithm>
using namespace std;
int main() {
    string s; cin >> s;
    unordered_set<char> chars;
    int l=0, res=0;
    for(int r=0; r<s.length(); r++) {
        while(chars.count(s[r])) { chars.erase(s[l]); l++; }
        chars.insert(s[r]);
        res = max(res, r-l+1);
    }
    return 0;
}