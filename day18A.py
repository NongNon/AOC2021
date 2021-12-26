# class Node:
#     def __init__(self,l,f,p) -> None:
#         self.left = l
#         self.right = f
#         self.parent = p

#     def __str__(self):
#         return "[" + str(self.left) + "," + str(self.right) + "]"


def snailFishNumber(str):
    stack = []
    for ptr in range(len(str)):
        if str[ptr].isnumeric():
            stack.append(int(str[ptr]))
        if str[ptr]=="]":
            a=stack.pop()
            b=stack.pop()
            stack.append([b,a])
    return stack[0]
    
def process(snNumber):
    i_stack = []
    i=0
    ds_stack = []
    ds=snNumber
    while i<len(ds) or ds !=  snNumber:
        if(isinstance(ds[i],list)):
            if len(ds_stack)>=3:
                i+=1
                continue;
            ds_stack.append(ds)
            i_stack.append(i+1)
            ds=ds[i]
            i=0;
        else:
            i+=1
        while i>=len(ds):
            if  ds==snNumber:
                break
            else:
                ds = ds_stack.pop()
                i = i_stack.pop()
        
# main function
if __name__ == "__main__":
    f = open("in.txt", "r").read().splitlines()
    q=[]
    for line in f:
        
        q.append(process(snailFishNumber(line)))

    

    

