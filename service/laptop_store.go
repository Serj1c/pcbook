package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Serj1c/pcbook/pb"
	"github.com/jinzhu/copier"
)

var (
	// ErrAlreadyExists is returned when a record with a given ID exists already
	ErrAlreadyExists = errors.New("record already exists")
)

// LaptopStore is an interface to store a laptop
type LaptopStore interface {
	// Save saves the laptop to the store
	Save(laptop *pb.Laptop) error
	// Find finds laptop by ID
	Find(id string) *pb.Laptop
	// Search searches for laptops with filter and returns one by one via the found function
	Search(filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}

//InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// DBLaptopStore is a stub for the future
type DBLaptopStore struct{}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// Save saves the laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data %v", err)
	}
	store.data[other.Id] = other
	return nil
}

// Find finds laptop by ID
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, err
	}

	return other, nil
}
