package employeeApp

import "context"

type StatusFlipsDTO struct {
	ID int64 `form:"id"`
}

func (svc *EmployeeService) StatusFlips(ctx context.Context,
	status int, id int64) error {
	err := svc.repo.UpdateStatusByID(ctx, status, id)
	return err
}
