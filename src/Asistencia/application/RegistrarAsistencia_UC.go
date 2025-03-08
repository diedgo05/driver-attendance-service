package application

import "driver-attendance/src/Asistencia/domain"

type RegistrarAsistenciaUC struct {
	repo domain.IAsistenciaRepository
}

func NewRegistrarAsistenciaUC (repo domain.IAsistenciaRepository) *RegistrarAsistenciaUC {
	return &RegistrarAsistenciaUC{repo: repo}
}

func (uc *RegistrarAsistenciaUC) Run(asistencia domain.Asistencia) error {
	return uc.repo.RegistrarAsistencia(asistencia)
}