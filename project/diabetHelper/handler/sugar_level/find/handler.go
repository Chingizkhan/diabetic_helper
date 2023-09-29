package find

import (
	"context"
	"diabetHelperTelegramBot/diabetHelper/handler/common"
	"diabetHelperTelegramBot/diabetHelper/storage/postgres"
	"diabetHelperTelegramBot/diabetHelper/utils"
	pb "diabetHelperTelegramBot/proto/diabetHelper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

const defaultLimit int32 = 10

func Handle(ctx context.Context, req *pb.FindSLRequest) (*pb.SugarLevels, error) {
	dto := &DTO{}
	if err := dto.validate(req); err != nil {
		return nil, status.Error(codes.Code(http.StatusBadRequest), err.Error())
	}

	repo := postgres.NewSugarLevelStorage()

	findDto := mapSLDtoToFindDto(dto)

	countAll, err := repo.GetCountAllForFind(ctx, findDto)
	if err != nil {
		return nil, err
	}

	perPage := defaultLimit
	if dto.pagination.Limit != 0 {
		perPage = dto.pagination.Limit
	}

	pager := utils.NewPagination(dto.pagination.Page, perPage, dto.pagination.Offset, countAll)
	findDto.Limit = pager.PerPage()
	findDto.Offset = pager.Offset()

	levels, err := repo.Find(ctx, findDto)
	if err != nil {
		return nil, status.Error(codes.Code(http.StatusInternalServerError), err.Error())
	}

	res := make([]*pb.SugarLevel, 0, len(levels))
	for _, sl := range levels {
		res = append(res, common.MapSLToPb(sl))
	}

	return &pb.SugarLevels{
		PagesCount: pager.GetPagesCount(),
		TotalItems: countAll,
		PerPage:    pager.PerPage(),
		Data:       res,
	}, nil
}
