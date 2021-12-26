const fs = require('fs');
let data="";
try {
    data = fs.readFileSync('./in.txt', 'utf8');
}catch(e){
    console.log(e);
}
//console.log(data);
data = data.split("\n");

let incompTable={
    ')':1,
    ']':2,
    '}':3,
    '>':4
}

let compScore=[];
for(let line=0;line<data.length;line++){
    let stack = [];
    let currupted=false
    for(let i=0;i<data[line].length;i++){
        if(data[line][i].match(/[\(\[\{\<]/))
            stack.push(data[line][i]);
        else{
            let p = stack.pop()+data[line][i];
            if(p!=="()" && p!=="[]" && p!=="{}" && p!=="<>"){
                currupted=true;
                break;
            }
        }
    }
    if(currupted) continue;
    let score =0;
    while(stack.length>0){
        let c = stack.pop();
        score*=5;
        if(c==="(") score+= incompTable[')']
        else if(c==="[") score+= incompTable[']']
        else if(c==="{") score+= incompTable['}']
        else if(c==="<") score+= incompTable['>']
    }
    compScore.push(score);
}

compScore.sort((a,b)=> a-b);
console.log(compScore[Math.floor(compScore.length/2)]);