#include <iostream>
#include <string>
#include <unordered_map>
using namespace std;
int main() {
    int n; if(cin >> n) {
        unordered_map<string, int> h;
        h.reserve(n);
        for(int i=0; i<n; i++) h[to_string(i)] = i;
    }
    return 0;
}