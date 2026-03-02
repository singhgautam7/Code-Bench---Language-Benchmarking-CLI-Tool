const fs = require('fs');
let s = fs.readFileSync(0, 'utf-8').trim().split(/\s+/);
let counts = new Map();
for(let w of s) {
    counts.set(w, (counts.get(w) || 0) + 1);
}
