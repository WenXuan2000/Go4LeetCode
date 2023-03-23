package main

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		// 由于有重复的数字，所以无法判断那半区间是有序或无序的，这个时候重新更新mid
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
			//	怎么判断区间有序？mid比左端>=时，说明左边有序
		} else if nums[l] <= nums[mid] {
			// 左侧有序 升序
			if nums[l] <= target && target < nums[mid] {
				//说明target 在左边
				r = mid - 1
			} else {
				//说明target在右边-无序
				l = mid + 1
			}
		} else {
			// 右侧有序 升序
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}

func main() {
	nums := []int{0}
	target := 6
	//当一侧无序时另一侧一定是有序的
	println(search(nums, target))
}
