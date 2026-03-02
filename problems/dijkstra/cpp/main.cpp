#include <iostream>
#include <vector>
#include <queue>
using namespace std;
struct Edge { int to, weight; };
int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int V, E;
    if(cin >> V >> E) {
        vector<vector<Edge>> adj(V);
        for(int i=0; i<E; i++) {
            int u, v, w; cin>>u>>v>>w;
            adj[u].push_back({v, w});
        }
        vector<int> dist(V, 1e9); dist[0] = 0;
        priority_queue<pair<int,int>, vector<pair<int,int>>, greater<pair<int,int>>> pq;
        pq.push({0, 0});
        while(!pq.empty()) {
            int d = pq.top().first, u = pq.top().second; pq.pop();
            if(d > dist[u]) continue;
            for(auto& edge : adj[u]) {
                if(dist[u] + edge.weight < dist[edge.to]) {
                    dist[edge.to] = dist[u] + edge.weight;
                    pq.push({dist[edge.to], edge.to});
                }
            }
        }
    }
    return 0;
}