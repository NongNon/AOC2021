#include <iostream>
#include <string>

int main()
{
    int ndigits=12, nline=1000;
    char data[nline][ndigits];
    //input string 
    for(int i=0; i<nline; i++)
    {
        std::cin >> data[i];
    }
    // copy data to new array
    char newdata[nline][ndigits];
    for(int i=0; i<nline; i++)
    {
        for(int j=0; j<ndigits; j++)
        {
            newdata[i][j]=data[i][j];
        }
    }

    // //output string
    // for(int i=0; i<nline; i++)
    // {
    //     for(int j=0; j<ndigits; j++)
    //     {
    //         std::cout << data[i][j];
    //     }
    //     std::cout << std::endl;
    // }

    //gamma value
    int size = nline;
    for (int col=0; col<ndigits; col++)
    {
        int bitcount = 0;
        for (int row=0; row<size; row++)
        {
            if (data[row][col] == '1')
            {
                bitcount++;
            }
        }
        int mostbits = (bitcount >= size-bitcount)?'1':'0';
        int newsize=0;
        for (int row=0; row<size; row++)
        {
            if (data[row][col] == mostbits)
            {   
                for (int i=0; i<ndigits; i++)
                {
                    data[newsize][i] = data[row][i];    
                }
                newsize++;
            }
        } 
        size = newsize;
    }

    size = nline;
    for (int col=0; col<ndigits; col++)
    {
        int bitcount = 0;
        for (int row=0; row<size; row++)
        {
            if (newdata[row][col] == '1')
            {
                bitcount++;
            }
        }
        int lestbits = (bitcount >= size-bitcount)?'0':'1';
        int newsize=0;
        for (int row=0; row<size; row++)
        {
            if (newdata[row][col] == lestbits)
            {   
                for (int i=0; i<ndigits; i++)
                {
                    newdata[newsize][i] = newdata[row][i];    
                }
                newsize++;
            }
        } 
        size = newsize;
    }

    int gamma=0,co=0;
    for (int col=0;col<ndigits;col++)
    {
        if (data[0][col]=='1')
        {
            gamma+=(1<<ndigits-col-1);
        }
        if (newdata[0][col]=='1')
        {
            co+=(1<<ndigits-col-1);
        }
    }
    std::cout << gamma << " " << co << std::endl;
    std::cout << gamma*co << std::endl;

}