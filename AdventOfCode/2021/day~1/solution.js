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

console.log('Part 1:', increase);

let windowIncrease = 0;

for (let i = 3; i < input.length; i++) {
  const firstWindow = parseInt(input[i-1]) + parseInt(input[i-2]) + parseInt(input[i-3]);
  const secondWindow = parseInt(input[i]) + parseInt(input[i-1]) + parseInt(input[i-2]);

  if(firstWindow < secondWindow) {
    windowIncrease++;
  }
}

console.log("Part 2:", windowIncrease);