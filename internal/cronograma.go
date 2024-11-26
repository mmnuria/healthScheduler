package health_scheduler

type Asignacion struct {
	Empleado   Empleado
	Area       string
	Turno      Turno
	EsFlexible bool
}

type CronogramaDiario struct {
	Fecha        string
	Asignaciones []Asignacion
}

func NuevaAsignacion(empleado Empleado, area string, turno Turno, esFlexible bool) *Asignacion {
	return &Asignacion{
		Empleado:   empleado,
		Area:       area,
		Turno:      turno,
		EsFlexible: esFlexible,
	}
}
