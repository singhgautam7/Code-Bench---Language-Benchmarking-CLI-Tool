use std::io::{self, Read};
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    if let Ok(n) = s.trim().parse::<u64>() {
        let (mut a, mut b) = (0u64, 1u64);
        for _ in 0..n { let t = a; a = b; b = b.wrapping_add(t); }
    }
}