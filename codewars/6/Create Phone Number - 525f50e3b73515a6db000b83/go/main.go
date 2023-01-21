package kata

func CreatePhoneNumber(numbers [10]uint) string {
	return string([]byte{
		'(', digit(numbers[0]), digit(numbers[1]), digit(numbers[2]), ')', ' ',
		digit(numbers[3]), digit(numbers[4]), digit(numbers[5]), '-',
		digit(numbers[6]), digit(numbers[7]), digit(numbers[8]), digit(numbers[9]),
	})
}

func digit(x uint) byte {
	return byte(x) + '0'
}
