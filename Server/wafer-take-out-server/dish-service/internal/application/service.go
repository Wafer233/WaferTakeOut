package application

import (
	"context"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/internal/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/pkg/ai"

	"github.com/jinzhu/copier"
)

type DishService struct {
	repo domain.DishRepository
	svc  *rpc.CategoryService
}

func NewDishService(dishRepo domain.DishRepository, svc *rpc.CategoryService) *DishService {
	return &DishService{
		repo: dishRepo,
		svc:  svc,
	}
}

func (svc *DishService) Create(ctx context.Context, dto *DishDTO, curId int64) error {

	dishEntity := &domain.Dish{
		Id:          dto.ID, // 现在还没有值
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

	if dto.Description != "" {
		newDescription, _ := ai.GetDescriptionRanking(dto.Description)
		dishEntity.Description = newDescription
	}

	flavorEntities := make([]*domain.Flavor, len(dto.Flavors))
	for index, value := range dto.Flavors {
		flavorEntities[index] = &domain.Flavor{
			Id:     value.Id,     //现在还没有值
			DishId: value.DishId, //现在还没有值
			Name:   value.Name,
			Value:  value.Value,
		}
	}

	err := svc.repo.Create(ctx, dishEntity, flavorEntities)

	return err
}

func (svc *DishService) Update(ctx context.Context, dto *DishDTO, curID int64) error {

	dishEntity := &domain.Dish{
		Id:          dto.ID, // 这个是必须的
		Name:        dto.Name,
		CategoryId:  dto.CategoryId,
		Price:       dto.Price,
		Image:       dto.Image,
		Description: dto.Description,
		Status:      dto.Status,
		UpdateTime:  time.Now(),
		UpdateUser:  curID,
	}

	if dto.Description != "" {
		newDescription, _ := ai.GetDescriptionRanking(dto.Description)
		dishEntity.Description = newDescription
	}

	flavors := make([]*domain.Flavor, len(dto.Flavors))
	for index, value := range dto.Flavors {
		flavors[index] = &domain.Flavor{
			Id:     value.Id,
			DishId: dishEntity.Id,
			Name:   value.Name,
			Value:  value.Value,
		}
	}

	err := svc.repo.Update(ctx, dishEntity, flavors)

	return err
}

func (svc *DishService) FindPage(ctx context.Context, dto *PageDTO) (PageVO, error) {

	name := dto.Name
	categoryId := dto.CategoryId
	page := dto.Page
	pageSize := dto.PageSize
	status, _ := strconv.Atoi(dto.Status)
	if dto.Status == "" {
		status = -1
	}

	dishes, total, err := svc.repo.FindPage(ctx, name, categoryId, status, page, pageSize)
	if err != nil {
		return PageVO{}, err
	}

	categoryNames := make([]string, 0)
	for index, _ := range dishes {
		curName, er := svc.svc.FindNameById(ctx, dishes[index].CategoryId)
		if er != nil {
			return PageVO{}, er
		}
		categoryNames = append(categoryNames, curName)
	}

	records := make([]Record, len(dishes))
	for index, record := range records {
		record.Id = dishes[index].Id
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

func (svc *DishService) FindByCategoryId(ctx context.Context, cid int64) ([]Record, error) {

	dishes, err := svc.repo.FindByCategoryId(ctx, cid)

	// 防止溢出
	if len(dishes) == 0 {
		return []Record{}, nil
	}

	records := make([]Record, len(dishes))

	curName, err := svc.svc.FindNameById(ctx, dishes[0].CategoryId)
	if err != nil {
		return nil, err
	}

	for index, d := range dishes {
		records[index] = Record{
			Id:           d.Id,
			Name:         d.Name,
			CategoryId:   d.CategoryId,
			Price:        d.Price,
			Image:        d.Image,
			Description:  d.Description,
			Status:       d.Status,
			UpdateTime:   d.UpdateTime.Format("2006-01-02 15:04"),
			CategoryName: curName,
		}
	}
	return records, nil
}

func (svc *DishService) FindById(ctx context.Context, id int64) (DishVO, error) {

	dishEntity, flavors, err := svc.repo.FindById(ctx, id)
	if err != nil {
		return DishVO{}, err
	}

	flavorVOs := make([]Flavor, len(flavors))
	for index, flavor := range flavors {
		flavorVOs[index] = Flavor{
			Id:     flavor.Id,
			Name:   flavor.Name,
			DishId: flavor.DishId,
			Value:  flavor.Value,
		}
	}

	curName, err := svc.svc.FindNameById(ctx, dishEntity.CategoryId)
	if err != nil {
		return DishVO{}, err
	}

	vo := DishVO{
		CategoryId:   dishEntity.CategoryId,
		CategoryName: curName,
		Description:  dishEntity.Description,
		Flavors:      flavorVOs,
		ID:           dishEntity.Id,
		Image:        dishEntity.Image,
		Name:         dishEntity.Name,
		Price:        dishEntity.Price,
		Status:       dishEntity.Status,
		UpdateTime:   dishEntity.UpdateTime.Format("2006-01-02 15:04"),
	}
	return vo, nil
}

func (svc *DishService) UpdateStatus(ctx context.Context,
	id int64, status int, curId int64) error {

	entity := &domain.Dish{
		Id:         id,
		Status:     status,
		UpdateTime: time.Now(),
		UpdateUser: curId,
	}
	err := svc.repo.UpdateStatus(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

func (svc *DishService) Delete(ctx context.Context, idArr []int64) error {

	err := svc.repo.Delete(ctx, idArr)
	if err != nil {
		return err
	}

	return nil
}

func (svc *DishService) FindByCategoryIdFlavor(ctx context.Context,
	categoryId int64) ([]DishVO, error) {

	dishes, mapping, err := svc.repo.FindByCategoryIdFlavor(ctx, categoryId)
	if err != nil {
		return []DishVO{}, err
	}

	if len(dishes) == 0 {
		return []DishVO{}, nil
	}

	records := make([]DishVO, len(dishes))

	mappingVO := make(map[int64][]Flavor)

	err = copier.Copy(&mappingVO, &mapping)
	if err != nil {
		return []DishVO{}, err
	}

	curName, err := svc.svc.FindNameById(ctx, dishes[0].CategoryId)
	if err != nil {
		return []DishVO{}, err
	}

	for index, d := range dishes {
		records[index] = DishVO{
			CategoryId:   d.CategoryId,
			CategoryName: curName,
			Description:  d.Description,
			Flavors:      mappingVO[d.Id],
			ID:           d.Id,
			Image:        d.Image,
			Name:         d.Name,
			Price:        d.Price,
			Status:       d.Status,
			UpdateTime:   d.UpdateTime.Format("2006-01-02 15:04"),
		}
	}

	return records, nil
}
