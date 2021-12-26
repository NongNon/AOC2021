let dfs=(row,col,data)=>{
    if(row>=data.length||col>=data[0].length||row<0||col<0)
        return 0;
    if (data[row][col]==9)
        return 0;
    data[row][col]=9;
    return 1+dfs(row+1,col,data)+dfs(row-1,col,data)+dfs(row,col+1,data)+dfs(row,col-1,data);
}

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

let list=[]
for(let row = 0; row < data.length; row++) {
    for(let col = 0; col < data[row].length; col++) {
        if(data[row][col]<9){
            //console.log(row+" "+col+" "+dfs(row,col,data))
            list.push(dfs(row,col,data))
        }
    }
}
list=list.sort((a,b)=>b-a)
console.log(list[0]*list[1]*list[2])
