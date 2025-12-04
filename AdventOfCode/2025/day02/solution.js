const fs = require('fs');

const file = fs.readFileSync('./input.txt','utf8');
const input = file.split(',');

let idRanges = [];
let partOneSolutionTotal = 0;

for (let line of input) {
  const range = line.split('-')
  idRanges.push({
    start: range[0],
    end: range[1],
  });
}

for (let idRange of idRanges) {
  for (var i = parseInt(idRange.start); i <= parseInt(idRange.end); i ++) {
    const stringified = i.toString();
    if (stringified.length % 2 != 0) {
      continue;
    }
    const halfPoint = stringified.length / 2;
    const firstHalf = stringified.slice(0, halfPoint);
    const secondHalf = stringified.slice(halfPoint);
    if (firstHalf === secondHalf) {
      partOneSolutionTotal += i;
    }
  }
}
let partTwoSolutionTotal = 0;

for (let idRange of idRanges) {
  for (var i = parseInt(idRange.start); i <= parseInt(idRange.end); i ++) {
    const stringified = i.toString();
    for (var lengthToCheck = 1; lengthToCheck <= stringified.length / 2; lengthToCheck++){
      if (stringified.length % lengthToCheck != 0) {
        continue;
      }
      const sections = stringified.match(new RegExp(".{1," + lengthToCheck + "}", "g"));
      let copy = sections[0];
      if (sections.every((item) => item === copy)) {
        partTwoSolutionTotal += i;
        break;
      }
    }
  }
  console.log('\n');
}

console.log("part 1 solution: ", partOneSolutionTotal);
console.log("part 2 solution: ", partTwoSolutionTotal);