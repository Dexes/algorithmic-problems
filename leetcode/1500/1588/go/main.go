package main

func sumOddLengthSubarrays(arr []int) int {
	length, result, halfLen := len(arr), 0, len(arr)/2
	for i := 0; i < halfLen; i++ {
		result += ((length-i)*(i+1) + 1) / 2 * (arr[i] + arr[length-i-1])
	}

	if len(arr)&1 == 1 {
		result += ((halfLen+1)*(halfLen+1) + 1) / 2 * arr[halfLen]
	}

	return result
}
