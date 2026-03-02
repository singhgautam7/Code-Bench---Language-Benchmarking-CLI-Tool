const fs = require('fs');
let n = parseInt(fs.readFileSync(0, 'utf-8'));
if(n >= 2) {
    let sieve = new Uint8Array(n+1);
    let count = 0;
    for(let p=2; p<=n; p++) {
        if(sieve[p] === 0) {
            count++;
            for(let i=p*p; i<=n; i+=p) sieve[i] = 1;
        }
    }
}
