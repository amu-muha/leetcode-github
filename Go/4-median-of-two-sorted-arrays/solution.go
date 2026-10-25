
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // Always binary search the smaller array
    if len(nums1) > len(nums2) {
        return findMedianSortedArrays(nums2, nums1)
    }

    m, n := len(nums1), len(nums2)
    left, right := 0, m

    for left <= right {
        partitionX := (left + right) / 2
        partitionY := (m+n+1)/2 - partitionX

        maxLeftX := -int(^uint(0)>>1) - 1
        minRightX := int(^uint(0) >> 1)
        maxLeftY := -int(^uint(0)>>1) - 1
        minRightY := int(^uint(0) >> 1)

        if partitionX > 0 {
            maxLeftX = nums1[partitionX-1]
        }
        if partitionX < m {
            minRightX = nums1[partitionX]
        }
        if partitionY > 0 {
            maxLeftY = nums2[partitionY-1]
        }
        if partitionY < n {
            minRightY = nums2[partitionY]
        }

        if maxLeftX <= minRightY && maxLeftY <= minRightX {
            if (m+n)%2 == 0 {
                return (float64(max(maxLeftX, maxLeftY)) +
                    float64(min(minRightX, minRightY))) / 2.0
            }
            return float64(max(maxLeftX, maxLeftY))
        } else if maxLeftX > minRightY {
            right = partitionX - 1
        } else {
            left = partitionX + 1
        }
    }

    return 0.0
}
