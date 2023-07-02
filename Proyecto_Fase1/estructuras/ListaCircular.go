package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Cliente struct {
	IdCliente string
	Nombre    string
}

type ListaCircular struct {
	Inicio   *NodoCircular
	Longitud int
}

func (l *ListaCircular) Insertar(idCliente string, nombre string) {
	nuevoCliente := &Cliente{IdCliente: idCliente, Nombre: nombre}
	if l.Longitud == 0 {
		l.Inicio = &NodoCircular{Cliente: nuevoCliente, Siguiente: nil}
		l.Inicio.Siguiente = l.Inicio
		l.Longitud++
	} else {
		if l.Longitud == 1 {
			l.Inicio.Siguiente = &NodoCircular{Cliente: nuevoCliente, Siguiente: l.Inicio}
			l.Longitud++
		} else {
			aux := l.Inicio
			for i := 0; i < l.Longitud-1; i++ {
				aux = aux.Siguiente
			}
			aux.Siguiente = &NodoCircular{Cliente: nuevoCliente, Siguiente: l.Inicio}
			l.Longitud++
		}
	}
}

func (l *ListaCircular) Mostrar() {
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		fmt.Println(aux.Cliente.IdCliente, aux.Cliente.Nombre)
		aux = aux.Siguiente
	}
}

func LeerClientes(ruta string, listaAux *ListaCircular) {
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

		listaAux.Insertar(linea[0], linea[1])
	}
}

func (l *ListaCircular) ReporteClientes() {
	nombreArchivo := "./reportes/listacircular.dot"
	nombreImagen := "./reportes/listacircular.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.Cliente.IdCliente + ",\n" + aux.Cliente.Nombre + "\", shape=box];\n"
		aux = aux.Siguiente
	}
	// Conectar los nodos en un ciclo circular
	for i := 0; i < l.Longitud-1; i++ {
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(i+1) + ";\n"
	}
	// Conectar el último nodo al primer nodo
	texto += "nodo" + strconv.Itoa(l.Longitud-1) + "->nodo0;\n"

	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
