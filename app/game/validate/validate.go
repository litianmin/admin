package validate

import (
	"admin/app/game/entity"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// CreateGame 增加游戏的数据验证
func CreateGame(l *entity.NewGame) bool {
	_, err := govalidator.ValidateStruct(l)
	if err != nil {
		return false
	}
	return true
}
