package validate

import (
	"admin/app/user/entity"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Login 验证登陆结构体
func Login(l *entity.LoginAuth) bool {
	_, err := govalidator.ValidateStruct(l)
	if err != nil {
		return false
	}
	return true
}
