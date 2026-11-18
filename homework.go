package homework01

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}	

	// TODO: implement
	m := make(map[int]int)
	for _,v := range nums{
		m[v]++
	}
	for key, value := range m {
		if value==1 {
			return key
		}
	}
	return 0
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// TODO: implement
	//回文数指的是对称数，即正反颠倒数字仍旧一致
	//思路，将数字拆分为数组，然后首位相加除2仍旧是自己，直到算尽
	if x < 0 {
		return false
	}
	tmp := x
	rev :=0
	for x > 0 {
		rev = rev * 10 + x % 10
		x = x/10
	}
	return tmp==rev
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
    var slice []byte  = make([]byte,0)
    for i := 0;i<len(s);i++{
        str := s[i]
        switch str {
            case '(','[','{':
                slice = append(slice,str)
            case ')':
                if len(slice) == 0  ||  slice[len(slice)-1] != '('{
                    return false
                }
                slice = slice[:len(slice)-1]
            case ']':
                if len(slice) == 0  ||  slice[len(slice)-1] != '['{
                    return false
                }
                slice = slice[:len(slice)-1]
            case '}':
                if len(slice) == 0  ||  slice[len(slice)-1] != '{'{
                    return false
                }
                slice = slice[:len(slice)-1]
        }
    }
    return len(slice)==0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for _, v := range strs[1:] {   // 从第 2 个开始比即可
        i := 0
        for i < len(prefix) && i < len(v) && prefix[i] == v[i] {
            i++
        }
        prefix = prefix[:i]        // 直接截取公共部分
        if prefix == "" {          // 提前退出
            return ""
        }
    }
    return prefix
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	//因为考虑到会进一，所以这里需要逆序遍历,同时需要考虑到进位的问题
    for i := len(digits) - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++          // 不是 9，加 1 后不进位，直接结束
            return digits
        }
        digits[i] = 0            // 是 9，变 0，继续往前看
    }
    // 走到这里说明全是 9，需要多一位
    return append([]int{1}, digits...)
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	//使用快慢指针，慢指针存有效位有效值，快指针快读做比对
    //移位的位置存储快指针的值
    //另外，默认是有序的
    if len(nums)==0 {
        return 0
    }
    slow :=0
    for fast:=1;fast<len(nums);fast++{
        if(nums[slow]!=nums[fast]){
            slow++
            nums[slow] = nums[fast]
        }
    }
    return slow+1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// TODO: implement
	//数组排序
	n :=len(intervals)
	for i :=0;i<n-1;i++{
		swapped := false
		for j:= 0;j<n-1-i;j++{
			if intervals[j][0] > intervals[j+1][0] {
                intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
                swapped = true
            }
		}
		if !swapped {   // 已有序，提前退出
            break
        }
	}
	ranArr := make([][]int,0)
	for _,v := range intervals {
		if(len(ranArr)==0){
			ranArr = append(ranArr,v)
			continue
		}
		maxRng := ranArr[len(ranArr)-1]
		tmpRng := v
		if(tmpRng[1]<maxRng[0]){
			ranArr = append([][]int{tmpRng},ranArr...)
		}else if(tmpRng[0]>maxRng[1]){
			ranArr = append(ranArr,tmpRng)
		}else {
			if(tmpRng[0]<maxRng[0]){
				maxRng[0] = tmpRng[0]
			}
			if(tmpRng[1]>maxRng[1]){
				maxRng[1] = tmpRng[1]
			}
		}
	}
	
	return ranArr
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// TODO: implement
	m := make(map[int]int)
    for i,v :=range nums{
        if j, ok := m[target-v]; ok{
            return []int{j,i}
        }
        m[v] = i
    }
    return []int{}
}
