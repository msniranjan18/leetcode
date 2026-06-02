package main

import (
	"fmt"
)

func main() {

	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(s))
}

func maxArea(height []int) int {
    i,j:=0,len(height)-1
    maxArea := 0
    for i<j {
        h := min(height[i], height[j]) // take the min height as water can go to the min height bar
        w := j-i // distance between two pole note that we are not using j-i+1
        area := h*w // calculate the area
        maxArea = max(maxArea, area) // check for the maximum
        if height[i] < height[j] { // move i as height i is less and we might find a bigger hieght moving i keep the max height of j
            i++
        } else {
            j-- // move j is height j is less then height i
        }
        
    }
    return maxArea
}
