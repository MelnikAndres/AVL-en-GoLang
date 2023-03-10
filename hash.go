package diccionario

import (
	"fmt"
	"hash/fnv"
)

type estado int

const (
	OCUPADO estado = iota
	BORRADO
)
const (
	CARGAMAXIMA   = 6556 // 6556 / 3 = 2185 al aumentar la carga se divide
	CARGAMINIMA   = 728  // 728 * 3 = 2184 al reducir la carga se multiplica
	FACTORTAMANO  = 3
	TAMANOINICIAL = 31
)

type elem[K comparable, V any] struct {
	clave K
	valor V
	modo  estado
}

func nuevoElem[K comparable, V any](clave K, valor V) *elem[K, V] {
	return &elem[K, V]{clave, valor, OCUPADO}
}

type hash[K comparable, V any] struct {
	arreglo  []*elem[K, V]
	cant     int
	borrados int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hash[K, V]{make([]*elem[K, V], TAMANOINICIAL), 0, 0}
}

func (h *hash[K, V]) ocupado(pos int) bool {
	return h.arreglo[pos].modo == OCUPADO
}

func (h *hash[K, V]) libre(pos int) bool {
	return h.arreglo[pos] == nil
}

func (h *hash[K, V]) Cantidad() int {
	return h.cant
}
func (h *hash[K, V]) Pertenece(clave K) bool {
	byteado := convertirABytes(clave)
	pos := int(hashear(byteado) % uint64(len(h.arreglo)))
	for i := pos; i < len(h.arreglo); i++ {
		if h.libre(i) {
			return false
		}
		if h.arreglo[i].clave == clave {
			return h.ocupado(i)
		}
	}
	for i := 0; i < pos; i++ {
		if h.libre(i) {
			return false
		}
		if h.arreglo[i].clave == clave {
			return h.ocupado(i)
		}
	}
	return false
}

func (h *hash[K, V]) Guardar(clave K, dato V) {
	h.redimensionar(false)
	byteado := convertirABytes(clave)
	pos := int(hashear(byteado) % uint64(len(h.arreglo)))
	candidato := -1
	for i := pos; i < len(h.arreglo); i++ {
		if h.libre(i) {
			if candidato != -1 {
				h.arreglo[candidato] = nuevoElem(clave, dato)
				h.borrados--
			} else {
				h.arreglo[i] = nuevoElem(clave, dato)
			}
			h.cant++
			return
		}
		if !h.ocupado(i) && candidato == -1 {
			candidato = i
		}
		if h.arreglo[i].clave == clave {
			if !h.ocupado(i) {
				h.arreglo[candidato] = nuevoElem(clave, dato)
				h.borrados--
				h.cant++
				return
			}
			h.arreglo[i].valor = dato
			return
		}
	}
	for i := 0; i < pos; i++ {
		if h.libre(i) {
			if candidato != -1 {
				h.arreglo[candidato] = nuevoElem(clave, dato)
				h.borrados--
			} else {
				h.arreglo[i] = nuevoElem(clave, dato)
			}
			h.cant++
			return
		}
		if !h.ocupado(i) && candidato == -1 {
			candidato = i
		}
		if h.arreglo[i].clave == clave {
			if !h.ocupado(i) {
				h.arreglo[candidato] = nuevoElem(clave, dato)
				h.borrados--
				h.cant++
				return
			}
			h.arreglo[i].valor = dato
			return
		}
	}
	return
}

