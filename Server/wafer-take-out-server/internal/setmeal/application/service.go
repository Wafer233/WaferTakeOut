package application

import (
	"context"
	"strconv"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/setmeal_dish"
	setmeal "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/domain"
)

type SetMealService struct {
	setRepo  setmeal.SetMealRepository
	dishRepo setmeal_dish.SetMealDishRepository
	cateRepo domain.CategoryRepository
}

func NewSetMealService(setRepo setmeal.SetMealRepository,
	dishRepo setmeal_dish.SetMealDishRepository,
	cateRepo domain.CategoryRepository) *SetMealService {
	return &SetMealService{
		setRepo:  setRepo,
		dishRepo: dishRepo,
		cateRepo: cateRepo,
	}
}

func (svc *SetMealService) AddSetMeal(ctx context.Context, dto *AddSetMealDTO, curId int64) error {

	set := &setmeal.SetMeal{
		Id:          dto.Id,
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

	mealDishes := make([]*setmeal_dish.SetMealDish, len(dto.SetMealDishes))

	err := svc.setRepo.Insert(ctx, set)
	if err != nil {
		return err
	}

	for index, value := range dto.SetMealDishes {
		mealDishes[index] = &setmeal_dish.SetMealDish{
			Id:        value.ID,
			SetMealId: set.Id,
			DishId:    value.DishId,
			Name:      value.Name,
			Price:     value.Price,
			Copies:    value.Copies,
		}
	}

	err = svc.dishRepo.Inserts(ctx, mealDishes)
	if err != nil {
		return err
	}
	return nil
}

func (svc *SetMealService) Deletes(ctx context.Context, ids []int64) error {

	err := svc.setRepo.DeletesByIds(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}

func (svc *SetMealService) Edit(ctx context.Context, dto *AddSetMealDTO, curId int64) error {

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
	err := svc.setRepo.UpdateById(ctx, set)
	if err != nil {
		return err
	}

	if len(dto.SetMealDishes) == 0 {
		return nil
	}

	sets := make([]*setmeal_dish.SetMealDish, len(dto.SetMealDishes))
	for i, dish := range dto.SetMealDishes {
		sets[i] = &setmeal_dish.SetMealDish{
			Id:        dish.ID,
			SetMealId: set.Id,
			DishId:    dish.DishId,
			Name:      dish.Name,
			Price:     dish.Price,
			Copies:    dish.Copies,
		}
	}
	err = svc.dishRepo.UpdatesBySetMealId(ctx, sets)
	if err != nil {
		return err
	}
	return nil
}

func (svc *SetMealService) StatusFlip(ctx context.Context, id int64, curId int64, status int) error {

	entity := &setmeal.SetMeal{
		Id:         id,
		Status:     status,
		UpdateUser: curId,
		UpdateTime: time.Now(),
	}
	err := svc.setRepo.UpdateStatusById(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (svc *SetMealService) IdQuery(ctx context.Context, id int64) (GetSetMealVO, error) {

	set, err := svc.setRepo.GetById(ctx, id)
	if err != nil {
		return GetSetMealVO{}, err
	}
	dishes, err := svc.dishRepo.GetsBySetMealId(ctx, set.Id)
	if err != nil {
		return GetSetMealVO{}, err
	}

	dishVOs := make([]SetMealDish, len(dishes))

	for i, dish := range dishes {
		dishVOs[i] = SetMealDish{
			Copies:    dish.Copies,
			DishId:    dish.DishId,
			ID:        dish.Id,
			Name:      dish.Name,
			Price:     dish.Price,
			SetmealId: set.Id,
		}
	}

	category, err := svc.cateRepo.FindById(ctx, set.CategoryId)
	if err != nil {
		return GetSetMealVO{}, err
	}

	vo := GetSetMealVO{
		CategoryId:    set.CategoryId,
		CategoryName:  category.Name,
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

func (svc *SetMealService) PageQuery(ctx context.Context, dto *PageDTO) (PageVO, error) {

	categoryId := dto.CategoryId
	name := dto.Name
	page := dto.Page
	pageSize := dto.PageSize
	status, _ := strconv.Atoi(dto.Status)
	if dto.Status == "" {
		status = -1
	}
	records, total, err := svc.setRepo.GetsPaged(ctx, categoryId, name, page, pageSize, status)
	if err != nil {
		return PageVO{}, err
	}

	catNames := make([]string, len(records))
	for index, value := range records {
		cat, er := svc.cateRepo.FindById(ctx, value.CategoryId)
		if er != nil {
			return PageVO{}, er
		}
		catNames[index] = cat.Name
	}

	recordVOs := make([]Record, len(records))
	for index, value := range records {
		recordVOs[index] = Record{
			Id:           value.Id,
			CategoryId:   value.CategoryId,
			Name:         value.Name,
			Price:        value.Price,
			Status:       strconv.Itoa(value.Status),
			Description:  value.Description,
			Image:        value.Image,
			UpdateTime:   value.UpdateTime.Format("2006-01-02 15:04"),
			CategoryName: catNames[index],
		}
	}

	vo := PageVO{
		Total:   total,
		Records: recordVOs,
	}

	return vo, err
}
