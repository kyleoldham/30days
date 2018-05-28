package twosum

func twoSum(input []int, target int) []int {
	theMap := make(map[int]int, len(input))
	for i, num1 := range input {
		if j, num2 := theMap[target-num1]; num2 {
			return []int{j, i}
		}
		theMap[num1] = i
	}
	return nil
}
