package employeeApp

import "context"

func (svc *EmployeeService) StatusFlips(ctx context.Context,
	status int, id int64) error {
	err := svc.repo.UpdateStatusByID(ctx, status, id)
	return err
}
