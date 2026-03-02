use std::io::{self, BufRead};
fn main() {
    let stdin = io::stdin();
    let c = stdin.lock().lines().count();
}