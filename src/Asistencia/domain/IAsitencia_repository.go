package domain

type IAsistenciaRepository interface {
	RegistrarAsistencia(asistencia Asistencia) error
} 