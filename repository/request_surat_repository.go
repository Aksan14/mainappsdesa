package repository

import (
	"godesaapps/model"
)

type RequestSuratRepository interface {
	FindByNik(nik string) (*model.DataWarga, error)
	FindDataWargaByNIK(nik string) (*model.DataWarga, error)
	InsertRequestSurat(request model.RequestSuratWarga) error
}
