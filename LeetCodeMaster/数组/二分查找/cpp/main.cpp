#include <iostream>
#include <vector>

using namespace std;

int search(vector<int>& nums, int target) 
{
    int left = 0;
    int right = nums.size() - 1;

    while (left <= right)
    {
        int middle = left + ((right - left) / 2);
        if (nums[middle] == target)
        {
            return middle;
        }

        if (nums[middle] > target)
        {
            right = middle - 1;    
        } 
        else 
        {
            left = middle + 1;
        }
    }

    return -1;
}

int main(int argc ,char ** argv)
{
    vector<int> nums = {-1,0,3,5,9,12};
    int target = 2;

    int idx = search(nums, target);
    printf("index is %d\n", idx);
    return 1;
}