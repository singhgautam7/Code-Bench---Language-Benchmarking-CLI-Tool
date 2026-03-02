const fs = require('fs');
let n = parseInt(fs.readFileSync(0, 'utf-8').trim());
let h = new Map();
for(let i=0; i<n; i++) h.set(i.toString(), i);
