package ArbolAVL

import (
	"Fase2/estructuras/ColaPedidos"
	"Fase2/estructuras/GenerarArchivos"
	"Fase2/estructuras/Peticiones"
	"fmt"
	"math"
	"strconv"
)

type Arbol struct {
	Raiz *NodoArbol
}

func (a *Arbol) altura(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return raiz.Altura
}

func (a *Arbol) equilibrio(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return (a.altura(raiz.Derecho) - a.altura(raiz.Izquierdo))
}

func (a *Arbol) InsertarElemento(id_cliente int, nombre_imagen string) {
	nuevoNodo := &NodoArbol{Valor: &Peticiones.Pedido{Id_Cliente: id_cliente, Nombre_Imagen: nombre_imagen}}
	a.Raiz = a.insertarNodo(a.Raiz, nuevoNodo)
}

func (a *Arbol) rotacionI(raiz *NodoArbol) *NodoArbol {
	raiz_derecho := raiz.Derecho
	hijo_izquierdo := raiz_derecho.Izquierdo
	raiz_derecho.Izquierdo = raiz
	raiz.Derecho = hijo_izquierdo
	numeroMax := math.Max(float64(a.altura(raiz.Izquierdo)), float64(a.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numeroMax)
	numeroMax = math.Max(float64(a.altura(raiz_derecho.Izquierdo)), float64(a.altura(raiz_derecho.Derecho)))
	raiz_derecho.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = a.equilibrio(raiz)
	raiz_derecho.Factor_Equilibrio = a.equilibrio(raiz_derecho)
	return raiz_derecho
}

func (a *Arbol) rotacionD(raiz *NodoArbol) *NodoArbol {
	raiz_izquierdo := raiz.Izquierdo
	hijo_derecho := raiz_izquierdo.Derecho
	raiz_izquierdo.Derecho = raiz
	raiz.Izquierdo = hijo_derecho
	numeroMax := math.Max(float64(a.altura(raiz.Izquierdo)), float64(a.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numeroMax)
	numeroMax = math.Max(float64(a.altura(raiz_izquierdo.Izquierdo)), float64(a.altura(raiz_izquierdo.Derecho)))
	raiz_izquierdo.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = a.equilibrio(raiz)
	raiz_izquierdo.Factor_Equilibrio = a.equilibrio(raiz_izquierdo)
	return raiz_izquierdo
}

func (a *Arbol) insertarNodo(raiz *NodoArbol, nuevoNodo *NodoArbol) *NodoArbol {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Valor.Id_Cliente > nuevoNodo.Valor.Id_Cliente {
			raiz.Izquierdo = a.insertarNodo(raiz.Izquierdo, nuevoNodo)
		} else {
			raiz.Derecho = a.insertarNodo(raiz.Derecho, nuevoNodo)
		}
	}
	numeroMax := math.Max(float64(a.altura(raiz.Izquierdo)), float64(a.altura(raiz.Derecho)))
	raiz.Altura = 1 + int(numeroMax)
	balanceo := a.equilibrio(raiz)
	raiz.Factor_Equilibrio = balanceo
	/*Rotacion simple a la izquierda*/
	if balanceo > 1 && nuevoNodo.Valor.Id_Cliente > raiz.Derecho.Valor.Id_Cliente {
		return a.rotacionI(raiz)
	}
	if balanceo < -1 && nuevoNodo.Valor.Id_Cliente < raiz.Izquierdo.Valor.Id_Cliente {
		return a.rotacionD(raiz)
	}
	if balanceo > 1 && nuevoNodo.Valor.Id_Cliente < raiz.Derecho.Valor.Id_Cliente {
		raiz.Derecho = a.rotacionD(raiz.Derecho)
		return a.rotacionI(raiz)
	}
	if balanceo < -1 && nuevoNodo.Valor.Id_Cliente > raiz.Izquierdo.Valor.Id_Cliente {
		raiz.Izquierdo = a.rotacionI(raiz.Izquierdo)
		return a.rotacionD(raiz)
	}
	return raiz
}

func (a *Arbol) Graficar() {
	cadena := ""
	nombre_archivo := "./Reportes/arbolAVL.dot"
	nombre_imagen := "./Reportes/arbolAVL.jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol{ "
		cadena += a.retornarValoresArbol(a.Raiz, 0)
		cadena += "}"
	}
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(cadena, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (a *Arbol) retornarValoresArbol(raiz *NodoArbol, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
		cadena += " - "
		cadena += (raiz.Valor.Nombre_Imagen) //Esto es para que se muestre la imagen
		cadena += "\" ;"
		if raiz.Izquierdo != nil && raiz.Derecho != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += " - "
			cadena += (raiz.Valor.Nombre_Imagen) //Esto es para que se muestre la imagen
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierdo, numero)
			cadena += "\""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += " - "
			cadena += (raiz.Valor.Nombre_Imagen) //Esto es para que se muestre la imagen
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecho, numero)
			cadena += "{rank=same" + "\"" + strconv.Itoa(raiz.Izquierdo.Valor.Id_Cliente) + " - " + (raiz.Izquierdo.Valor.Nombre_Imagen) + "\"" + " -> " + "\"" + strconv.Itoa(raiz.Derecho.Valor.Id_Cliente) + " - " + (raiz.Derecho.Valor.Nombre_Imagen) + "\"" + " [style=invis]}; "
		} else if raiz.Izquierdo != nil && raiz.Derecho == nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += " - "
			cadena += (raiz.Valor.Nombre_Imagen) //Esto es para que se muestre la imagen
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierdo, numero)
			cadena += "\""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += " - "
			cadena += (raiz.Valor.Nombre_Imagen) //Esto es para que se muestre la imagen
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "{rank=same" + "\"" + strconv.Itoa(raiz.Izquierdo.Valor.Id_Cliente) + " - " + (raiz.Izquierdo.Valor.Nombre_Imagen) + "\"" + " -> " + "x" + strconv.Itoa(numero) + " [style=invis]}; "
		} else if raiz.Izquierdo == nil && raiz.Derecho != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += " - "
			cadena += (raiz.Valor.Nombre_Imagen) //Esto es para que se muestre la imagen
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "; \""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += " - "
			cadena += (raiz.Valor.Nombre_Imagen) //Esto es para que se muestre la imagen
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecho, numero)
			cadena += "{rank=same" + " x" + strconv.Itoa(numero) + " -> \"" + strconv.Itoa(raiz.Derecho.Valor.Id_Cliente) + " - " + (raiz.Derecho.Valor.Nombre_Imagen) + "\"" + " [style=invis]}; "
		}
	}
	return cadena
}

func (a *Arbol) RecorridoInorden(raiz *NodoArbol, colaActual *ColaPedidos.Cola) {
	if raiz != nil {
		if raiz.Izquierdo != nil {
			a.RecorridoInorden(raiz.Izquierdo, colaActual)
		}
		colaActual.Encolar(raiz.Valor.Id_Cliente, raiz.Valor.Nombre_Imagen)
		fmt.Print(" ", raiz.Valor, " ")
		if raiz.Derecho != nil {
			a.RecorridoInorden(raiz.Derecho, colaActual)
		}
	}
}
