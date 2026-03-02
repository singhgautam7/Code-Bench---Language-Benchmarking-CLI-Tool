#include <iostream>
#include <vector>
#include <queue>
using namespace std;
int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int V, E;
    if(cin >> V >> E) {
        vector<vector<int>> adj(V);
        for(int i=0; i<E; i++) {
            int u, v, w; cin>>u>>v>>w;
            adj[u].push_back(v);
        }
        queue<int> q; q.push(0);
        vector<bool> vis(V, false); vis[0] = true;
        while(!q.empty()) {
            int curr = q.front(); q.pop();
            for(int nxt : adj[curr]) {
                if(!vis[nxt]) { vis[nxt]=true; q.push(nxt); }
            }
        }
    }
    return 0;
}