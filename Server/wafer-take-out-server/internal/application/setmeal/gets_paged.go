package setmealApp

import (
	"context"
	"strconv"
)

type PageDTO struct {
	CategoryId int64  `form:"categoryId"`
	Name       string `form:"name"`
	Page       int    `form:"page"`
	PageSize   int    `form:"pageSize"`
	Status     string `form:"status"`
}
type PageVO struct {
	Total   int64    `json:"total"`
	Records []Record `json:"records"`
}

type Record struct {
	Id           int64   `json:"id"`
	CategoryId   int64   `json:"categoryId"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	Description  string  `json:"description"`
	Image        string  `json:"image"`
	UpdateTime   string  `json:"updateTime"`
	CategoryName string  `json:"categoryName"`
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
		cat, er := svc.cateRepo.GetById(ctx, value.CategoryId)
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
