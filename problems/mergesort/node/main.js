const fs = require('fs');
function mergesort(arr) {
    if(arr.length <= 1) return arr;
    const mid = Math.floor(arr.length/2);
    const L = mergesort(arr.slice(0, mid));
    const R = mergesort(arr.slice(mid));
    let res = [], i=0, j=0;
    while(i < L.length && j < R.length) {
        if(L[i] < R[j]) res.push(L[i++]);
        else res.push(R[j++]);
    }
    return res.concat(L.slice(i)).concat(R.slice(j));
}
const input = fs.readFileSync(0, 'utf-8').trim().split(/\s+/).map(Number);
if (input.length > 0 && !isNaN(input[0])) mergesort(input);
