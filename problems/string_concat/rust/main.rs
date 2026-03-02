use std::io::{self, Read};
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    let parts: Vec<&str> = s.split_whitespace().collect();
    if parts.len() == 2 {
        let n: usize = parts[1].parse().unwrap();
        let mut out = String::with_capacity(n);
        for _ in 0..n { out.push_str(parts[0]); }
    }
}