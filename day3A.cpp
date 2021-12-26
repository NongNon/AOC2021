#include <iostream>

int main()
{
    int ndigits=12;
    int line=0;
    int bitcount[ndigits]={0};
    while (true)
    {
        long input;
        std::cin >> input;
        if (input==0)
            break;
        for(int i = 0; i < ndigits; i++){
            int digit = input % 10;
            if(digit==1){
                bitcount[i]++;
            }
            input=input/10;
        }
        line++;
    }
    int gramma=0,epsilon=0;
    for (int i = 0; i < ndigits; i++)
    {
        if(bitcount[i]>line-bitcount[i]){
            gramma+=(1<<i);
        }
        else{
            epsilon+=(1<<i);
        }
    }
    std::cout << gramma*epsilon << std::endl;
}