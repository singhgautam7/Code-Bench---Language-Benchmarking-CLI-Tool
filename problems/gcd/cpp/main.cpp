#include <iostream>
using namespace std;
int gcd(int a, int b) {
    while(b) { int t = b; b = a % b; a = t; }
    return a;
}
int main() {
    int n; if(cin >> n) {
        for(int i=1; i<=n; i++) gcd(1000000007, i);
    }
    return 0;
}