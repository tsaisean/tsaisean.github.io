package repositories

import (
	"gorm.io/gorm"
	"iHR/db"
	. "iHR/db/models"
)

//go:generate mockery --all --output=./mocks
type EmployeeRepository interface {
	CreateEmployee(employee *Employee) (*Employee, error)
	GetAllEmployees() ([]Employee, error)
	GetEmployeeByID(id uint) (*Employee, error)
	UpdateEmployeeByID(id uint, updated *Employee) (*Employee, error)
	DeleteEmployee(id uint) error
}

type EmployeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) *EmployeeRepo {
	return &EmployeeRepo{db: db}
}

func (r *EmployeeRepo) CreateEmployee(employee *Employee) (*Employee, error) {
	if err := r.db.Create(employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *EmployeeRepo) GetAllEmployees() ([]Employee, error) {
	var employees []Employee
	if err := r.db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepo) GetEmployeeByID(id uint) (*Employee, error) {
	employee := new(Employee)
	if err := r.db.First(employee, id).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *EmployeeRepo) UpdateEmployeeByID(id uint, updated *Employee) (*Employee, error) {
	employee, err := r.GetEmployeeByID(id)
	if err != nil {
		return nil, err
	}

	if err := db.DB.Model(&employee).Updates(updated).Error; err != nil {
		return employee, err
	}
	return employee, nil
}

func (r *EmployeeRepo) DeleteEmployee(id uint) error {
	_, err := r.GetEmployeeByID(id)
	if err != nil {
		return err
	}
	return db.DB.Delete(&Employee{}, id).Error
}
