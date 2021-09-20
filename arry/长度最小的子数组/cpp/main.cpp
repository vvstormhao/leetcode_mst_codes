#include <iostream>
#include <vector>

using namespace std;

void printWnd(vector<int> &nums, int start, int end)
{
    printf("[");
    for (int i = start; i <= end; ++i)
    {
        printf("%d ", nums[i]);
    }

    printf("]\n");
}

int minSubArry(vector<int> &nums, int sum)
{
    int minCount = 0;
    int start = 0;
    int end = 0;
    int curSum = nums[start];

    for (;end < nums.size();)
    {
        printWnd(nums, start, end);
        printf("curSum = %d \n", curSum);
        if (nums[start] >= sum)
        {
            minCount = 1;
            return minCount;
        }

        if (curSum < sum)
        {
            curSum += nums[++end];
            continue;
        } else {
            minCount = end - start + 1;
            curSum -= nums[start++];
        }
        
        printf("minCount = %d\n", minCount);
    }

    return minCount;
}

int main(int argc, char ** argv)
{
    vector<int> nums = {7,1,1,1,1,1,1,1};
    int count = minSubArry(nums, 4);
    printf("%d\n", count);
}