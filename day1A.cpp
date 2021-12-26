#include <iostream>

int main()
{
    int n,last=999999;
    int count=0;
    while (true)
    {
        std::cin >> n;
        if (n==0)
            break;
        if (n>last)
            count++;
        last=n;
    }
    std::cout << count << std::endl;
    return 0;
}