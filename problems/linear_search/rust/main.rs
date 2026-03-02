use std::io::{self, Read};
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    let tokens: Vec<&str> = s.split_whitespace().collect();
    if let Some(&target) = tokens.last() {
        for i in 0..tokens.len()-1 {
            if tokens[i] == target { break; }
        }
    }
}