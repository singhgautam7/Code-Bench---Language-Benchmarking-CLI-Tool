#include <iostream>
#include <vector>
using namespace std;
vector<vector<int>> adj;
vector<bool> vis;
void dfs(int u) {
    vis[u] = true;
    for(int v : adj[u]) { if(!vis[v]) dfs(v); }
}
int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int V, E;
    if(cin >> V >> E) {
        adj.resize(V); vis.assign(V, false);
        for(int i=0; i<E; i++) {
            int u, v, w; cin>>u>>v>>w;
            adj[u].push_back(v);
        }
        dfs(0);
    }
    return 0;
}