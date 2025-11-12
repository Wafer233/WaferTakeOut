package categoryApp

import (
	"context"
)

type PageDTO struct {
	Name     string `form:"name"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Type     int    `form:"type"`
}

type PageVO struct {
	Total   int64    `json:"total"`
	Records []Record `json:"records"`
}
type Record struct {
	ID         int64  `json:"id"`
	Type       int    `json:"type"`
	Name       string `json:"name"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	CreateUser int64  `json:"createUser"`
	UpdateUser int64  `json:"updateUser"`
}

func (svc *CategoryService) PageQuery(ctx context.Context, dto *PageDTO) (PageVO, error) {

	curName := dto.Name
	page := dto.Page
	pageSize := dto.PageSize
	curType := dto.Type

	entities, total, err := svc.repo.GetsByPaged(ctx, curName, curType, page, pageSize)
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
