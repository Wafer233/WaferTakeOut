package application

import (
	"context"
	"strconv"
	"time"

	setmeal "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/internal/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/pkg/ai"
	"github.com/jinzhu/copier"
)

type SetMealService struct {
	repo        setmeal.SetMealRepository
	categorySvc *rpc.CategoryService
	dishSvc     *rpc.DishService
}

func NewSetMealService(
	repo setmeal.SetMealRepository,
	categorySvc *rpc.CategoryService,
	dishSvc *rpc.DishService,
) *SetMealService {
	return &SetMealService{
		repo:        repo,
		categorySvc: categorySvc,
		dishSvc:     dishSvc,
	}
}

func (svc *SetMealService) Create(ctx context.Context, dto *SetMealDTO, curId int64) error {

	setEntity := &setmeal.SetMeal{
		Id:          dto.Id, // 现在还不知道
		CategoryId:  dto.CategoryId,
		Name:        dto.Name,
		Price:       dto.Price,
		Status:      dto.Status,
		Description: dto.Description,
		Image:       dto.Image,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		CreateUser:  curId,
		UpdateUser:  curId,
	}

	if dto.Description != "" {
		newDescription, _ := ai.GetDescriptionRanking(dto.Description)
		setEntity.Description = newDescription
	}

	dishEntities := make([]*setmeal.SetMealDish, len(dto.SetMealDishes))

	for index, value := range dto.SetMealDishes {
		dishEntities[index] = &setmeal.SetMealDish{
			Id:        value.ID,     // 主键
			SetMealId: setEntity.Id, // 现在还不知道
			DishId:    value.DishId,
			Name:      value.Name,
			Price:     value.Price,
			Copies:    value.Copies,
		}
	}

	err := svc.repo.Create(ctx, setEntity, dishEntities)

	return err
}

func (svc *SetMealService) Deletes(ctx context.Context, ids []int64) error {

	err := svc.repo.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}

func (svc *SetMealService) Update(ctx context.Context, dto *SetMealDTO, curId int64) error {

	set := &setmeal.SetMeal{
		Id:          dto.Id,
		CategoryId:  dto.CategoryId,
		Name:        dto.Name,
		Price:       dto.Price,
		Description: dto.Description,
		Image:       dto.Image,
		UpdateTime:  time.Now(),
		UpdateUser:  curId,
	}

	if dto.Description != "" {
		newDescription, _ := ai.GetDescriptionRanking(dto.Description)
		set.Description = newDescription
	}

	dishes := make([]*setmeal.SetMealDish, len(dto.SetMealDishes))
	for index, value := range dto.SetMealDishes {
		dishes[index] = &setmeal.SetMealDish{
			Id:        value.ID,
			SetMealId: set.Id,
			DishId:    value.DishId,
			Name:      value.Name,
			Price:     value.Price,
			Copies:    value.Copies,
		}
	}

	err := svc.repo.Update(ctx, set, dishes)

	return err
}

func (svc *SetMealService) UpdateStatus(ctx context.Context, id int64,
	curId int64, status int) error {

	entity := &setmeal.SetMeal{
		Id:         id,
		Status:     status,
		UpdateUser: curId,
		UpdateTime: time.Now(),
	}
	err := svc.repo.UpdateStatus(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (svc *SetMealService) FindById(ctx context.Context, id int64) (SetMealVO, error) {

	set, dishes, err := svc.repo.FindById(ctx, id)
	if err != nil {
		return SetMealVO{}, err
	}

	dishVOs := make([]SetMealDish, len(dishes))

	for i, d := range dishes {
		dishVOs[i] = SetMealDish{
			Copies:    d.Copies,
			DishId:    d.DishId,
			ID:        d.Id,
			Name:      d.Name,
			Price:     d.Price,
			SetmealId: set.Id,
		}
	}

	curName, err := svc.categorySvc.FindNameById(ctx, set.CategoryId)
	if err != nil {
		return SetMealVO{}, err
	}

	vo := SetMealVO{
		CategoryId:    set.CategoryId,
		CategoryName:  curName,
		Description:   set.Description,
		Id:            set.Id,
		Image:         set.Image,
		Name:          set.Name,
		Price:         set.Price,
		SetMealDishes: dishVOs,
		Status:        set.Status,
		UpdateTime:    set.UpdateTime.Format("2006-01-02 15:04"),
	}

	return vo, err

}

func (svc *SetMealService) FindPage(ctx context.Context, dto *PageDTO) (PageVO, error) {

	categoryId := dto.CategoryId
	name := dto.Name
	page := dto.Page
	pageSize := dto.PageSize
	status, _ := strconv.Atoi(dto.Status)
	if dto.Status == "" {
		status = -1
	}
	records, total, err := svc.repo.FindPage(ctx, categoryId, name, page, pageSize, status)
	if err != nil || len(records) == 0 {
		return PageVO{}, err
	}
	recordVOs := make([]Record, len(records))
	_ = copier.Copy(&recordVOs, &records)

	for index, _ := range records {
		curName, er := svc.categorySvc.FindNameById(ctx, records[index].CategoryId)
		if er != nil {
			return PageVO{}, er
		}
		recordVOs[index].Status = strconv.Itoa(records[index].Status)
		recordVOs[index].UpdateTime = records[index].UpdateTime.Format("2006-01-02 15:04")
		recordVOs[index].CategoryName = curName
	}

	vo := PageVO{
		Total:   total,
		Records: recordVOs,
	}

	return vo, err
}

func (svc *SetMealService) FindByCategoryId(ctx context.Context, cid int64) ([]FindByCategoryVO, error) {

	setmealEntity, err := svc.repo.FindByCategoryId(ctx, cid)
	if err != nil {
		return []FindByCategoryVO{}, err
	}

	vo := make([]FindByCategoryVO, len(setmealEntity))
	for index, value := range setmealEntity {
		vo[index] = FindByCategoryVO{
			CategoryId:  value.CategoryId,
			CreateTime:  value.CreateTime.Format("2006-01-02 15:04"),
			CreateUser:  value.CreateUser,
			Description: value.Description,
			ID:          value.Id,
			Image:       value.Image,
			Name:        value.Name,
			Price:       value.Price,
			Status:      value.Status,
			UpdateTime:  value.UpdateTime.Format("2006-01-02 15:04"),
			UpdateUser:  value.UpdateUser,
		}
	}
	return vo, nil

}

func (svc *SetMealService) FindDishById(ctx context.Context, setId int64) ([]DishVO, error) {

	dishes, err := svc.repo.FindDishById(ctx, setId)
	if err != nil {
		return []DishVO{}, err
	}

	dishVOs := make([]DishVO, len(dishes))
	for index, d := range dishes {
		des, img, er := svc.dishSvc.FindDescriptionById(ctx, d.DishId)
		if er != nil {
			return []DishVO{}, err
		}
		dishVOs[index] = DishVO{
			Copies: d.Copies,
			// 这两个都是dishid
			Description: des,
			Image:       img,
			Name:        d.Name,
		}
	}
	return dishVOs, err
}
