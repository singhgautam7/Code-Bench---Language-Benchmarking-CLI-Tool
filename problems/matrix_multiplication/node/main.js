const fs = require('fs');
let tokens = fs.readFileSync(0, 'utf-8').trim().split(/\s+/);
if (tokens.length === 0 || tokens[0] === '') process.exit(0);
let n = parseInt(tokens[0]);
let idx = 1;
let A = Array.from({length: n}, () => new Int32Array(n));
let B = Array.from({length: n}, () => new Int32Array(n));
let C = Array.from({length: n}, () => new Int32Array(n));
for(let i=0; i<n; i++) for(let j=0; j<n; j++) A[i][j] = parseInt(tokens[idx++]);
for(let i=0; i<n; i++) for(let j=0; j<n; j++) B[i][j] = parseInt(tokens[idx++]);
for(let i=0; i<n; i++) {
    for(let k=0; k<n; k++) {
        for(let j=0; j<n; j++) C[i][j] += A[i][k] * B[k][j];
    }
}