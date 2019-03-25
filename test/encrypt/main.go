package main

import (
	"admin/utils"
	"fmt"
)

func main() {
	str := "Hbk5551412"
	pwd := utils.PwdSha1Encrypt(str)
	fmt.Println(pwd)
}
