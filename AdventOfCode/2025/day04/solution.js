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

const getAccessableRolls = (warehouse) => {
  let accessable = 0;
  const spacesToClear = []
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
        spacesToClear.push([y,x])
      }
    }
  }
  return [accessable, spacesToClear];
}

console.log("Part 1 solution: ", getAccessableRolls(warehouse)[0]);
let freedSpaces = 0
let totalFreedSpaces = 0
do {
  response = getAccessableRolls(warehouse);
  freedSpaces = response[0];
  const spacesToClear = response[1];
  for (let space of spacesToClear) {
    warehouse[space[0]][space[1]] = "space";
  }
  totalFreedSpaces += freedSpaces;
} while(freedSpaces > 0)

console.log("Part 2 solution: ", totalFreedSpaces);
