#include <iostream>
#include <vector>

using namespace std;


int remove(vector<int> &nums, int target)
{
    int slowIndex = 0;
    for (int fastIndex = 0; fastIndex < nums.size();fastIndex++)
    {
        if (target != nums[fastIndex])
        {
            nums[slowIndex++] = nums[fastIndex];
        }
    }

    return slowIndex;
}

int main(int argc, char ** argv)
{
    vector<int> nums = {0,1,2,2,3,0,4,2};
    int count = remove(nums, 2);
    for (int i = 0; i < count; ++i)
    {
        printf("%d\n", nums[i]);
    }

    return 1;
}