func (h *hash[K, V]) Obtener(clave K) V {
	byteado := convertirABytes(clave)
	pos := int(hashear(byteado) % uint64(len(h.arreglo)))
	for i := pos; i < len(h.arreglo); i++ {
		if h.libre(i) {
			panic("La clave no pertenece al diccionario")
		}
		if h.arreglo[i].clave == clave {
			if !h.ocupado(i) {
				panic("La clave no pertenece al diccionario")
			}
			return h.arreglo[i].valor
		}
	}
	for i := 0; i < pos; i++ {
		if h.libre(i) {
			panic("La clave no pertenece al diccionario")
		}
		if h.arreglo[i].clave == clave {
			if !h.ocupado(i) {
				panic("La clave no pertenece al diccionario")
			}
			return h.arreglo[i].valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hash[K, V]) Borrar(clave K) V {
	defer h.redimensionar(true)
	byteado := convertirABytes(clave)
	pos := int(hashear(byteado) % uint64(len(h.arreglo)))
	for i := pos; i < len(h.arreglo); i++ {
		if h.libre(i) {
			panic("La clave no pertenece al diccionario")
		}
		if h.arreglo[i].clave == clave {
			if !h.ocupado(i) {
				panic("La clave no pertenece al diccionario")
			}
			h.arreglo[i].modo = BORRADO
			h.cant--
			h.borrados++
			return h.arreglo[i].valor
		}
	}
	for i := 0; i < pos; i++ {
		if h.libre(i) {
			panic("La clave no pertenece al diccionario")
		}
		if h.arreglo[i].clave == clave {
			if !h.ocupado(i) {
				panic("La clave no pertenece al diccionario")
			}
			h.arreglo[i].modo = BORRADO
			h.cant--
			h.borrados++
			return h.arreglo[i].valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := 0; i < len(h.arreglo); i++ {
		if !h.libre(i) && h.ocupado(i) {
			if !visitar(h.arreglo[i].clave, h.arreglo[i].valor) {
				return
			}
		}
	}
}

func (h *hash[K, V]) recolocar(nuevoHash *hash[K, V]) {
	recolocar := func(clave K, dato V) bool {
		nuevoHash.Guardar(clave, dato)
		return true
	}
	h.Iterar(recolocar)
	*h = *nuevoHash
}

func (h *hash[K, V]) esRedimensionable(carga int, borrando bool) (bool, int) {
	if carga >= CARGAMAXIMA {
		return true, len(h.arreglo) * FACTORTAMANO
	}
	if borrando && carga < CARGAMINIMA && len(h.arreglo) > TAMANOINICIAL {
		return true, len(h.arreglo) / FACTORTAMANO
	}
	return false, 0
}

func (h *hash[K, V]) redimensionar(borrando bool) {
	carga := ((h.cant + h.borrados) * 10000) / len(h.arreglo)
	if hayRedimension, tamano := h.esRedimensionable(carga, borrando); hayRedimension {
		nuevoHash := &hash[K, V]{make([]*elem[K, V], tamano), 0, 0}
		h.recolocar(nuevoHash)
	}
}

func (h *hash[K, V]) Iterador() IterDiccionario[K, V] {
	nuevoIterador := new(iterador[K, V])
	nuevoIterador.hasheo = h
	nuevoIterador.pos = nuevoIterador.buscarPrimero()
	return nuevoIterador
}

type iterador[K comparable, V any] struct {
	hasheo *hash[K, V]
	pos    int
}

func (iter *iterador[K, V]) buscarPrimero() int {
	for i := 0; i < len(iter.hasheo.arreglo); i++ {
		if !iter.hasheo.libre(i) && iter.hasheo.ocupado(i) {
			return i
		}
	}
	return -1
}

func (iter *iterador[K, V]) HaySiguiente() bool {
	return iter.pos != -1
}

func (iter *iterador[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.hasheo.arreglo[iter.pos]
	return actual.clave, actual.valor
}

func (iter *iterador[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	clave := iter.hasheo.arreglo[iter.pos].clave
	for i := iter.pos + 1; i < len(iter.hasheo.arreglo); i++ {
		if !iter.hasheo.libre(i) && iter.hasheo.ocupado(i) {
			iter.pos = i
			return clave
		}
	}
	iter.pos = -1
	return clave
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hashear(byteado []byte) uint64 {
	hasheado := fnv.New64a()
	hasheado.Write(byteado)
	return hasheado.Sum64()
}
