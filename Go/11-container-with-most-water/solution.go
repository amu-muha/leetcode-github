func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxWater := 0

    for left < right {
        h := height[left]
        if height[right] < h {
            h = height[right]
        }

        water := h * (right - left)

        if water > maxWater {
            maxWater = water
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxWater
}