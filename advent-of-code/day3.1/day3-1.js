const file = require('fs')

const directions = file.readFileSync('data.txt', 'utf8')

function findDirection(directions){
    let visited = new Set()
        let x = 0, y = 0
        visited.add(`${x}, ${y}`)

        for (let i = 0; i < directions.length;i++){
            switch(directions[i]){
                case '^':
                    y += 1
                    break;
                case 'v':
                    y += 1
                    break;
                case '>':
                    x += 1
                    break;
                case '<':
                    x += 1
                    break;
            }
            visited.add(`${x}, ${y}`)
        }
        return visited.size

}
const count = findDirection(directions)
console.log(count)