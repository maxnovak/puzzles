const fs = require('fs');

const file = fs.readFileSync('./input.txt','utf8');
const input = file.split('\n');

let batteryBanks = [];

for (let line of input) {
  batteryBanks.push(line);
}

let totalJoltage = 0;

for (let bank of batteryBanks) {
  let highestJoltageInBank = 0;
  for (let batteryIndex in bank) {
    if (parseInt(batteryIndex) + 1 === bank.length) {
      break;
    }
    for (let i = parseInt(batteryIndex) + 1; i < bank.length; i++) {
      const testJoltage = parseInt(bank[batteryIndex] + bank[i]);
      if (testJoltage > highestJoltageInBank) {
        highestJoltageInBank = testJoltage;
      }
    }
  }
  totalJoltage += highestJoltageInBank;
}

console.log("Total Highest Joltage: ", totalJoltage);