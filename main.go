package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println(removeDuplicateLetter("bananas"))
	// fmt.Println(isAnagram("Tamat", "Kiamat"))
	fmt.Println(isAnagram("Aku", "Kua"))
	fmt.Println(removeDuplicate([]int{1, 2, 2, 3, 3, 4, 4, 4}))
	words := []string{"this", "is", "a", "kanoha"}
	exclude := []string{"is", "a"}
	fmt.Println(Capitalize(words, exclude))
	fmt.Println(oddBeforeEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(plusOne([]int{1, 4, 8, 9}, 2))
	fmt.Println(sameElement([]int{1, 2, 7, 4, 7, 8}, []int{7, 7, 3, 2, 9}))
	fmt.Println(countFrequentElement([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}))
	fmt.Println(addDigits([]int{9, 2, 7}, []int{1, 3, 5}))
	for i := 0; i <= 4; i++ {
		fmt.Println(sliceOperation([]int{1, 2, 3, 4, 5}, []int{2, 4, 6, 7}, i))
	}
	for i := 0; i <= 2; i++ {
		fmt.Println(sliceFruits(
			[]string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"},
			[]string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"},
			i))
	}

}

func removeDuplicateLetter(words string) string {
	// Quiz 1
	fmt.Println("Quiz 1")
	Hash := make(map[string]int)
	count := 1
	var result string
	for i := 0; i < len(words); i++ {
		if _, ok := Hash[string(words[i])]; ok {
			continue
		}
		Hash[string(words[i])] = count
		result += string(words[i])
	}
	return result
}

func isAnagram(word1 string, word2 string) bool {
	// Quiz 2
	fmt.Println("Quiz 2")
	if len(word1) != len(word2) {
		return false
	}

	temp_rune_first := []rune(strings.ToLower(word1))
	temp_rune_second := []rune(strings.ToLower(word2))

	sort.Slice(temp_rune_first, func(first, last int) bool {
		return temp_rune_first[first] < temp_rune_first[last]
	})
	sort.Slice(temp_rune_second, func(first, last int) bool {
		return temp_rune_second[first] < temp_rune_second[last]
	})
	// fmt.Print(temp_rune_first, temp_rune_second)
	return string(temp_rune_first) == string(temp_rune_second)
}

func Capitalize(words []string, exceptions []string) []string {
	fmt.Println("Quiz 3")
	var Capitalize []string

	// Create a map for exceptions for O(1) lookup time
	ExceptionMap := make(map[string]bool)
	for _, exception := range exceptions {
		ExceptionMap[exception] = true
	}

	for _, PerWord := range words {
		// Check if the word is an exception
		if ExceptionMap[PerWord] {
			Capitalize = append(Capitalize, PerWord)
		} else {
			// Capitalize the first letter of the word
			capitalizedWord := strings.ToUpper(string(PerWord[0])) + PerWord[1:]
			Capitalize = append(Capitalize, capitalizedWord)
		}
	}

	return Capitalize
}

func removeDuplicate(nums []int) []int {
	fmt.Println("Quiz 4")
	var result []int
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			result = append(result, nums[left])
			left = right
			continue
		}
		left = right
	}
	result = append(result, nums[left])
	return result
}

func oddBeforeEven(nums []int) []int {
	fmt.Println("Quiz 5")
	sort.Ints(nums)
	var odd []int
	var even []int
	for _, digit := range nums {
		if digit%2 == 0 && digit != 0 {
			even = append(even, digit)
		} else {
			odd = append(odd, digit)
		}
	}
	odd = append(odd, even...)
	return odd
}

func plusOne(nums []int, add int) any {
	fmt.Println("Quiz 6")
	if add == 0 {
		return false
	}
	size := len(nums)
	// var result []int
	nums[size-1] += add
	for i := size - 1; i > 0; i-- {
		carry_number := nums[i] / 10
		nums[i] %= 10
		nums[i-1] += carry_number
	}

	if nums[0] == 0 {
		carry_number := nums[0] / 10
		nums[0] %= 10
		nums = append([]int{carry_number}, nums...)
	}
	return nums
}

