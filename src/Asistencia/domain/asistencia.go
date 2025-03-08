package domain

import "time"

type Asistencia struct {
	ID          int `json:"id"`
	ChoferID    int `json:"chofer_id"`
	BusID       int `json:"bus_id"`
	FechaInicio time.Time `json:"fecha_inicio"`
	// FechaFin    time.Time `json:"fecha_fin"`
	Estado      string `json:"estado"`
}

func NewAsistencia (id int, choferID int, busID int, fechaInicio time.Time, fechaFin time.Time, estado string) *Asistencia {
	return &Asistencia{
		ID: id,
		ChoferID: choferID,
		BusID: busID,
		FechaInicio: fechaInicio,
		// FechaFin: fechaFin,
		Estado: estado,
	}
}