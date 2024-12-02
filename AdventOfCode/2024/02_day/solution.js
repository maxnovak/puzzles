const fs = require('fs');

const file = fs.readFileSync('./input.txt','utf8');
const input = file.split('\r\n');

// Solution 1
let reactorData = [];
for (line of input) {
    const reactorLevels = line.split(" ");
    reactorData.push([...reactorLevels.map(element => parseInt(element))]);
}



const isSafe = (data) => {
    let previous;
    let direction = "";
    for (level of data) {
        if (previous == null) {
            previous = level;
            continue;
        }
        if (direction === "") {
            if (level > previous) {
                direction = "increasing";
            }
            if (level < previous) {
                direction = "decreasing";
            }
            if (level == previous) {
                return false;
            }
        }
        if (direction === "decreasing") {
            if ((previous - level) > 3 || (previous - level) <= 0) {
                return false;
            }
        }
        if (direction === "increasing") {
            if ((level - previous) > 3 || (level - previous) <= 0) {
                return false;
            }
        }
        previous = level
    }
    return true;
}

let safeCount = 0;
for (data of reactorData) {
    if (isSafe(data)) {
        safeCount++;
    }
}

console.log('part 1 solution: ', safeCount);

let safeCountWithDampener = 0;
for (data of reactorData) {
    if (isSafe(data)) {
        safeCountWithDampener++
    } else {
        for (let i = 0; i < data.length; i++) {
            if (isSafe([...data.slice(0,i), ...data.slice(i+1)])) {
                safeCountWithDampener++;
                break;
            }
        }
    }
}

console.log('part 2 solution: ', safeCountWithDampener);