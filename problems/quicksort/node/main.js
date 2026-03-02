const fs = require('fs');

function quicksort(arr) {
    if (arr.length <= 1) return arr;
    const pivot = arr[Math.floor(arr.length / 2)];
    const left = [], middle = [], right = [];
    for (let x of arr) {
        if (x < pivot) left.push(x);
        else if (x > pivot) right.push(x);
        else middle.push(x);
    }
    return quicksort(left).concat(middle, quicksort(right));
}

const input = fs.readFileSync(0, 'utf-8').trim().split(/\s+/).map(Number);
if (input.length > 0 && !isNaN(input[0])) {
    quicksort(input);
}
