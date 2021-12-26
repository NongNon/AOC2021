#include <iostream>
#include <string>

int main()
{
    int x=0,y=0;
    int aim=0;
    while(true){
        std::string s;
        std::cin >> s;
        int n ;
        std::cin >> n;
        if(s=="forward"){
            x+=n;
            y+=(n*aim);
        }
        else if(s=="up"){
            aim-=n;
        }
        else if(s=="down"){
            aim+=n;
        }
        else if(s=="exit"){
            break;
        }
    }
    std::cout << x*y << std::endl;
}