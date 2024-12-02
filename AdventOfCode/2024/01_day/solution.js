const fs = require('fs');

const file = fs.readFileSync('./input.txt','utf8');
const input = file.split('\r\n');

// Solution 1
let firstList = [];
let secondList = [];
for (let line of input) {
    const numbers = line.split("   ");
    firstList.push(parseInt(numbers[0]));
    secondList.push(parseInt(numbers[1]));
}

firstList = firstList.sort();
secondList = secondList.sort();

let distanceScore = 0;
for (let i = 0; i<firstList.length; i++) {
    distanceScore += Math.abs(firstList[i] - secondList[i])
}

console.log('part 1 solution: ', distanceScore);

let similarityScore = 0;
for (number of firstList) {
    let count = 0;
    for (comparison of secondList) {
        if (number == comparison) {
            count++;
        }
    }
    similarityScore += count * number;
}

console.log('part 2 solution: ', similarityScore);