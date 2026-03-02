use std::io::{self, Read};
use std::collections::BinaryHeap;
use std::cmp::Ordering;
#[derive(Copy, Clone, Eq, PartialEq)]
struct State { cost: usize, position: usize }
impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering { other.cost.cmp(&self.cost).then_with(|| self.position.cmp(&other.position)) }
}
impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> { Some(self.cmp(other)) }
}
struct Edge { node: usize, cost: usize }

fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    let mut tokens = s.split_whitespace().filter_map(|x| x.parse::<usize>().ok());
    if let Some(v) = tokens.next() {
        let e = tokens.next().unwrap();
        let mut adj: Vec<Vec<Edge>> = (0..v).map(|_| Vec::new()).collect();
        for _ in 0..e {
            let u = tokens.next().unwrap();
            let nxt = tokens.next().unwrap();
            let w = tokens.next().unwrap();
            adj[u].push(Edge { node: nxt, cost: w });
        }
        let mut dist = vec![usize::MAX; v];
        let mut heap = BinaryHeap::new();
        dist[0] = 0;
        heap.push(State { cost: 0, position: 0 });
        while let Some(State { cost, position }) = heap.pop() {
            if cost > dist[position] { continue; }
            for edge in &adj[position] {
                let next = State { cost: cost + edge.cost, position: edge.node };
                if next.cost < dist[next.position] {
                    heap.push(next);
                    dist[next.position] = next.cost;
                }
            }
        }
    }
}