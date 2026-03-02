use std::io::{self, Read};
fn dfs(u: usize, adj: &Vec<Vec<usize>>, vis: &mut Vec<bool>) {
    vis[u] = true;
    for &v in &adj[u] {
        if !vis[v] { dfs(v, adj, vis); }
    }
}
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    let mut tokens = s.split_whitespace().filter_map(|x| x.parse::<usize>().ok());
    if let Some(v) = tokens.next() {
        let e = tokens.next().unwrap();
        let mut adj = vec![vec![]; v];
        for _ in 0..e {
            let u = tokens.next().unwrap();
            let nxt = tokens.next().unwrap();
            tokens.next(); // w
            adj[u].push(nxt);
        }
        let mut vis = vec![false; v];
        dfs(0, &adj, &mut vis);
    }
}