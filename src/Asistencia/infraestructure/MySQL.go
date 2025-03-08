package infraestructure

import (
	"database/sql"
	"driver-attendance/src/Asistencia/domain"
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{DB: db}
}

func (r *MySQL) RegistrarAsistencia(asistencia domain.Asistencia) error {
	query := "INSERT INTO asistencias (chofer_id, bus_id, fecha_inicio, estado) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, asistencia.ChoferID, asistencia.BusID, asistencia.FechaInicio, asistencia.Estado)

	return err

}