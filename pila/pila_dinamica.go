package pila

const _CAPACIDAD_INICIAL = 30
const _FACTOR_REDIMENSION = 2

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) redimensionar(tamanoNuevo int) {
	nuevosDatos := make([]T, tamanoNuevo)
	copy(nuevosDatos, pila.datos)
	pila.datos = nuevosDatos
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
	if pila.cantidad > len(pila.datos)-1 {
		pila.redimensionar(len(pila.datos) * _FACTOR_REDIMENSION)
	}
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	elementoTope := pila.datos[pila.cantidad-1]
	pila.cantidad--
	if pila.cantidad*4 <= len(pila.datos)-1 {
		pila.redimensionar(len(pila.datos) / _FACTOR_REDIMENSION)
	}
	return elementoTope
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAPACIDAD_INICIAL)
	pila.cantidad = 0
	return pila
}
