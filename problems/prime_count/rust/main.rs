use std::io::{self, Read};
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    if let Ok(n) = s.trim().parse::<usize>() {
        if n < 2 { return; }
        let mut sieve = vec![true; n+1];
        let mut count = 0;
        for p in 2..=n {
            if sieve[p] {
                count += 1;
                let mut i = p * p;
                while i <= n {
                    sieve[i] = false;
                    i += p;
                }
            }
        }
    }
}