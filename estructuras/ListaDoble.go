package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Imagen struct {
	Nombre string
	Capas  string
}

type ListaDoble struct {
	Inicio   *NodoDoble
	Longitud int
}

func (l *ListaDoble) estaVacia() bool {
	return l.Longitud == 0
}

func (l *ListaDoble) Insertar(nombre string, capas string) {
	nuevaImagen := &Imagen{Nombre: nombre, Capas: capas}
	// 0x8494AB8 = {20200000, Jose}
	//nuevoAlumno := &Alumno{carnet, nombre}
	if l.estaVacia() {
		l.Inicio = &NodoDoble{nuevaImagen, nil, nil}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = &NodoDoble{nuevaImagen, nil, aux}
		l.Longitud++
	}
}

func (l *ListaDoble) Mostrar() {
	aux := l.Inicio

	for aux != nil {
		fmt.Println(aux.Imagen.Nombre, aux.Imagen.Capas)
		aux = aux.Siguiente
	}
}

func LeerImagen(ruta string, listaAux *ListaDoble) {
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

func (l *ListaDoble) ReporteImagen() {
	nombreArchivo := "./reportes/listadoble.dot"
	nombreImagen := "./reportes/listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Inicio
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.Imagen.Nombre + "\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
