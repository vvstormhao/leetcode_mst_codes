#include <iostream>
#include <vector>

using namespace std;
vector<vector<int>> generateMetrix(int n)
{
    vector<vector<int>> res(n, vector<int>(n, 0));
    int startx = 0;
    int starty = 0;
    int loop = n/2;
    int mid = n / 2;
    int offset = 0;
    int count = 1;

    while(loop--) 
    {
        int i = startx;
        int j = starty;

        for (;j<starty + n - offset; j++)
        {
            res[i][j] = count++;
        }

        for (;i<startx+n - offset; i++)
        {
            res[i][j] = count++;
        }

        for (;j > startx;j--)
        {
            res[i][j] = count++;
        }

        for (;i > starty;i--)
        {
            res[i][j] = count++;
        }

        startx++;
        starty++;
        offset += 2;
    }

    if (n % 2) {
        res[mid][mid] = count;
    }

    return res;
}

void PrintMetrix(vector<vector<int>> metrix)
{
    auto it = metrix.begin();
    auto itEnd = metrix.end();

    for (;it != itEnd;++it)
    {
        auto iter = it->begin();
        auto iterEnd = it->end();
        for (;iter != iterEnd;++iter)
        {
            printf("%d ", *iter);
        }
        printf("\n");
    }
}

int main()
{
    auto res = generateMetrix(3);
    PrintMetrix(res);
    return 1;
}

