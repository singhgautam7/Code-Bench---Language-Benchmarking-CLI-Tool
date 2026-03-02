const fs = require('fs');
let n = parseInt(fs.readFileSync(0, 'utf-8'));
let s = 1n;
for(let i=1n; i<=BigInt(n); i++) s *= i;
