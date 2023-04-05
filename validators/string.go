package validators

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrLenIsInvalid          = errors.New("len of string field is invalid")
	ErrLenIsLessThenMin      = errors.New("len of string field is less then min")
	ErrLenIsBiggerThenMax    = errors.New("len of string field is bigger then max")
	ErrCantFindStringInArray = errors.New("cant find string field value in \"in\" validator tag array")
)

// IsStringFieldValid return nil if string field of struct is valid. And an error in opposite case.
func IsStringFieldValid(value string, tag string) error {
	splitted := strings.Split(tag, ":")
	validator := splitted[0]
	arguments := splitted[1]
	switch validator {
	case "len":
		expectedLen, _ := strconv.Atoi(arguments)
		if len(value) != expectedLen {
			return ErrLenIsInvalid
		}
	case "min":
		minExpectedLen, _ := strconv.Atoi(arguments)
		if len(value) < minExpectedLen {
			return ErrLenIsLessThenMin
		}
	case "max":
		maxExpectedLen, _ := strconv.Atoi(arguments)
		if len(value) > maxExpectedLen {
			return ErrLenIsBiggerThenMax
		}
	case "in":
		if !IsFieldValueInArray(strings.Split(arguments, ","), value) {
			return ErrCantFindStringInArray
		}
	}
	return nil
}

// IsFieldValueInArray checks if target value belongs to array of values.
func IsFieldValueInArray(inputArr []string, target string) bool {
	for _, val := range inputArr {
		if target == val {
			return true
		}
	}
	return false
}
