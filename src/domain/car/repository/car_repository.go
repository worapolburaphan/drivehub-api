package repository

import (
	"errors"
	"github.com/drivehub-api/src/entity"
	"github.com/drivehub-api/src/external"
	"github.com/mitchellh/mapstructure"
	"log"
)

type repository struct {
	db external.Storage
}

type CarRepository interface {
	FindAll() ([]*entity.Car, error)
	Insert(*entity.Car) (*entity.Car, error)
	Update(entity *entity.Car) (*entity.Car, error)
	Delete(*entity.Car) error
}

func NewCarRepository(db external.Storage) CarRepository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]*entity.Car, error) {
	all, err := r.db.FindAll()
	if err != nil {
		return nil, err
	}

	cars := make([]*entity.Car, 0)

	for _, record := range all {
		car := entity.Car{}
		err := mapstructure.Decode(record, &car)
		if err != nil {
			log.Printf("error decoding record: %v", err)
			continue
		}

		cars = append(cars, &car)
	}

	return cars, nil
}

func (r *repository) Insert(car *entity.Car) (*entity.Car, error) {
	res, err := r.db.Insert(car)
	if err != nil {
		return nil, err
	}

	retCar, ok := res.(*entity.Car)
	if !ok {
		return nil, errors.New("could not insert car")
	}

	return retCar, nil
}

func (r *repository) Update(car *entity.Car) (*entity.Car, error) {
	update, err := r.db.Update(car)
	if err != nil {
		return nil, err
	}

	retCar, ok := update.(*entity.Car)
	if !ok {
		return nil, errors.New("could not update car")
	}

	return retCar, nil
}

func (r *repository) Delete(car *entity.Car) error {
	return r.db.Delete(car)
}
