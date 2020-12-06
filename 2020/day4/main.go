package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var fieldValidations = map[string](func(string) bool){
	"byr": validateYear(1920, 2002),
	"iyr": validateYear(2010, 2020),
	"eyr": validateYear(2020, 2030),
	"hgt": validateHeight(),
	"hcl": validateHairColor(),
	"ecl": validateEyeColor(),
	"pid": validatePassportID(),
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	passports := strings.Split(string(content), "\n\n")

	passportsWithExptectedRequiredFields := getPassportsWithExpectedRequiredFields(passports)
	passportsWithExtededValidation := getPassportsWithExtendedValidation(passports)

	fmt.Println("passports with expected requiredFields:", passportsWithExptectedRequiredFields)
	fmt.Println("passports with extended validation rules:", passportsWithExtededValidation)
}

func getFields(passwort string) (fieldMap map[string]string) {
	rows := strings.Split(passwort, "\n")
	fieldMap = make(map[string]string)
	for _, r := range rows {
		fields := strings.Split(r, " ")
		for _, f := range fields {
			key := strings.Split(f, ":")[0]
			val := strings.Split(f, ":")[1]
			fieldMap[key] = val
		}
	}
	return fieldMap
}

func getPassportsWithExpectedRequiredFields(passports []string) int {
	validPassports := len(passports)
	for _, p := range passports {
		fields := getFields(p)
		for requiredFieldName := range fieldValidations {
			_, ok := fields[requiredFieldName]
			if !ok {
				validPassports--
				break
			}
		}
	}
	return validPassports
}

func getPassportsWithExtendedValidation(passports []string) int {
	validPassports := len(passports)
	for _, p := range passports {
		fields := getFields(p)
		for requiredFieldName, validate := range fieldValidations {
			val, ok := fields[requiredFieldName]
			if !ok || !validate(val) {
				validPassports--
				break
			}
		}
	}
	return validPassports
}

/**
 * Validation Rules
 */
func validateYear(min, max int) func(string) bool {
	return func(val string) bool {
		if len(val) != 4 {
			return false
		}
		intVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("error while validating")
		}
		if intVal < min {
			return false
		}
		if intVal > max {
			return false
		}
		return true
	}
}

func validateHeight() func(string) bool {
	return func(val string) bool {
		unit := val[len(val)-2:]
		num, _ := strconv.Atoi(val[:len(val)-2])
		switch unit {
		case "in":
			if num < 59 || num > 76 {
				return false
			}
			return true
		case "cm":
			if num < 150 || num > 193 {
				return false
			}
			return true
		default:
			return false
		}
	}
}

func validateHairColor() func(string) bool {
	return func(val string) bool {
		re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		return re.Match([]byte(val))
	}
}

func validateEyeColor() func(string) bool {
	return func(val string) bool {
		for _, s := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if val == s {
				return true
			}
		}
		return false
	}
}

func validatePassportID() func(string) bool {
	return func(val string) bool {
		re := regexp.MustCompile(`^[0-9]{9}$`)
		return re.Match([]byte(val))
	}
}
