f = open("in.txt", "r").read()
bottom_y = int(f.split(" ")[-1].split("=")[-1].split(".")[0])
print(bottom_y*(bottom_y+1)//2)