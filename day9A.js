const fs = require('fs')

let data
try {
    data = fs.readFileSync('./in.txt', 'utf8')
} catch (err) {
    console.error(err)
}
data = data.split('\n')
for(let index in  data) {
    data[index] = data[index].split('')
}
data = data.map(r => r.map(c => parseInt(c)))
let total=0;
for(let i = 0; i < data.length; i++) {
    for(let j = 0; j < data[i].length; j++) {
        let left= data[i][j-1]===undefined?10:data[i][j-1];
        let right = data[i][j+1]===undefined?10:data[i][j+1];
        let top = data[i-1] ? data[i-1][j] : 10;
        let bottom = data[i+1] ? data[i+1][j] : 10;
        if(data[i][j] < left && data[i][j] < right && data[i][j] < top && data[i][j] < bottom) {
            console.log(`${i}, ${j}, ${data[i][j]}`)
            total += data[i][j]+1
        }
    }
}
console.log(total)