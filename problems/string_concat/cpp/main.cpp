#include <iostream>
#include <string>
using namespace std;
int main() {
    string c; int n;
    if(cin >> c >> n) {
        string s = ""; s.reserve(n);
        for(int i=0; i<n; i++) s += c;
    }
    return 0;
}