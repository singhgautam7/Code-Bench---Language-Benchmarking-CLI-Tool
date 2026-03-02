const fs = require('fs');
let s = fs.readFileSync(0, 'utf-8').trim();
let chars = new Set();
let l=0, res=0;
for(let r=0; r<s.length; r++) {
    while(chars.has(s[r])) { chars.delete(s[l]); l++; }
    chars.add(s[r]);
    res = Math.max(res, r-l+1);
}