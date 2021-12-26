import re
f = open("in.txt", "r").read()
f= re.sub(r"[a-z:,=]", "", f).replace("..", " ").strip()
L_x,R_x,B_y,T_y = [int(i) for i in f.split(" ")]

max_x = R_x
min_x=0
for i in range(0,L_x):
    if (i*(i+1))>=L_x*2 and (i*(i+1))<=R_x*2:
        min_x = i
        break

max_y= abs(B_y+1)
min_y= B_y

count=0;
for i in range(min_x,max_x+1):
    for j in range(min_y,max_y+1):
        x,y= [0,0]
        add_x=i
        add_y=j
        while y>=min_y and x<=max_x:
            if x>=L_x and x<=R_x and y>=B_y and y<=T_y:
                count+=1
                break
            x+=add_x
            y+=add_y
            add_y-=1
            if(add_x>0):
                add_x-=1
            
print(count)