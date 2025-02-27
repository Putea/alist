package db

import (
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/pkg/errors"
)

// why don't need `cache` for storage?
// because all storage store in `operations.storagesMap`
// the most of the read operation is from `operations.storagesMap`
// just for persistence in database

// CreateStorage just insert storage to database
func CreateStorage(storage *model.Storage) error {
	return errors.WithStack(db.Create(storage).Error)
}

// UpdateStorage just update storage in database
func UpdateStorage(storage *model.Storage) error {
	return errors.WithStack(db.Save(storage).Error)
}

// DeleteStorageById just delete storage from database by id
func DeleteStorageById(id uint) error {
	return errors.WithStack(db.Delete(&model.Storage{}, id).Error)
}

// GetStorages Get all storages from database order by index
func GetStorages(pageIndex, pageSize int) ([]model.Storage, int64, error) {
	storageDB := db.Model(&model.Storage{})
	var count int64
	if err := storageDB.Count(&count).Error; err != nil {
		return nil, 0, errors.Wrapf(err, "failed get storages count")
	}
	var storages []model.Storage
	if err := storageDB.Order(columnName("index")).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&storages).Error; err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return storages, count, nil
}

// GetStorageById Get Storage by id, used to update storage usually
func GetStorageById(id uint) (*model.Storage, error) {
	var storage model.Storage
	storage.ID = id
	if err := db.First(&storage).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &storage, nil
}
