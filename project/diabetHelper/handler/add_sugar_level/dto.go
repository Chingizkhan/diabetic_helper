package add_sugar_level

import (
	pb "diabetHelperTelegramBot/proto/diabetHelper"
	"errors"
)

type DTO struct {
	userId int64
	value  string
}

func (dto *DTO) validate(req *pb.AddSLRequest) (err error) {
	if req.UserId == 0 {
		return errors.New("username can not be empty")
	} else {
		dto.userId = req.UserId
	}

	if req.Value == "" {
		return errors.New("value can not be empty")
	} else {
		dto.value = req.Value
	}

	return nil
}
