#include <iostream>
using namespace std;
// C++ BigInt is non-trivial, wrap around u64 for benchmark purposes
int main() {
    unsigned long long n, s=1;
    if(cin >> n) {
        for(unsigned long long i=1; i<=n; i++) s *= i;
    }
    return 0;
}