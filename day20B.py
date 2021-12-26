def bina(img,i,j,o):
    neighbors =((i-1,j-1),(i-1,j),(i-1,j+1),
                (i,j-1),(i,j),(i,j+1),
                (i+1,j-1),(i+1,j),(i+1,j+1))
    binary=""
    for neighbor in neighbors:
        if neighbor[0]>=0 and neighbor[0]<len(img) and neighbor[1]>=0 and neighbor[1]<len(img[0]):
            binary+=img[neighbor[0]][neighbor[1]]
        else :
            binary+=o
    return binary

if __name__ == "__main__":
    f= open("in.txt","r").read()
    enchancement,img_str=[i.replace(".","0").replace("#","1") for i in f.split("\n\n")]
    img=[list(i) for i in img_str.split("\n")]
    #print(img)
    for i in range(50):
        o="0" if i%2 == 0  else "1"
        new_img=[]
        for i in range(-1,len(img)+1):
            new_img.append(["."]*(len(img[0])+2))
            for j in range(-1,len(img[0])+1):
                new_img[i+1][j+1]=enchancement[int(bina(img,i,j,o),2)]
        img=new_img
    
    print(sum([i.count("1") for i in img]))
    #print(img)