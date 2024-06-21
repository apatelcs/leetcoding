package main

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	sum := 0
	gains := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			sum += customers[i]
		}

		if i+minutes <= len(customers) {
			curr := 0
			for j := i; j < i+minutes; j++ {
				curr += grumpy[j] * customers[j]
			}
			if curr > gains {
				gains = curr
			}
		}
	}

	return sum + gains
}
