const fs = require('fs');
let s = fs.readFileSync(0, 'utf-8').trim();
for(let i=0; i<s.length/2; i++) {
   if(s[i] !== s[s.length-1-i]) return;
}