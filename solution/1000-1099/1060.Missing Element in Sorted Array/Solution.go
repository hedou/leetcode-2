func missingElement(nums []int, k int) int {
	missing := func(i int) int {
		return nums[i] - nums[0] - i
	}
	n := len(nums)
	if k > missing(n-1) {
		return nums[n-1] + k - missing(n-1)
	}
	l, r := 0, n-1
	for l < r {
		mid := (l + r) >> 1
		if missing(mid) >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l-1] + k - missing(l-1)
}