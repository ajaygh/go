package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const tagName = "validate"

var emailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

//generic data validator
type Validator interface {
	Validate(interface{}) (bool, error)
}

type DefaultValidator struct {
}

func (defaultValidator DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

type StringValidator struct {
	Min int
	Max int
}

func (v StringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))

	if l == 0 {
		return false, fmt.Errorf("cannot be blank")
	}

	if l < v.Min {
		return false, fmt.Errorf("should be atleast %v characters long", v.Min)
	}
	if l > v.Max {
		return false, fmt.Errorf("should be atmost %v characters long", v.Max)
	}
	return true, nil
}

//Number validator performs numerical validator
type NumberValidator struct {
	Min int
	Max int
}

func (v NumberValidator) Validate(val interface{}) (bool, error) {
	l := val.(int)

	if l < v.Min {
		return false, fmt.Errorf("should be greater than %v", v.Min)
	}

	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf("should be less than %v", v.Max)
	}

	return true, nil
}

type EmailValidator struct {
}

func (v EmailValidator) Validate(val interface{}) (bool, error) {
	if !emailRe.MatchString(val.(string)) {
		return false, fmt.Errorf("Not a valid email address")
	}

	return true, nil
}

//Return validator struct corresponding to validator type
func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "number":
		validator := NumberValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "string":
		validator := StringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "email":
		validator := EmailValidator{}
		return validator
	}

	return DefaultValidator{}
}

//Perform actual data validation
func ValidateStruct(s interface{}) []error {
	errs := []error{}

	v := reflect.ValueOf(s)
	fmt.Println("value ", v)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagName)

		//skip if the tag is not defined
		if tag == "" || tag == "-" {
			continue
		}

		//Get a validator that corresponds to a tag
		validator := getValidatorFromTag(tag)

		//perform validation
		valid, err := validator.Validate(v.Field(i).Interface())

		//Append errors to results
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}
	return errs
}

type User struct {
	Id    int    `validate:"number,min=1,max=1000"`
	Name  string `validate:"string,min=2,max=10"`
	Bio   string `validate:"string"`
	Email string `validate:"email,required"`
}

func main() {
	user := User{
		Id:    0,
		Name:  "tommmmmmmmmmmmmm",
		Bio:   "",
		Email: "tomgmail.com",
	}

	t := reflect.TypeOf(user)
	fmt.Println("Type: ", t.Name())
	fmt.Println("Kind: ", t.Kind())

	//iterate over all fields and read tags
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		fmt.Printf("%d. %v (%v), tag: '%v' \n", i+1, field.Name, field.Type.Name(), tag)
	}

	//Validate
	fmt.Println("Errors:")
	for i, err := range ValidateStruct(user) {
		fmt.Printf("\t%d. %s\n", i+1, err.Error())
	}
}
