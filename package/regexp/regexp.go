package regexp

import "regexp"

func CheckValidAccOrPassword(param string) bool {
	var digit, _ = regexp.Compile(`\d`)
	var word, _ = regexp.Compile("[a-z]")
	var valid, _ = regexp.Compile(`(?m)^[0-9a-z]{6,20}$`)
	if !digit.MatchString(param) || !word.MatchString(param) || !valid.MatchString(param) {
		return false
	}
	return true
}

func CheckValidWithdrawalPassword(param string) bool {
	var valid, _ = regexp.Compile(`(?m)^[0-9]{4}$`)
	return valid.MatchString(param)
}

func CheckBankCard(param string) bool {
	var valid, _ = regexp.Compile(`(?m)^[0-9]{16,20}$`)
	return valid.MatchString(param)
}

func CheckName(param string) bool {
	var valid, _ = regexp.Compile(`(?m)^[\x{4e00}-\x{9fa5}]+$`)
	return valid.MatchString(param)
}