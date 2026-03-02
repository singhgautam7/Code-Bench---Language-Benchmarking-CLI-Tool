use std::io::{self, Read};
use std::collections::HashSet;
use std::cmp::max;
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    let s = s.trim().as_bytes();
    let mut chars = HashSet::new();
    let (mut l, mut res) = (0, 0);
    for r in 0..s.len() {
        while chars.contains(&s[r]) { chars.remove(&s[l]); l+=1; }
        chars.insert(&s[r]);
        res = max(res, r-l+1);
    }
}