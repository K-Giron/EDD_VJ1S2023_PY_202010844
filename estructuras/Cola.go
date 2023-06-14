package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Cola struct {
	Primero  *NodoCola
	Longitud int
}

type ListaNuevoCliente struct {
	Inicio   *NodoCircular
	Longitud int
}

func (c *Cola) Encolar(idCliente string, nombre string) {

	if idCliente == "X" {
		idCliente = generarIDAleatorio()

	}

	nuevoCliente := &Cliente{IdCliente: idCliente, Nombre: nombre}
	if c.Longitud == 0 {
		nuevoNodo := &NodoCola{nuevoCliente, nil}
		c.Primero = nuevoNodo
		c.Longitud++
	} else {
		nuevoNodo := &NodoCola{nuevoCliente, nil}
		aux := c.Primero
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func generarIDAleatorio() string {
	rand.Seed(time.Now().UnixNano())
	randomID := rand.Intn(100000)
	return strconv.Itoa(randomID)
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		fmt.Println("No hay alumnos pendientes en la cola")
	} else {
		c.Primero = c.Primero.Siguiente
		c.Longitud--
	}
}

func LeerCola(ruta string, listaAux *Cola) {
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

		listaAux.Encolar(linea[0], linea[1])
	}
}

func (c *Cola) Graficar() {
	nombre_archivo := "./reportes/cola.dot"
	nombre_imagen := "./reportes/cola.jpg"
	texto := "digraph cola{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := c.Primero
	contador := 0
	for i := 0; i < c.Longitud; i++ {
		texto = texto + "nodo" + strconv.Itoa(i) + "[label=\"{ID: " + aux.Cliente.IdCliente + ", Nombre: " + aux.Cliente.Nombre + "|}\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < c.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}
