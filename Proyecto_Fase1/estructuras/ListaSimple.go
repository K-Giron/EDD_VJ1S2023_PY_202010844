package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Empleado struct {
	IdEmpleado string
	Nombre     string
	Cargo      string
	Password   string
}

type ListaSimple struct {
	Inicio   *NodoSimple
	Longitud int
}

func (l *ListaSimple) estaVacia() bool {
	return l.Longitud == 0
}

func (l *ListaSimple) Insertar(idEmpleado string, nombre string, cargo string, password string) {
	nuevoEmpleado := &Empleado{IdEmpleado: idEmpleado, Nombre: nombre, Cargo: cargo, Password: password}
	if l.estaVacia() {
		l.Inicio = &NodoSimple{nuevoEmpleado, nil}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = &NodoSimple{nuevoEmpleado, nil}
		l.Longitud++
	}
}

func (l *ListaSimple) Mostrar() {
	aux := l.Inicio

	for aux != nil {
		fmt.Println(aux.Empleado.IdEmpleado, aux.Empleado.Nombre, aux.Empleado.Cargo, aux.Empleado.Password)
		aux = aux.Siguiente
	}
}

func LeerEmpleados(ruta string, listaAux *ListaSimple) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	encabezado := true
	for {
		linea, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude la línea del csv")
			return
		}
		if encabezado {
			encabezado = false
			continue
		}

		listaAux.Insertar(linea[0], linea[1], linea[2], linea[3])
	}
}

func (l *ListaSimple) ReporteEmpleados() {
	nombreArchivo := "./reportes/listasimple.dot"
	nombreImagen := "./reportes/listasimple.png"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + "ID: " + aux.Empleado.IdEmpleado + ",\n" + " Nombre: " + aux.Empleado.Nombre + "\", shape=box];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
	}

	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
