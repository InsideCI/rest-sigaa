package repository

import "github.com/InsideCI/nego/model"

// DepartmentRepository abstracts CRUD methods for Department.
type DepartmentRepository interface {
	Create(deps []model.Department) (int, error)
	Fetch(limit int) ([]*model.Department, error)
}
