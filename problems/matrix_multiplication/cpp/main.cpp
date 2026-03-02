#include <iostream>
#include <vector>
using namespace std;
int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int n; if(!(cin>>n)) return 0;
    vector<vector<int>> A(n, vector<int>(n));
    vector<vector<int>> B(n, vector<int>(n));
    vector<vector<int>> C(n, vector<int>(n, 0));
    for(int i=0; i<n; i++) for(int j=0; j<n; j++) cin >> A[i][j];
    for(int i=0; i<n; i++) for(int j=0; j<n; j++) cin >> B[i][j];
    for(int i=0; i<n; i++) {
        for(int k=0; k<n; k++) {
            for(int j=0; j<n; j++) C[i][j] += A[i][k] * B[k][j];
        }
    }
    return 0;
}