package main

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	FirstName string
	LastName  string
}

func main() {
	type Post struct {
		Title    string `valid:"alphanum,required"`
		Message  string `valid:"duck,ascii"`
		AuthorIP string `valid:"ipv4"`
		Date     string `valid:"-"`
	}
	post := &Post{
		Title:    "My Example Post",
		Message:  "duck",
		AuthorIP: "123.234.54.3",
	}

	// Add your own struct validation tags
	govalidator.TagMap["duck"] = govalidator.Validator(func(str string) bool {
		return str == "duck"
	})

	result, err := govalidator.ValidateStruct(post)
	if err != nil {
		println("error: " + err.Error())
	}
	println(result)
}
