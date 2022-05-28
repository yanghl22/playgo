package algs

func isPalindrome(num int) bool {
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}
	var reversedNum int

	for num > reversedNum {
		reversedNum = reversedNum*10 + num%10
		num = num / 10
	}
	return reversedNum == num || reversedNum/10 == num
}
