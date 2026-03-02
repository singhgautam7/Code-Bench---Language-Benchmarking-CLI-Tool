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
let vis = new Uint8Array(V);
function dfs(u) {
    vis[u] = 1;
    for(let v of adj[u]) { if(vis[v] === 0) dfs(v); }
}
try { dfs(0); } catch(e) {} // Handle potential callstack size exceeded
