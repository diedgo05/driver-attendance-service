package dependencies

import (
	"driver-attendance/src/Asistencia/application"
	"driver-attendance/src/Asistencia/infraestructure"
	"driver-attendance/src/Asistencia/infraestructure/http/controllers"
	"driver-attendance/src/core"
	"fmt"
)

var (
	mySQL infraestructure.MySQL
)

func Init() {
	db, err := core.InitMySQL()
	if err != nil {
		fmt.Println(err)
	}
	mySQL = *infraestructure.NewMySQL(db)
}

func RegistrarAsistenciaCtrl() *controllers.RegistrarAsistenciaCtrl {
	ucRegistrarAsistencia := application.NewRegistrarAsistenciaUC(&mySQL)
	return controllers.NewRegistrarAsistenciaCtrl(ucRegistrarAsistencia)
}
