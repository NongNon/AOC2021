const fs = require('fs');
let data="";
try {
    data = fs.readFileSync('./in.txt', 'utf8');
}catch(e){
    console.log(e);
}
//console.log(data);
data = data.split("\n");

let errTable={
    ')':3,
    ']':57,
    '}':1197,
    '>':25137
}

let errScore=0;
for(let line=0;line<data.length;line++){
    let stack = [];
    for(let i=0;i<data[line].length;i++){
        if(data[line][i].match(/[\(\[\{\<]/))
            stack.push(data[line][i]);
        else{
            let p = stack.pop()+data[line][i];
            if(p!=="()" && p!=="[]" && p!=="{}" && p!=="<>")
                errScore+=errTable[data[line][i]];
        }

    }
}
console.log(errScore);


