const fs = require('fs');
let tokens = fs.readFileSync(0, 'utf-8').trim().split(/\s+/);
if(tokens.length === 0 || tokens[0] === "") process.exit(0);
let V = parseInt(tokens[0]), E = parseInt(tokens[1]);
let adj = Array.from({length: V}, () => []);
let idx = 2;
for(let i=0; i<E; i++) {
    let u = parseInt(tokens[idx++]), v = parseInt(tokens[idx++]); idx++;
    adj[u].push(v);
}
let q = [0], head = 0;
let vis = new Uint8Array(V); vis[0] = 1;
while(head < q.length) {
    let curr = q[head++];
    for(let nxt of adj[curr]) {
        if(vis[nxt] === 0) { vis[nxt] = 1; q.push(nxt); }
    }
}
