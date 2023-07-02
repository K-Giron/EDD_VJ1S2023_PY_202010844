package Lista

type NodoSimple struct {
	Empleado  *Empleado
	Siguiente *NodoSimple
}

type Empleado struct {
	IdEmpleado string
	Nombre     string
	Cargo      string
	Password   string
}
