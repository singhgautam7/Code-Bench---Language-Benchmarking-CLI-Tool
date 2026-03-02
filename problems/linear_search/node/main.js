const fs = require('fs');
let tokens = fs.readFileSync(0, 'utf-8').trim().split(/\s+/);
if(tokens.length > 0) {
    let target = tokens[tokens.length-1];
    for(let i=0; i<tokens.length-1; i++) {
        if(tokens[i] === target) break;
    }
}