package application

import (
	"context"
	cate "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/domain"
	dish "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/flavor"
	"strconv"
	"time"
)

type DishService struct {
	dishRepo dish.DishRepository
	cateRepo cate.CategoryRepository
	flavRepo flavor.FlavorRepository
}

func NewDishService(dishRepo dish.DishRepository,
	cateRepo cate.CategoryRepository,
	flavRepo flavor.FlavorRepository) *DishService {
	return &DishService{
		dishRepo: dishRepo,
		cateRepo: cateRepo,
		flavRepo: flavRepo,
	}
}

func (svc *DishService) Insert(ctx context.Context, dto *DishDTO, curId int64) error {

	entity := &domain.Dish{
		Name:        dto.Name,
		CategoryId:  dto.CategoryId,
		Price:       dto.Price,
		Image:       dto.Image,
		Description: dto.Description,
		Status:      dto.Status,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		CreateUser:  curId,
		UpdateUser:  curId,
	}
	err := svc.dishRepo.Insert(ctx, entity)
	if err != nil {
		return err
	}

	if len(dto.Flavors) != 0 {
		entities := make([]*flavor.Flavor, len(dto.Flavors))
		for index, value := range entities {
			value = &flavor.Flavor{}
			value.DishId = entity.Id
			value.Name = dto.Flavors[index].Name
			value.Value = dto.Flavors[index].Value

			entities[index] = value
		}

		err = svc.flavRepo.Inserts(ctx, entities)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *DishService) UpdateDish(ctx context.Context,
	dto *DishDTO, curID int64) error {

	entity := &domain.Dish{
		Id:          dto.ID,
		Name:        dto.Name,
		CategoryId:  dto.CategoryId,
		Price:       dto.Price,
		Image:       dto.Image,
		Description: dto.Description,
		Status:      dto.Status,
		UpdateTime:  time.Now(),
		UpdateUser:  curID,
	}
	err := svc.dishRepo.UpdateById(ctx, entity)
	if err != nil {
		return err
	}

	if len(dto.Flavors) == 0 {
		return nil
	}

	flavorsDTO := dto.Flavors

	flavors := make([]*flavor.Flavor, len(flavorsDTO))
	for i, f := range flavorsDTO {
		flavors[i] = &flavor.Flavor{
			Id:     f.Id,
			DishId: entity.Id,
			Name:   f.Name,
			Value:  f.Value,
		}
	}
	err = svc.flavRepo.UpdatesByDishId(ctx, flavors)
	return nil
}

func (svc *DishService) PageQuery(ctx context.Context, dto *PageDTO) (PageVO, error) {

	name := dto.Name
	categoryId := dto.CategoryId
	page := dto.Page
	pageSize := dto.PageSize
	status, _ := strconv.Atoi(dto.Status)
	if dto.Status == "" {
		status = -1
	}

	dishes, total, err := svc.dishRepo.GetsPaged(ctx, name, categoryId, status, page, pageSize)
	if err != nil {
		return PageVO{}, err
	}

	categoryNames := make([]string, 0)
	for index, _ := range dishes {
		category, er := svc.cateRepo.FindById(ctx, dishes[index].CategoryId)
		if er != nil {
			return PageVO{}, er
		}
		categoryNames = append(categoryNames, category.Name)
	}

	records := make([]Record, len(dishes))
	for index, record := range records {
		record.ID = dishes[index].Id
		record.Name = dishes[index].Name
		record.CategoryId = dishes[index].CategoryId
		record.Price = dishes[index].Price
		record.Image = dishes[index].Image
		record.Description = dishes[index].Description
		record.Status = dishes[index].Status
		record.UpdateTime = dishes[index].UpdateTime.Format("2006-01-02 15:04")
		record.CategoryName = categoryNames[index]

		records[index] = record
	}

	vo := PageVO{
		Total:   total,
		Records: records,
	}
	return vo, nil
}

func (svc *DishService) GetDishCategory(ctx context.Context, id int64) ([]*Record, error) {

	dishes, err := svc.dishRepo.GetsByCategoryId(ctx, id)
	if err != nil || len(dishes) == 0 {
		return nil, err
	}

	records := make([]*Record, len(dishes))

	category, err := svc.cateRepo.FindById(ctx, dishes[0].CategoryId)
	if err != nil {
		return nil, err
	}

	for index, d := range dishes {
		records[index] = &Record{
			ID:           d.Id,
			Name:         d.Name,
			CategoryId:   d.CategoryId,
			Price:        d.Price,
			Image:        d.Image,
			Description:  d.Description,
			Status:       d.Status,
			UpdateTime:   d.UpdateTime.Format("2006-01-02 15:04"),
			CategoryName: category.Name,
		}
	}
	return records, nil
}

func (svc *DishService) GetDishId(ctx context.Context, id int64) (DishVO, error) {

	dish, err := svc.dishRepo.GetById(ctx, id)
	if err != nil {
		return DishVO{}, err
	}
	category, err := svc.cateRepo.FindById(ctx, dish.CategoryId)
	if err != nil {
		return DishVO{}, err
	}

	flavors, err := svc.flavRepo.GetsByDishId(ctx, dish.Id)
	if err != nil {
		return DishVO{}, err
	}

	flavorVOs := make([]Flavor, len(flavors))
	for index, flavor := range flavors {
		flavorVOs[index].DishId = flavor.DishId
		flavorVOs[index].Value = flavor.Value
		flavorVOs[index].Name = flavor.Name
		flavorVOs[index].Id = flavor.Id
	}

	vo := DishVO{
		CategoryId:   dish.CategoryId,
		CategoryName: category.Name,
		Description:  dish.Description,
		Flavors:      flavorVOs,
		ID:           dish.Id,
		Image:        dish.Image,
		Name:         dish.Name,
		Price:        dish.Price,
		Status:       dish.Status,
		UpdateTime:   dish.UpdateTime.Format("2006-01-02 15:04"),
	}
	return vo, nil
}

func (svc *DishService) StatusFlip(ctx context.Context, id int64, status int, curId int64) error {

	entity := &domain.Dish{
		Id:         id,
		Status:     status,
		UpdateTime: time.Now(),
		UpdateUser: curId,
	}
	err := svc.dishRepo.UpdateStatusById(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

func (svc *DishService) DeleteDishes(ctx context.Context, idArr []int64) error {

	err := svc.dishRepo.DeletesById(ctx, idArr)
	if err != nil {
		return err
	}

	return nil
}
