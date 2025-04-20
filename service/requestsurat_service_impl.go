package service

import (
	"fmt"
	"godesaapps/dto"
	"godesaapps/model"
	"godesaapps/repository"
	"time"
)

type requestSuratServiceImpl struct {
	repo repository.RequestSuratRepository
}

func NewRequestSuratService(repo repository.RequestSuratRepository) RequestSuratService {
	return &requestSuratServiceImpl{repo}
}

func (s *requestSuratServiceImpl) FindByNik(nik string) (*model.DataWarga, error) {
	warga, err := s.repo.FindByNik(nik)
	if err != nil {
		return nil, fmt.Errorf("gagal mencari warga dengan NIK %s: %v", nik, err)
	}

	if warga == nil {
		return nil, fmt.Errorf("warga dengan NIK %s tidak ditemukan", nik)
	}

	return warga, nil
}

func (s *requestSuratServiceImpl) RequestSurat(input dto.RequestSuratDTO) error {
	warga, err := s.repo.FindDataWargaByNIK(input.NIK)
	if err != nil {
		return fmt.Errorf("gagal mencari data warga dengan NIK %s: %v", input.NIK, err)
	}

	if warga == nil {
		return fmt.Errorf("data warga dengan NIK %s tidak ditemukan", input.NIK)
	}

	request := model.RequestSuratWarga{
		IDWarga:          warga.ID,
		JenisSurat:       input.JenisSurat,
		Status:           "Pending",
		NomorSurat:       "",
		Keterangan:       input.Keterangan,
		TanggalPengajuan: time.Now(),
		TanggalSelesai:   nil,
		NIK:              warga.NIK,
		NamaLengkap:      warga.NamaLengkap,
		TempatLahir:      warga.TempatLahir,
		TanggalLahir:     warga.TanggalLahir,
		JenisKelamin:     warga.JenisKelamin,
		Pendidikan:       warga.Pendidikan,
		Pekerjaan:        warga.Pekerjaan,
		Agama:            warga.Agama,
		StatusPernikahan: warga.StatusPernikahan,
		Kewarganegaraan:  warga.Kewarganegaraan,
		Alamat:           warga.Alamat,
		Penghasilan:      input.Penghasilan,
		LamaTinggal:      input.LamaTinggal,
		NamaUsaha:        input.NamaUsaha,
		JenisUsaha:       input.JenisUsaha,
		AlamatUsaha:      input.AlamatUsaha,
		AlamatTujuan:     input.AlamatTujuan,
		AlasanPindah:     input.AlasanPindah,
		KeperluanPindah:  input.KeperluanPindah,
		TujuanPindah:     input.TujuanPindah,
	}

	err = s.repo.InsertRequestSurat(request)
	if err != nil {
		return fmt.Errorf("gagal menyimpan permintaan surat untuk warga NIK %s: %v", input.NIK, err)
	}

	return nil
}