package estructuras

// Nodo is just a node
type NodoSimple struct {
	Empleado  *Empleado
	Siguiente *NodoSimple
}

type NodoDoble struct {
	Imagen    *Imagen
	Siguiente *NodoDoble
	Anterior  *NodoDoble
}

type NodoCircular struct {
	Cliente   *Cliente
	Siguiente *NodoCircular
}

type NodoCola struct {
	Cliente   *Cliente
	Siguiente *NodoCola
}

type NodoPila struct {
	Pedido    *Pedido
	Siguiente *NodoPila
}

type NodoMatriz struct {
	Siguiente *NodoMatriz
	Anterior  *NodoMatriz
	Abajo     *NodoMatriz
	Arriba    *NodoMatriz
	PosX      int
	PosY      int
	Color     string
}
