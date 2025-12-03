const fs = require('fs');

const file = fs.readFileSync('./input.txt','utf8');
const input = file.split('\n');

const dialStart = 50;
let commands = [];

for (let line of input) {
  commands.push({
    direction: line[0],
    distance: parseInt(line.slice(1)),
  });
}

// Solution 1
let zeroCount = 0;
let position = dialStart;
for (let command of commands) {
  if (command.direction === 'R') {
    position += command.distance;
  }
  if (command.direction === 'L') {
    position -= command.distance;
  }
  position = position % 100
  if (position < 0) {
    position += 100
  }
  if (position == 0) {
    zeroCount++;
  }
  console.log('position after turn', position);
}

console.log('Solution 1 Zero Count: ', zeroCount)

// Solution 2
zeroCount = 0;
position = dialStart;
movement = 0;
for (let command of commands) {
  if (command.direction === 'R') {
    movement = position + command.distance;
    zeroCount += Math.abs(Math.floor((position-1)/100) - Math.floor((movement-1)/100))
  }
  if (command.direction === 'L') {
    movement = position  - command.distance;
    zeroCount += Math.abs(Math.floor((position)/100) - Math.floor((movement)/100))
  }

  position = movement % 100
  if (position < 0) {
    position += 100
  }

  console.log('position after turn: ', position, ' zeroCount: ', zeroCount, command);
}

console.log('Solution 2 Zero Count: ', zeroCount);
