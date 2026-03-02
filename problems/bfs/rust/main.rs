use std::io::{self, Read};
use std::collections::VecDeque;
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
        let mut q = VecDeque::new();
        let mut vis = vec![false; v];
        q.push_back(0);
        vis[0] = true;
        while let Some(curr) = q.pop_front() {
            for &nxt in &adj[curr] {
                if !vis[nxt] {
                    vis[nxt] = true;
                    q.push_back(nxt);
                }
            }
        }
    }
}