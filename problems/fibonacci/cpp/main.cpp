#include <iostream>
using namespace std;
int main() {
    long long n, a=0, b=1, t;
    if(cin >> n) {
        for(long long i=0; i<n; i++) { t=a; a=b; b=t+b; }
    }
    return 0;
}