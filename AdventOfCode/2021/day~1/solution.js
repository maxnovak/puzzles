const fs = require('fs');

const file = fs.readFileSync('./input.txt','utf8');
const input = file.split('\n')
let increase = 0;

for (let i = 0; i < input.length; i++) {
  if (i === 0) {
    continue;
  }
  if (parseInt(input[i]) > parseInt(input[i-1])) {
    increase++;
  }
};

console.log(increase);