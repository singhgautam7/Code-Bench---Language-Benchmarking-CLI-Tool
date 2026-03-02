use std::io::{self, Read};
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    let mut tokens = s.split_whitespace().filter_map(|x| x.parse::<i32>().ok());
    if let Some(n) = tokens.next() {
        let n = n as usize;
        let mut a = vec![vec![0; n]; n];
        let mut b = vec![vec![0; n]; n];
        let mut c = vec![vec![0; n]; n];
        for i in 0..n { for j in 0..n { a[i][j] = tokens.next().unwrap(); } }
        for i in 0..n { for j in 0..n { b[i][j] = tokens.next().unwrap(); } }
        for i in 0..n {
            for k in 0..n {
                let aik = a[i][k];
                for j in 0..n { c[i][j] += aik * b[k][j]; }
            }
        }
    }
}