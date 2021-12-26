#include <iostream>
#include <stdio.h>

int main(){
    int input=500;
    int sqsize=1000;
    int square[sqsize][sqsize];
    //fill square with 0
    for(int i=0;i<sqsize;i++){
        for(int j=0;j<sqsize;j++){
            square[i][j]=0;
        }
    }

    while (input--){
        int x1,y1,x2,y2;
        scanf("%d,%d -> %d,%d",&x1,&y1,&x2,&y2);
        if (x1!=x2 && y1!=y2){
            if((y2-y1)/double(x2-x1)!=1.0&&(y2-y1)/double(x2-x1)!=(-1.0)){
                std::cout<<"booya"<<std::endl;
                std::cout << (y2-y1)/double(x2-x1) << std::endl;
            }
            if(x1<x2==y1<y2){
                int s_x=(x1<x2)?x1:x2;
                int s_y=(y1<y2)?y1:y2;
                int e_x=(x1>x2)?x1:x2;
                while(s_x<=e_x){
                    square[s_y][s_x]++;
                    s_x++;
                    s_y++;
                }
            }else{
                int s_x=(x1<x2)?x1:x2;
                int s_y=(y1>y2)?y1:y2;
                int e_x=(x1>x2)?x1:x2;
                while(s_x<=e_x){
                    square[s_y][s_x]++;
                    s_x++;
                    s_y--;
                }
            }
        }
        else if (x1==x2){
            // swap y1 y2 if y1>y2
            if (y1>y2){
                int temp=y1;
                y1=y2;
                y2=temp;
            }
            for (int i=y1;i<=y2;i++){
                square[i][x1]++;
            }
        }
        else if (y1==y2){
            // swap x1 x2 if x1>x2
            if (x1>x2){
                int temp=x1;
                x1=x2;
                x2=temp;
            }
            for (int i=x1;i<=x2;i++){
                square[y1][i]++;
            }
        }
    }
    // count value more than and equal to 2
    long count=0;
    for (int i=0;i<sqsize;i++){
        for (int j=0;j<sqsize;j++){
            if (square[i][j]>=2){
                count++;
            }
        }
    }
    std::cout<<count<<std::endl;

    // print square
    // for (int i=0;i<10;i++){
    //     for (int j=0;j<10;j++){
    //         printf("%d ",square[i][j]);
    //     }
    //     printf("\n");
    // }
    return 0;
}