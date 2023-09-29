package find_sugar_level

import (
	"diabetHelperTelegramBot/diabetHelper/storage"
	pb "diabetHelperTelegramBot/proto/diabetHelper"
	"time"
)

type DTO struct {
	UserId          int64
	Value           string
	CreatedAtStart  time.Time
	CreatedAtFinish time.Time
	pagination      *pb.Pagination
}

func (dto *DTO) validate(req *pb.FindSLRequest) (err error) {
	if req.UserId != 0 {
		dto.UserId = req.UserId
	}

	if req.Value != "" {
		dto.Value = req.Value
	}

	if req.CreatedAtStart != 0 {
		dto.CreatedAtStart = time.Unix(req.CreatedAtStart, 0)
	}

	if req.CreatedAtFinish != 0 {
		dto.CreatedAtFinish = time.Unix(req.CreatedAtFinish, 0)
	}

	if req.Pagination == nil {
		dto.pagination = &pb.Pagination{}
	} else {
		dto.pagination = &pb.Pagination{
			Page:   req.Pagination.Page,
			Limit:  req.Pagination.Limit,
			Offset: req.Pagination.Offset,
		}
	}

	return nil
}

func mapSLDtoToFindDto(dto *DTO) storage.SugarLevelFindDTO {
	return storage.SugarLevelFindDTO{
		UserId:          dto.UserId,
		Value:           dto.Value,
		CreatedAtStart:  dto.CreatedAtStart,
		CreatedAtFinish: dto.CreatedAtFinish,
	}
}
