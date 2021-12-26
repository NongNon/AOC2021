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
    return start_index+5


def dfs(binary_package):
    v = int(binary_package[0:3],2)
    #print("v = "+str(v))
    t = int(binary_package[3:6],2)

    if t == 4:
        return lit_package(binary_package[6:])+6,v
    else :
        i = binary_package[6]
        #l = int(binary_package[7:22] if i == '0' else binary_package[7:18],2)
        if i == '0':
            l = int(binary_package[7:22],2)
            sub_package = binary_package[22:22+l]
            sum_v=v
            while len(sub_package) > 0:
                end_index,child_v = dfs(sub_package)
                sum_v += child_v
                sub_package = sub_package[end_index:]
            return 22+l,sum_v
        else :
            l = int(binary_package[7:18],2)
            sub_package = binary_package[18:]
            sum_v=v
            sum_index= 18;
            for w in range(l):  
                end_index,child_v = dfs(sub_package)
                sum_v += child_v
                sum_index += end_index
                sub_package = sub_package[end_index:]  
            return sum_index,sum_v


if __name__ == "__main__":
    a = input()
    binary = bin(int(a, 16))[2:].zfill(len(a)*4)
    _, sum_v = dfs(binary)
    print(sum_v)


