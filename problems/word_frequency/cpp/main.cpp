#include <iostream>
#include <string>
#include <unordered_map>
using namespace std;
int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    unordered_map<string, int> counts;
    string w;
    while(cin >> w) counts[w]++;
    return 0;
}