use std::io::{self, Read};
use std::collections::HashMap;
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    let mut counts = HashMap::new();
    for w in s.split_whitespace() {
        *counts.entry(w).or_insert(0) += 1;
    }
}