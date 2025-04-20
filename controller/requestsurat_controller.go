package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"godesaapps/dto"
	"godesaapps/service"
	"github.com/julienschmidt/httprouter"
)

type RequestSuratController struct {
	service service.RequestSuratService
}

func NewRequestSuratController(service service.RequestSuratService) *RequestSuratController {
	return &RequestSuratController{service}
}

func (c *RequestSuratController) FindWargaByNik(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	nik := ps.ByName("nik")
	if nik == "" {
		http.Error(w, "NIK tidak boleh kosong", http.StatusBadRequest)
		return
	}

	warga, err := c.service.FindByNik(nik)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal mencari data warga: %v", err), http.StatusInternalServerError)
		return
	}

	if warga == nil {
		http.Error(w, "Data warga tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(warga)
}

func (c *RequestSuratController) CreateRequestSurat(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var input dto.RequestSuratDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := c.service.RequestSurat(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Permintaan surat berhasil dikirim.",
	})
}

