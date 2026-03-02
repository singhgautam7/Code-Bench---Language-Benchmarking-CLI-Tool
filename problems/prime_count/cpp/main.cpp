#include <iostream>
#include <vector>
using namespace std;
int main() {
    int n; if(cin >> n && n>=2) {
        vector<bool> sieve(n+1, true);
        int count = 0;
        for(long long p=2; p<=n; p++) {
            if(sieve[p]) {
                count++;
                for(long long i=p*p; i<=n; i+=p) sieve[i] = false;
            }
        }
    }
    return 0;
}