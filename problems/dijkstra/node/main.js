const fs = require('fs');
let tokens = fs.readFileSync(0, 'utf-8').trim().split(/\s+/);
if(tokens.length === 0 || tokens[0] === "") process.exit(0);
let V = parseInt(tokens[0]), E = parseInt(tokens[1]);
let adj = Array.from({length: V}, () => []);
let idx = 2;
for(let i=0; i<E; i++) {
    let u = parseInt(tokens[idx++]), v = parseInt(tokens[idx++]), w = parseInt(tokens[idx++]);
    adj[u].push({to: v, w: w});
}
let dist = new Int32Array(V).fill(1e9); dist[0] = 0;
// NodeJS lacks a native Priority Queue. We simulate with an array sort for benchmark overhead purposes, understanding it alters true O bounds slightly.
let pq = [{n: 0, d: 0}];
while(pq.length > 0) {
    pq.sort((a,b) => a.d - b.d);
    let curr = pq.shift();
    if(curr.d > dist[curr.n]) continue;
    for(let e of adj[curr.n]) {
        if(dist[curr.n] + e.w < dist[e.to]) {
            dist[e.to] = dist[curr.n] + e.w;
            pq.push({n: e.to, d: dist[e.to]});
        }
    }
}
