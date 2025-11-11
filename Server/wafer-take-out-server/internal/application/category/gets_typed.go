package categoryApp

import "context"

type GetsTypedVO []Record

func (svc *CategoryService) TypeQuery(ctx context.Context, curType int) (GetsTypedVO, error) {

	entities, err := svc.repo.GetsByType(ctx, curType)
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
		record.CreateTime = entities[index].CreateTime.Format("2006-01-02 15:04:05")
		record.UpdateTime = entities[index].UpdateTime.Format("2006-01-02 15:04:05")
		record.CreateUser = entities[index].CreateUser
		record.UpdateUser = entities[index].UpdateUser
		vo[index] = record
	}

	return vo, nil

}
