package dishApp

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
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	CategoryId   int64   `json:"categoryId"`
	Price        float64 `json:"price"`
	Image        string  `json:"image"`
	Description  string  `json:"description"`
	Status       int     `json:"status"`
	UpdateTime   string  `json:"updateTime"`
	CategoryName string  `json:"categoryName"`
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

	categoryName := ""
	category, err := svc.cateRepo.GetById(ctx, categoryId)
	if err != nil || category == nil {
		categoryName = ""
	} else {
		categoryName = category.Name
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
		record.UpdateTime = dishes[index].UpdateTime.Format("2006-01-02 15:04:05")
		records[index].CategoryName = categoryName
		records[index] = record
	}

	if dto.CategoryId == 0 {
		for index, record := range records {
			category, err = svc.cateRepo.GetById(ctx, record.CategoryId)
			if err != nil || category == nil {
				categoryName = ""
			} else {
				categoryName = category.Name
			}
			record.CategoryName = categoryName
			records[index] = record
		}
	}

	vo := PageVO{
		Total:   total,
		Records: records,
	}
	return vo, nil
}