func sameElement(nums1 []int, nums2 []int) []int {
	// nums1 = append(nums1, nums2...)
	fmt.Println("Quiz 7")
	ExactElement := make(map[int]bool)
	for _, nums := range nums1 {
		ExactElement[nums] = true
	}

	result_map := make(map[int]bool)
	for _, nums := range nums2 {
		if ExactElement[nums] {
			result_map[nums] = true
		}
	}

	result := make([]int, 0, len(result_map))

	for num := range result_map {
		result = append(result, num)
	}
	sort.Ints(result)
	return result
}

func sliceOperation(nums1 []int, nums2 []int, typeOperation int) []int {
	fmt.Println("Quiz 8")
	nums1_temp := make(map[int]bool)
	nums2_temp := make(map[int]bool)

	for _, nums := range nums1 {
		nums1_temp[nums] = true
	}

	for _, nums := range nums2 {
		nums2_temp[nums] = true
	}

	result := make([]int, 0)
	switch typeOperation {
	case 1:
		{
			// Difference 1
			for nums := range nums1_temp {
				if nums2_temp[nums] {
					delete(nums1_temp, nums)
				}
			}
			for nums := range nums1_temp {
				result = append(result, nums)
			}
			return result
		}
	case 2:
		{
			// Difference 2
			for nums := range nums2_temp {
				if nums1_temp[nums] {
					delete(nums2_temp, nums)
				}
			}
			for nums := range nums2_temp {
				result = append(result, nums)
			}
			return result
		}
	case 3:
		{
			exact_map := make(map[int]bool)
			for nums := range nums1_temp {
				exact_map[nums] = true
			}
			for nums := range nums2_temp {
				exact_map[nums] = true
			}

			result := make([]int, 0)
			for nums := range exact_map {
				result = append(result, nums)
			}
			return result
		}
	case 4:
		{
			result := make([]int, 0)
			for nums := range nums1_temp {
				if nums2_temp[nums] {
					result = append(result, nums)
				}
			}
			return result
		}
	default:
		break
	}
	return []int{}
}

func countFrequentElement(nums []int) map[int]int {
	fmt.Println("Quiz 9")
	result := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		result[nums[i]]++
	}
	return result
}
func addDigits(nums1 []int, nums2 []int) any {
	fmt.Println("Quiz 10")
	if len(nums1) != len(nums2) {
		return false
	}
	var carry_number int
	carry_number = 0
	// for i := len(nums1) - 1; i > 0; i-- {
	// 	carry_number := nums1[i] + nums2[i]
	// 	nums1[i] = carry_number % 10
	// 	nums1[i-1] += carry_number / 10
	// }
	// if nums1[0] != 0 {
	// 	carry_number = nums1[0] + nums2[0]
	// 	nums1[0] = carry_number % 10
	// 	nums1 = append([]int{carry_number / 10}, nums1...)
	// } else {
	// 	carry_number := nums1[1] + nums2[1]
	// 	nums1[1] = carry_number % 10
	// 	nums1[0] += carry_number / 10
	// }

	for i := len(nums1) - 1; i >= 0; i-- {
		sum := nums1[i] + nums2[i] + carry_number
		nums1[i] = sum % 10
		carry_number = sum / 10
	}
	for carry_number > 0 {
		nums1 = append([]int{carry_number}, nums1...)
		carry_number /= 10
	}
	return nums1
}

func sliceFruits(fruits1 []string, fruits2 []string, operationType int) []string {
	fmt.Println("Quiz 11")
	toMapping := func(words []string) map[string]bool {
		temp := make(map[string]bool)
		for _, wording := range words {
			temp[wording] = true
		}
		return temp
	}
	switch operationType {
	case 1:
		{
			fruits1Map := toMapping(fruits1)
			var result []string
			for _, perWord := range fruits2 {
				if fruits1Map[perWord] {
					result = append(result, perWord)
				}
			}
			return result
		}
	case 2:
		{
			fruits1Map := toMapping(fruits1)
			var result []string
			for _, perWord := range fruits2 {
				if !fruits1Map[perWord] {
					result = append(result, perWord)
				}
			}
			return result
		}
	default:
		break
	}
	return []string{}
}
