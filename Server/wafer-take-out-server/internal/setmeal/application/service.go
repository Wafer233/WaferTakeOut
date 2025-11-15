package application

import (
	"context"
	"strconv"
	"time"

	category "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/domain"
	dish "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/domain"
	setmeal "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/domain"
)

type SetMealService struct {
	setRepo  setmeal.SetMealRepository
	cateRepo category.CategoryRepository
	dishRepo dish.DishRepository
}

func NewSetMealService(
	setRepo setmeal.SetMealRepository,
	cateRepo category.CategoryRepository,
	dishRepo dish.DishRepository) *SetMealService {
	return &SetMealService{
		setRepo:  setRepo,
		cateRepo: cateRepo,
		dishRepo: dishRepo,
	}
}

func (svc *SetMealService) Create(ctx context.Context, dto *AddSetMealDTO, curId int64) error {

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

	err := svc.setRepo.Create(ctx, setEntity, dishEntities)

	return err
}

func (svc *SetMealService) Deletes(ctx context.Context, ids []int64) error {

	err := svc.setRepo.Delete(ctx, ids)
	if err != nil {
		return err
	}
	return nil
}

func (svc *SetMealService) Update(ctx context.Context, dto *AddSetMealDTO, curId int64) error {

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

	err := svc.setRepo.Update(ctx, set, dishes)

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
	err := svc.setRepo.UpdateStatus(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (svc *SetMealService) FindById(ctx context.Context, id int64) (GetSetMealVO, error) {

	set, dishes, err := svc.setRepo.FindById(ctx, id)
	if err != nil {
		return GetSetMealVO{}, err
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

	cate, err := svc.cateRepo.FindById(ctx, set.CategoryId)
	if err != nil {
		return GetSetMealVO{}, err
	}

	vo := GetSetMealVO{
		CategoryId:    set.CategoryId,
		CategoryName:  cate.Name,
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
	records, total, err := svc.setRepo.FindPage(ctx, categoryId, name, page, pageSize, status)
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

func (svc *SetMealService) FindByCategoryId(ctx context.Context, cid int64) ([]FindByCategoryVO, error) {

	setmealEntity, err := svc.setRepo.FindByCategoryId(ctx, cid)
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

	dishes, err := svc.setRepo.FindDishById(ctx, setId)
	if err != nil {
		return []DishVO{}, err
	}

	var ids []int64
	for _, value := range dishes {
		//这个地方是dishid
		ids = append(ids, value.DishId)
	}

	descriptions, images, err := svc.dishRepo.FindByIds(ctx, ids)
	if err != nil || len(descriptions) != len(dishes) {
		return []DishVO{}, err
	}

	dishVOs := make([]DishVO, len(dishes))
	for index, d := range dishes {
		dishVOs[index] = DishVO{
			Copies: d.Copies,
			// 这两个都是dishid
			Description: descriptions[d.DishId],
			Image:       images[d.DishId],
			Name:        d.Name,
		}
	}
	return dishVOs, err
}
