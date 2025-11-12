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

func (svc *SetMealService) PageQuery(ctx context.Context,
	dto *PageDTO) (PageVO, error) {

	categoryId := dto.CategoryId
	name := dto.Name
	page := dto.Page
	pageSize := dto.PageSize
	status := dto.Status
	statusInt := 0
	if status == "" {
		statusInt = 2
	} else {
		statusInt, _ = strconv.Atoi(status)
	}

	records, total, err := svc.setRepo.GetsPaged(ctx, categoryId, name, page, pageSize, statusInt)
	recordVOs := make([]Record, len(records))

	catNames := make([]string, len(records))
	for index, value := range records {
		cat, er := svc.cateRepo.GetById(ctx, value.CategoryId)
		if er != nil {
			return PageVO{}, er
		}
		catNames[index] = cat.Name
	}

	for index, vo := range recordVOs {
		vo.Id = records[index].Id
		vo.CategoryId = records[index].CategoryId
		vo.Name = records[index].Name
		vo.Price = records[index].Price
		vo.Status = strconv.Itoa(records[index].Status)
		vo.Description = records[index].Description
		vo.Image = records[index].Image
		vo.UpdateTime = records[index].UpdateTime.Format("2006-01-02 15:04:05")
		vo.CategoryName = catNames[index]
	}

	vo := PageVO{
		Total:   total,
		Records: recordVOs,
	}

	return vo, err
}
