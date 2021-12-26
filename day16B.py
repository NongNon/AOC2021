import math
def lit_package(binary):
    lit_value = ""
    start_index=0;
    start_bit = binary[start_index]
    while start_bit != '0':
        lit_value=lit_value+binary[start_index+1:start_index+5]
        start_index += 5
        start_bit = binary[start_index]

    lit_value = lit_value + binary[start_index+1:start_index+5]
    #print("lit = "+str(int(lit_value,2)))
    return start_index+5,int(lit_value,2)


def dfs(binary_package):
    v = int(binary_package[0:3],2)
    #print("v = "+str(v))
    t = int(binary_package[3:6],2)

    if t == 4:
        end_index,lit_value =  lit_package(binary_package[6:])
        return end_index+6,lit_value
    else :
        i = binary_package[6]
        #l = int(binary_package[7:22] if i == '0' else binary_package[7:18],2)
        index=0
        lit_set=[]
        if i == '0':
            l = int(binary_package[7:22],2)
            sub_package = binary_package[22:22+l]
            while len(sub_package) > 0:
                end_index,lit_value = dfs(sub_package)
                lit_set.append(lit_value)
                sub_package = sub_package[end_index:]
            index = 22+l
        else :
            l = int(binary_package[7:18],2)
            sub_package = binary_package[18:]
            sum_index= 18;
            for w in range(l):  
                end_index,lit_value = dfs(sub_package)
                lit_set.append(lit_value)
                sum_index += end_index
                sub_package = sub_package[end_index:]  
            index = sum_index
        
        result=0;
        #print(t,lit_set)
        if t == 0:
            result = sum(lit_set)
        elif t == 1:
            result = math.prod(lit_set)
        elif t == 2:
            result = min(lit_set)
        elif t == 3:
            result = max(lit_set)
        elif t == 5:
            result = 1 if lit_set[0]>lit_set[1] else 0
        elif t == 6:
            result = 1 if lit_set[0]<lit_set[1] else 0
        elif t == 7:
            result = 1 if lit_set[0]==lit_set[1] else 0
        
        return index,result


if __name__ == "__main__":
    a = input()
    binary = bin(int(a, 16))[2:].zfill(len(a)*4)
    _,result=dfs(binary)
    print(result)

