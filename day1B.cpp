#include <iostream>

int main()
{
    int a, b, c;
    std::cin >> a >> b >> c;
    int last = a+b+c;
    int count=0;
    while (true)
    {
        a=b;
        b=c;
        std::cin >> c;
        if (c==0)
            break;
        if (a+b+c>last)
            count++;
        last=a+b+c;
    }
    std::cout << count << std::endl;
    return 0;
}