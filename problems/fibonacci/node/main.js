const fs = require('fs');
let n = parseInt(fs.readFileSync(0, 'utf-8'));
let a = 0n, b = 1n, t;
for(let i=0; i<n; i++) { t=a; a=b; b=t+b; }
