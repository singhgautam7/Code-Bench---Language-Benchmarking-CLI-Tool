use std::io::{self, Read};
fn main() {
    let mut s = String::new();
    io::stdin().read_to_string(&mut s).unwrap();
    // Simulate parse load since no serde/json in STL
    let _c = s.split('{').count();
}