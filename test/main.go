package main

import (
	"admin/utils"
	"fmt"
)

func main() {
	str := "Hbk5551412"
	encryptStr := utils.PwdSha1Encrypt(str)
	fmt.Println(encryptStr)
}
