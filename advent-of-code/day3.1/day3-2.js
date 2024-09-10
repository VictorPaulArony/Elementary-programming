const fs = require('fs');

// Read the data from the file
const directions = fs.readFileSync('data.txt', 'utf8');

function countHousesWithPresents(directions) {
    let visitedHouses = new Set();
    
    // Initial positions of Santa and Robo-Santa
    let santaX = 0, santaY = 0;
    let roboX = 0, roboY = 0;

    // Add the starting position to the set (visited by both Santa and Robo-Santa)
    visitedHouses.add(`0,0`);

    // Process each direction
    for (let i = 0; i < directions.length; i++) {
        // Determine who moves: Santa on odd steps, Robo-Santa on even steps
        if (i % 2 === 0) {  // Santa's turn
            switch(directions[i]) {
                case '^':
                    santaY += 1;
                    break;
                case 'v':
                    santaY -= 1;
                    break;
                case '>':
                    santaX += 1;
                    break;
                case '<':
                    santaX -= 1;
                    break;
            }
            visitedHouses.add(`${santaX},${santaY}`);
        } else {  // Robo-Santa's turn
            switch(directions[i]) {
                case '^':
                    roboY += 1;
                    break;
                case 'v':
                    roboY -= 1;
                    break;
                case '>':
                    roboX += 1;
                    break;
                case '<':
                    roboX -= 1;
                    break;
            }
            visitedHouses.add(`${roboX},${roboY}`);
        }
    }

    // Return the number of unique houses visited
    return visitedHouses.size;
}

// Execute the function with the directions from the file
const uniqueHouses = countHousesWithPresents(directions);
console.log("Number of houses that received at least one present:", uniqueHouses);
