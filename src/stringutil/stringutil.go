package stringutil

//FullName :Returns fullname by concatenating first and last name
func FullName(firstName, lastName string) (string, int) {
	fullName := firstName + " " + lastName
	length := len(fullName)
	return fullName, length
}

//FullNameNakedReturn :Returns fullname by concatenating first and last name but here, variable
//names have been mentioned in the return type itself and therefore need not be explicitly returned
func FullNameNakedReturn(firstName, lastName string) (fullName string, length int) {
	fullName = firstName + " " + lastName
	length = len(fullName)
	return
}
