const fs = require('fs');
let n = parseInt(fs.readFileSync(0, 'utf-8').trim());
function gcd(a, b) {
    while(b) { let t = b; b = a % b; a = t; }
    return a;
}
for(let i=1; i<=n; i++) gcd(1000000007, i);
