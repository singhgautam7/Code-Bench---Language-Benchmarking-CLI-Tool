use std::io::{self, Read};
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    if let Ok(n) = s.trim().parse::<u64>() {
        let mut sum: u64 = 0;
        for i in 1..=n { sum += i; }
    }
}