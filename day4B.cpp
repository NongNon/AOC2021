#include <iostream>
#include <string>
#include <cstring>  

class Bingo{
    private:
        bool isBingo(int x,int y){
            int i=0;
            for (i = 0; i < 5; i++) {
                if(!check[x][i]){
                    break;
                }
            }
            if(i == 5){
                return true;
            }
            for (i = 0; i < 5; i++) {
                if(!check[i][y]){
                    break;
                }
            }
            if(i == 5){
                return true;
            }else{
                return false;
            }
        };
    public:
        int num[5][5];
        bool check[5][5]={false};

        Bingo(){
            for(int i=0;i<5;i++){
                for(int j=0;j<5;j++){
                    num[i][j]=0;
                }
            }
        }

        void SetNumber(){
            for(int i=0;i<5;i++){
                for(int j=0;j<5;j++){
                    std::cin >> num[i][j];
                }
            }
        }

        bool CheckNumber(int n){
            for(int i=0;i<5;i++){
                for(int j=0;j<5;j++){
                    if(num[i][j]==n){
                        check[i][j]=true;
                        return isBingo(i,j);
                    }
                }
            }
            return false;
        }

        void Print(){
            for(int i=0;i<5;i++){
                for(int j=0;j<5;j++){
                    std::cout << num[i][j] << " ";
                }
                std::cout << std::endl;
            }
        }
};


int main(){
    int n_bingo=100;
    char drawn[10000];
    int q[100000];
    int size=0;
    std::cin >> drawn;
    char* ptr= std::strtok(drawn,",");
    while (ptr != NULL)  
    {  
        q[size++]=atoi(ptr);
        ptr= std::strtok(NULL,",");
    }  

    Bingo bingo[n_bingo];
    for(int i=0;i<n_bingo;i++){
        bingo[i].SetNumber();
    }
    int valid_bingo= n_bingo;
    bool is_bingo[n_bingo]={false};
    int whatisbingo=-1,lastnum=-1;
    for (int i = 0; i < size; i++) {
        for (int j = 0; j < n_bingo; j++) {
            if(is_bingo[j]){
                continue;
            }else
            if(bingo[j].CheckNumber(q[i])){
                whatisbingo=j;
                lastnum=q[i];
                is_bingo[j]=true;
                valid_bingo--;
                if(valid_bingo==0){
                    break;
                }
            }
        }
        if(valid_bingo==0){
            break;
        }
    }
    int sum=0;
    for(int i=0;i<5;i++){
        for(int j=0;j<5;j++){
            if(!bingo[whatisbingo].check[i][j]){
                sum+=bingo[whatisbingo].num[i][j];
            }
        }
    }

    std::cout << sum *lastnum << std::endl;

    // //print bingo
    // for(int i=0;i<n_bingo;i++){
    //     bingo[i].Print();
    //     std::cout << std::endl;
    // }

    return 0;
}