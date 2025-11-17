package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/category-service/internal/domain"
)

type CategoryAppService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) *CategoryAppService {
	return &CategoryAppService{
		repo: repo,
	}
}

// Create
func (svc *CategoryAppService) Create(ctx context.Context, dto *CategoryDTO, curId int64) error {

	entity := domain.Category{
		ID:         dto.ID,
		Name:       dto.Name,
		Sort:       dto.Sort,
		Type:       dto.Type,
		Status:     1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		CreateUser: curId,
		UpdateUser: curId,
	}

	err := svc.repo.Create(ctx, &entity)
	return err
}

func (svc *CategoryAppService) Delete(ctx context.Context, id int64) error {
	err := svc.repo.Delete(ctx, id)
	return err
}

func (svc *CategoryAppService) Update(ctx context.Context, dto *CategoryDTO, curId int64) error {

	entity := domain.Category{
		ID:         dto.ID,
		Name:       dto.Name,
		Sort:       dto.Sort,
		UpdateTime: time.Now(),
		UpdateUser: curId,
	}
	err := svc.repo.Update(ctx, &entity)
	if err != nil {
		return err
	}
	return nil
}

func (svc *CategoryAppService) UpdateStatus(ctx context.Context, id int64, status int, curId int64) error {

	entity := &domain.Category{
		ID:         id,
		Status:     status,
		UpdateTime: time.Now(),
		UpdateUser: curId,
	}
	err := svc.repo.UpdateStatus(ctx, entity)
	return err
}

func (svc *CategoryAppService) FindPage(ctx context.Context, dto *PageDTO) (PageVO, error) {

	curName := dto.Name
	page := dto.Page
	pageSize := dto.PageSize
	curType := dto.Type

	entities, total, err := svc.repo.FindPage(ctx, curName, curType, page, pageSize)
	if err != nil {
		return PageVO{}, err
	}
	records := make([]Record, len(entities))

	for index, record := range records {
		record.ID = entities[index].ID
		record.Type = entities[index].Type
		record.Name = entities[index].Name
		record.Sort = entities[index].Sort
		record.Status = entities[index].Status
		record.CreateTime = entities[index].CreateTime.Format("2006-01-02 15:04")
		record.UpdateTime = entities[index].UpdateTime.Format("2006-01-02 15:04")
		record.CreateUser = entities[index].CreateUser
		record.UpdateUser = entities[index].UpdateUser

		records[index] = record
	}

	return PageVO{
		Total:   total,
		Records: records,
	}, nil

}

func (svc *CategoryAppService) FindByType(ctx context.Context, curType int) ([]Record, error) {

	entities, err := svc.repo.FindByType(ctx, curType)
	if err != nil {
		return []Record{}, err
	}

	vo := make([]Record, len(entities))
	for index, entity := range entities {
		vo[index] = Record{
			ID:         entity.ID,
			Type:       entity.Type,
			Name:       entity.Name,
			Sort:       entity.Sort,
			Status:     entity.Status,
			CreateTime: entity.CreateTime.Format("2006-01-02 15:04"),
			UpdateTime: entity.UpdateTime.Format("2006-01-02 15:04"),
			CreateUser: entity.CreateUser,
			UpdateUser: entity.UpdateUser,
		}
	}

	return vo, nil

}
