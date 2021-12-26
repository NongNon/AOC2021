const fs = require('fs')

let data
try {
    data = fs.readFileSync('./in.txt', 'utf8')
} catch (err) {
    console.error(err)
}
let list=[];
data.match(/\| (.+)\n?/g).forEach(word => {
    word = word.replace(/\| /, '').replace(/\n/, '')
    list.push(...(word.split(' ')))
})
//console.log(list)

let total=0;
for (let i = 0; i < list.length; i++) {
    let n = list[i].length
    if(n==2||n==4||n==3||n==7){
        // console.log(list[i])
        total++;
    }
}

console.log(total)