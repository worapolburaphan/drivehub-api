package usecase

import (
	"github.com/drivehub-api/src/domain/car/repository"
	"github.com/drivehub-api/src/entity"
)

type usecase struct {
	carRepo repository.CarRepository
}

type CarUseCase interface {
	GetCars() ([]*entity.Car, error)
	CreateCar(*entity.Car) (*entity.Car, error)
	UpdateCar(*entity.Car) (*entity.Car, error)
	DeleteCar(id string) error
}

func NewCarUsecase(carRepo repository.CarRepository) CarUseCase {
	return &usecase{carRepo: carRepo}
}

func (u *usecase) GetCars() ([]*entity.Car, error) {
	return u.carRepo.FindAll()
}

func (u *usecase) CreateCar(car *entity.Car) (*entity.Car, error) {
	return u.carRepo.Insert(car)
}

func (u *usecase) UpdateCar(car *entity.Car) (*entity.Car, error) {
	return u.carRepo.Update(car)
}

func (u *usecase) DeleteCar(id string) error {
	car := entity.NewCar(id)
	return u.carRepo.Delete(car)
}
