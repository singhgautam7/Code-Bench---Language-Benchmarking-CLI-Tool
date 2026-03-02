use std::io::{self, Read};
fn gcd(mut a: i32, mut b: i32) -> i32 {
    while b != 0 { let t = b; b = a % b; a = t; }
    a
}
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    if let Ok(n) = s.trim().parse::<i32>() {
        for i in 1..=n { gcd(1000000007, i); }
    }
}