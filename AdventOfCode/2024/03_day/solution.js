const fs = require('fs');

const input = fs.readFileSync('./input.txt','utf8');

// Solution 1
const validRegexGroups = /mul\((\d{1,3}),(\d{1,3})\)/gm;

let validInputs = input.matchAll(validRegexGroups);
let total = 0;
for (const match of validInputs) {
    total += parseInt(match[1]) * parseInt(match[2]);
}

console.log("Solution 1: ", total)

// Solution 2
const doRegexGroups = /mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)/gm

let validCommands = input.matchAll(doRegexGroups);
let total2 = 0;
let enabled = true;
for (const match of validCommands) {
    if (match[0] === 'do()') {
        enabled = true;
        continue
    }
    if (match[0] === 'don\'t()') {
        enabled = false;
        continue
    }
    if (enabled) {
        total2 += parseInt(match[1]) * parseInt(match[2]);
    }
}

console.log("Solution 2: ", total2)