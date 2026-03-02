use std::io::{self, Read};
use std::collections::HashMap;
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    if let Ok(n) = s.trim().parse::<usize>() {
        let mut h = HashMap::with_capacity(n);
        for i in 0..n { h.insert(i.to_string(), i); }
    }
}