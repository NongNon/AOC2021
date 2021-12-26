const fs = require('fs')

let data
try {
    data = fs.readFileSync('./in.txt', 'utf8')
} catch (err) {
    console.error(err)
}
let list=data.split('\n');
let total=0;
for(let line=0;line<list.length;line++){
    let signal=list[line].split(' ').slice(0,10).sort((a,b)=>a.length-b.length);
    let dic={}
    while(signal.length>0){
        let s=signal.shift();
        if(s.length==2){
            dic[1]=s.split('');
            continue;
        }
        if(s.length==3){
            dic[7]=s.split('');
            continue;
        }
        if(s.length==4){
            dic[4]=s.split('');
            continue;
        }
        if(s.length==7){
            dic[8]=s.split('');
            continue;
        }

        // check if 5,3,2
        if(s.length==5){
            if(dic[1].every(c=>s.includes(c))){
                dic[3]=s.split('');
                continue;
            }
            else {
                let c=0;
                for(let i=0;i<dic[4].length;i++){
                    if(s.includes(dic[4][i])){
                        c++;
                    }
                }
                if(c==3){
                    dic[5]=s.split('');
                    continue;
                }else{
                    dic[2]=s.split('');
                    continue;
                }
            }
            
        }

        if(s.length==6){
            if(!dic[1].every(c=>s.includes(c))){
                dic[6]=s.split('');
                continue;
            }
            else if (dic[3].every(c=>s.includes(c))){
                dic[9]=s.split('');
                continue;
            }
            else{
                dic[0]=s.split('');
                continue;
            }
        }
    }
    //console.log(dic);

    let digit=list[line].split(' ').slice(11,15)
    let sum=0;
    for (let k=0;k<4;k++){
        let d=-1;
        if(digit[k].length==2){
            d=1;
        }
        else if(digit[k].length==3){
            d=7;
        }
        else if(digit[k].length==4){
            d=4;
        }
        else if(digit[k].length==7){
            d=8;
        }
        else if(digit[k].length==5){
            for(let keys of [2,3,5]){
                if(dic[keys].every(c=>digit[k].includes(c))){
                    d=keys;
                    break;
                }
            }
        }
        else if(digit[k].length==6){
            for(let keys of [0,6,9]){
                if(dic[keys].every(c=>digit[k].includes(c))){
                    d=keys;
                    break;
                }
            }
        }
        if (d===-1){ 
            console.log('error');
            console.log(digit[k])
            break;
        }
        sum+=(d*Math.pow(10,3-k));
    }
    total+=sum;
}
console.log(total);