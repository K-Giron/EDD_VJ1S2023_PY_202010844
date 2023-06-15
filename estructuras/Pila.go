package estructuras

import "fmt"

type Pila struct {
	Primero  *NodoPila
	Longitud int
}

type Pedido struct {
	nombreImagen string
	idCliente    string
}

func (p *Pila) Push(nombreImagen string, idCliente string) {
	nuevaPila := &Pedido{nombreImagen: nombreImagen, idCliente: idCliente}
	if p.Longitud == 0 {
		nuevoNodo := &NodoPila{nuevaPila, nil}
		p.Primero = nuevoNodo
		p.Longitud++
	} else {
		nuevoNodo := &NodoPila{nuevaPila, p.Primero}
		p.Primero = nuevoNodo
		p.Longitud++
	}
}

func (p *Pila) Pop() {
	if p.Longitud == 0 {
		fmt.Println("No hay elementos en la pila")
	} else {
		p.Primero = p.Primero.Siguiente
		p.Longitud--
	}
}

func (p *Pila) Graficar() {
	nombre_archivo := "./reportes/pila.dot"
	nombre_imagen := "./reportes/pila.jpg"
	texto := "digraph pila{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record]"
	aux := p.Primero
	texto += "nodo0 [label=\""
	for i := 0; i < p.Longitud; i++ {
		texto = texto + "|(Imagen: " + aux.Pedido.nombreImagen + ", ID cliente: " + aux.Pedido.idCliente + ")"
		aux = aux.Siguiente
	}
	texto += "\"]; \n}"
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}
