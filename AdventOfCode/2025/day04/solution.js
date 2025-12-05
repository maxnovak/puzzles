const fs = require('fs');

const file = fs.readFileSync('./input.txt','utf8');
const input = file.split('\n');

let warehouse = [];

for (let line of input) {
  let row = [];
  for (let item of line) {
    if (item === ".") {
      row.push("space");
    }
    if (item === "@") {
      row.push("paper");
    }
  }
  warehouse.push(row);
}

const searchCoords = [
  [-1,-1], [-1,0], [-1,1],
  [0,-1],  /*[0,0]*/  [0,1],
 [1,-1],  [1,0],  [1,1]
];

let accessable = 0;

for (let y in warehouse) {
  for (let x in warehouse[y]) {
    if (warehouse[y][x] === "space") {
      continue;
    }

    let filledSpaceCount = 0;
    for (let coord of searchCoords) {
      const searchY = parseInt(y) + coord[0];
      const searchX = parseInt(x) + coord[1];
      if (!warehouse[searchY]) {
        continue;
      }
      if (!warehouse[searchY][searchX]){
        continue;
      }
      if (warehouse[searchY][searchX] === "paper") {
        filledSpaceCount++;
      }
    }
    if (filledSpaceCount < 4) {
      accessable++;
    }
  }
}

console.log(accessable);
