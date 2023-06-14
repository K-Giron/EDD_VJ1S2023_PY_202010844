package main

import (
	"PF1_202010844/estructuras"
	"fmt"
)

func main() {
	fmt.Println("Hello, World! Esto es una prueba de go")
	ListaSimple := &estructuras.ListaSimple{Inicio: nil, Longitud: 0}
	ListaSimple.Insertar("202010844", "Jorge", "Estudiante", "1234")
	ListaSimple.Insertar("202010845", "Luis", "Trabajador", "12345")
	ListaSimple.Insertar("202010846", "Pedro", "Limpiador", "123456")
	ListaSimple.Insertar("202010847", "Guicho", "Gerente", "1234567")
	ListaSimple.Insertar("202010848", "Kevin", "Estudiante", "12345678")

	estructuras.LeerEmpleados("archivos_prueba/empleados.csv", ListaSimple)
	ListaSimple.Mostrar()
	ListaSimple.ReporteEmpleados()

	fmt.Println("-----------------------------------------")
	ListaDoble := &estructuras.ListaDoble{Inicio: nil, Longitud: 0}
	ListaDoble.Insertar("imagen1", "1")
	ListaDoble.Insertar("imagen2", "2")
	ListaDoble.Insertar("imagen3", "3")
	ListaDoble.Insertar("imagen4", "4")
	estructuras.LeerImagen("archivos_prueba/imagenes.csv", ListaDoble)
	ListaDoble.Mostrar()
	ListaDoble.ReporteImagen()

	fmt.Println("-----------------------------------------")
	ListaCircular := &estructuras.ListaCircular{Inicio: nil, Longitud: 0}
	ListaCircular.Insertar("202010844", "Matias")
	ListaCircular.Insertar("202010845", "Ana")
	ListaCircular.Insertar("202010846", "Susana")
	ListaCircular.Insertar("202010847", "jeremias")
	estructuras.LeerClientes("archivos_prueba/clientes_registrados.csv", ListaCircular)
	ListaCircular.Mostrar()
	ListaCircular.ReporteClientes()

	fmt.Println("-----------------------------------------")
	Cola := &estructuras.Cola{Primero: nil, Longitud: 0}
	Cola.Encolar("4343", "Cristian")
	Cola.Encolar("4344", "Juan")
	Cola.Encolar("4345", "Pedro")
	Cola.Encolar("4346", "Maria")
	estructuras.LeerCola("archivos_prueba/clientes_cola.csv", Cola)
	Cola.Graficar()

}
