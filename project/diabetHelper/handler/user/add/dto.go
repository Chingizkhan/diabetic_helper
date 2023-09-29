package add

import (
	pb "diabetHelperTelegramBot/proto/diabetHelper"
	"errors"
)

type DTO struct {
	id           int64
	username     string
	firstName    string
	lastName     string
	languageCode string
}

func (dto *DTO) validate(req *pb.AddUserRequest) error {
	if req.Id != 0 {
		dto.id = req.Id
	} else {
		return errors.New("user id can not be empty")
	}
	if req.Username != "" {
		dto.username = req.Username
	}
	if req.FirstName != "" {
		dto.firstName = req.FirstName
	}
	if req.LastName != "" {
		dto.lastName = req.LastName
	}
	if req.LanguageCode != "" {
		dto.languageCode = req.LanguageCode
	}
	return nil
}
