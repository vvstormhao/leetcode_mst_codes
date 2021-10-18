#include <iostream>
#include <vector>

using namespace std;

vector<int> squre(vector<int>& nums)
{
    int i = 0;
    int j = nums.size() - 1;

    vector<int> ret(nums.size());
    int k = j;
    for (;i <= j;)
    {
        if (nums[i] * nums[i] > nums[j] * nums[j])
        {
            ret[k--] = nums[i] * nums[i];
            i++;
        }
        else
        {
            ret[k--] = nums[j] * nums[j];
            j--;
        }
    }

    return ret;
}

int main(int argc, char ** argv)
{
    vector<int> nums = {-4,-1,0,3,10};
    vector<int>  ret = squre(nums);

    for (int i = 0; i < ret.size(); i++)
    {
        printf("%d\n", ret[i]);
    }

    return 1;
}