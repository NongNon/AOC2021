#include <iostream>
#include <string>

int main()
{
    int x=0,y=0;
    while(true){
        std::string s;
        std::cin >> s;
        int n ;
        std::cin >> n;
        if(s=="forward"){
            x+=n;
        }
        else if(s=="up"){
            y-=n;
        }
        else if(s=="down"){
            y+=n;
        }
        else if(s=="exit"){
            break;
        }
    }
    std::cout << x*y << std::endl;
}