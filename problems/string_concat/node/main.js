const fs = require('fs');
let [c, n] = fs.readFileSync(0, 'utf-8').trim().split(/\s+/);
let s = ""; for(let i=0; i<parseInt(n); i++) s += c;
