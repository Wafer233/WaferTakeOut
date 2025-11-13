package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/domain"
)

type CategoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (svc *CategoryService) Create(ctx context.Context, dto *AddCategoryDTO, curId int64) error {

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

func (svc *CategoryService) Delete(ctx context.Context, id int64) error {
	err := svc.repo.Delete(ctx, id)
	return err
}

func (svc *CategoryService) Update(ctx context.Context, dto *EditCategoryDTO, curId int64) error {

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

func (svc *CategoryService) UpdateStatus(ctx context.Context, id int64, status int, curId int64) error {

	entity := &domain.Category{
		ID:         id,
		Status:     status,
		UpdateTime: time.Now(),
		UpdateUser: curId,
	}
	err := svc.repo.UpdateStatus(ctx, entity)
	return err
}

func (svc *CategoryService) FindPage(ctx context.Context, dto *PageDTO) (PageVO, error) {

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

func (svc *CategoryService) FindByType(ctx context.Context, curType int) (GetsTypedVO, error) {

	entities, err := svc.repo.FindByType(ctx, curType)
	if err != nil {
		return nil, err
	}

	vo := make(GetsTypedVO, len(entities))
	for index, record := range vo {
		record.ID = entities[index].ID
		record.Type = entities[index].Type
		record.Name = entities[index].Name
		record.Sort = entities[index].Sort
		record.Status = entities[index].Status
		record.CreateTime = entities[index].CreateTime.Format("2006-01-02 15:04")
		record.UpdateTime = entities[index].UpdateTime.Format("2006-01-02 15:04")
		record.CreateUser = entities[index].CreateUser
		record.UpdateUser = entities[index].UpdateUser
		vo[index] = record
	}

	return vo, nil

}
