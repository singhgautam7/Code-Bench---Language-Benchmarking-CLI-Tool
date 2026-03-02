use std::io::{self, Read};
fn mergesort(arr: &mut [i32]) {
    let len = arr.len();
    if len <= 1 { return; }
    let mid = len / 2;
    mergesort(&mut arr[0..mid]);
    mergesort(&mut arr[mid..len]);
    let mut res = Vec::with_capacity(len);
    let (mut i, mut j) = (0, mid);
    while i < mid && j < len {
        if arr[i] < arr[j] { res.push(arr[i]); i += 1; }
        else { res.push(arr[j]); j += 1; }
    }
    res.extend_from_slice(&arr[i..mid]);
    res.extend_from_slice(&arr[j..len]);
    arr.copy_from_slice(&res);
}
fn main() {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();
    let mut nums: Vec<i32> = input.split_whitespace().filter_map(|s| s.parse().ok()).collect();
    mergesort(&mut nums);
}