use std::io::{self, Read};
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    let s = s.trim().as_bytes();
    for i in 0..s.len()/2 {
        if s[i] != s[s.len()-1-i] { return; }
    }
}