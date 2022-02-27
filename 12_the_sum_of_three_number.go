// 12.三数之和
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
// 使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

// 排序 + 双指针
// https://leetcode-cn.com/problems/3sum/solution/san-shu-zhi-he-by-leetcode-solution/

func threeSum(nums []int) [][]int {
    n := len(nums)
    sort.Ints(nums)
    ans := make([][]int, 0)
 
    // 枚举 a
    for first := 0; first < n; first++ {
        // 需要和上一次枚举的数不相同
        if first > 0 && nums[first] == nums[first - 1] {
            continue
        }
        // c 对应的指针初始指向数组的最右端
        third := n - 1
        target := -1 * nums[first]
        // 枚举 b
        for second := first + 1; second < n; second++ {
            // 需要和上一次枚举的数不相同
            if second > first + 1 && nums[second] == nums[second - 1] {
                continue
            }
            // 需要保证 b 的指针在 c 的指针的左侧
            for second < third && nums[second] + nums[third] > target {
                third--
            }
            // 如果指针重合，随着 b 后续的增加
            // 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
            if second == third {
                break
            }
            if nums[second] + nums[third] == target {
                ans = append(ans, []int{nums[first], nums[second], nums[third]})
            }
        }
    }
    return ans
}


class Solution {
	public:
		vector<vector<int>> threeSum(vector<int>& nums) {
			int n = nums.size();
			sort(nums.begin(), nums.end());
			vector<vector<int>> ans;
			// 枚举 a
			for (int first = 0; first < n; ++first) {
				// 需要和上一次枚举的数不相同
				if (first > 0 && nums[first] == nums[first - 1]) {
					continue;
				}
				// c 对应的指针初始指向数组的最右端
				int third = n - 1;
				int target = -nums[first];
				// 枚举 b
				for (int second = first + 1; second < n; ++second) {
					// 需要和上一次枚举的数不相同
					if (second > first + 1 && nums[second] == nums[second - 1]) {
						continue;
					}
					// 需要保证 b 的指针在 c 的指针的左侧
					while (second < third && nums[second] + nums[third] > target) {
						--third;
					}
					// 如果指针重合，随着 b 后续的增加
					// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
					if (second == third) {
						break;
					}
					if (nums[second] + nums[third] == target) {
						ans.push_back({nums[first], nums[second], nums[third]});
					}
				}
			}
			return ans;
		}
	};