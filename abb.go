package diccionario

import (
	TDAPila "diccionario/pila"
)

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
	vDer  int
	vIzq  int
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{nil, nil, clave, dato, 0, 0}
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	comparar func(K, K) int
	cant     int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, funcion_cmp, 0}
}

type iterAbb[K comparable, V any] struct {
	pilaIter TDAPila.Pila[*nodoAbb[K, V]]
}

type iterAbbRango[K comparable, V any] struct {
	pilaIter TDAPila.Pila[*nodoAbb[K, V]]
	desde    *K
	hasta    *K
	comparar func(K, K) int
}

func (a *abb[K, V]) Cantidad() int {
	return a.cant
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	actual := a.raiz
	comparacion := 0
	for actual != nil {
		comparacion = a.comparar(clave, actual.clave)
		if comparacion > 0 {
			actual = actual.der
		} else if comparacion < 0 {
			actual = actual.izq
		} else {
			return true
		}
	}
	return false
}

func (a *abb[K, V]) Obtener(clave K) V {
	actual := a.raiz
	comparacion := 0
	for actual != nil {
		comparacion = a.comparar(clave, actual.clave)
		if comparacion > 0 {
			actual = actual.der
		} else if comparacion < 0 {
			actual = actual.izq
		} else {
			return actual.dato
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterando := true
	a.iterar(&visitar, a.raiz, &iterando)
}

func (a *abb[K, V]) iterar(visitar *func(clave K, dato V) bool, actual *nodoAbb[K, V], iterando *bool) {
	//inorder,  itera de menor a mayor
	if actual == nil {
		return
	}
	a.iterar(visitar, actual.izq, iterando)
	if !*iterando || !(*visitar)(actual.clave, actual.dato) {
		*iterando = false
		return
	}
	a.iterar(visitar, actual.der, iterando)
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterando := true

	desdeFn := func(nodo *nodoAbb[K, V]) bool {
		return desde == nil || a.comparar(nodo.clave, *desde) >= 0
	}
	hastaFn := func(nodo *nodoAbb[K, V]) bool {
		return hasta == nil || a.comparar(nodo.clave, *hasta) <= 0
	}
	enRango := func(nodo *nodoAbb[K, V]) bool {
		return desdeFn(nodo) && hastaFn(nodo)
	}

	a.iterarRango(&enRango, &visitar, a.raiz, &iterando)
}

func (a *abb[K, V]) iterarRango(enRango *func(ab *nodoAbb[K, V]) bool, visitar *func(clave K, dato V) bool, actual *nodoAbb[K, V], iterando *bool) {
	//inorder,  itera de menor a mayor
	if actual == nil {
		return
	}
	a.iterarRango(enRango, visitar, actual.izq, iterando)
	if !*iterando {
		return
	}
	if (*enRango)(actual) {
		if !(*visitar)(actual.clave, actual.dato) {
			*iterando = false
			return
		}
	}
	a.iterarRango(enRango, visitar, actual.der, iterando)
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	nuevoIter := new(iterAbb[K, V])
	nuevoIter.pilaIter = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	nuevoIter.apilarRama(a.raiz)
	return nuevoIter
}

func (iter *iterAbb[K, V]) apilarRama(actual *nodoAbb[K, V]) {
	for actual != nil {
		iter.pilaIter.Apilar(actual)
		actual = actual.izq
	}
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pilaIter.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pilaIter.VerTope()
	return actual.clave, actual.dato
}

func (iter *iterAbb[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := iter.pilaIter.Desapilar()
	iter.apilarRama(dato.der)
	return dato.clave
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	nuevoIter := new(iterAbbRango[K, V])
	nuevoIter.pilaIter = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	nuevoIter.hasta = hasta
	nuevoIter.desde = desde
	nuevoIter.comparar = a.comparar
	nuevoIter.apilarRamaIzq(nuevoIter.cotaBaja(a.raiz))
	return nuevoIter
}

func (iter *iterAbbRango[K, V]) sobreCotaBaja(nodo *nodoAbb[K, V]) bool {
	return iter.desde == nil || iter.comparar(nodo.clave, *iter.desde) >= 0
}

func (iter *iterAbbRango[K, V]) debajoCotaAlta(nodo *nodoAbb[K, V]) bool {
	return iter.hasta == nil || iter.comparar(nodo.clave, *iter.hasta) <= 0
}

func (iter *iterAbbRango[K, V]) cotaBaja(actual *nodoAbb[K, V]) *nodoAbb[K, V] {
	for actual != nil {
		if iter.sobreCotaBaja(actual) {
			if !iter.debajoCotaAlta(actual) {
				actual = iter.cotaAlta(actual)
			}
			return actual
		}
		actual = actual.der
	}
	return nil
}

func (iter *iterAbbRango[K, V]) cotaAlta(actual *nodoAbb[K, V]) *nodoAbb[K, V] {
	for actual != nil {
		if iter.debajoCotaAlta(actual) {
			if !iter.sobreCotaBaja(actual) {
				actual = iter.cotaBaja(actual)
			}
			return actual
		}
		actual = actual.izq
	}
	return nil
}

func (iter *iterAbbRango[K, V]) apilarRamaIzq(actual *nodoAbb[K, V]) {
	for actual != nil {
		if !iter.sobreCotaBaja(actual) {
			actual = actual.der
		} else {
			iter.pilaIter.Apilar(actual)
			actual = actual.izq
		}
	}
}

func (iter *iterAbbRango[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pilaIter.VerTope()
	return actual.clave, actual.dato
}

func (iter *iterAbbRango[K, V]) HaySiguiente() bool {
	return !iter.pilaIter.EstaVacia()
}

func (iter *iterAbbRango[K, V]) Siguiente() K {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pilaIter.Desapilar()
	if iter.pilaIter.EstaVacia() {
		iter.apilarRamaIzq(iter.cotaBaja(actual.der))
	} else {
		iter.apilarRamaIzq(actual.der)
	}
	return actual.clave
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	a.guardar(clave, dato, &a.raiz)
}

func (a *abb[K, V]) guardar(clave K, dato V, z **nodoAbb[K, V]) int {
	if *z == nil {
		*z = crearNodoAbb(clave, dato)
		a.cant++
		return 1
	}
	comparacion := a.comparar(clave, (*z).clave)
	if comparacion > 0 {
		vDer := a.guardar(clave, dato, &(*z).der)
		(*z).vDer = vDer
		if (*z).vIzq-(*z).vDer < -1 {
			a.equilibrarIzquierda(z, (*z).der.der == nil || (*z).der.vDer < (*z).der.vIzq)
		}
		return (*z).vDer + 1
	} else if comparacion < 0 {
		vIzq := a.guardar(clave, dato, &(*z).izq)
		(*z).vIzq = vIzq
		if (*z).vDer-(*z).vIzq < -1 {
			a.equilibrarDerecha(z, (*z).izq.izq == nil || (*z).izq.vIzq < (*z).izq.vDer)
		}
		return (*z).vIzq + 1
	} else {
		(*z).dato = dato
		return max((*z).vIzq, (*z).vDer) + 1
	}
}

func (a *abb[K, V]) Raiz() *nodoAbb[K, V] {
	return a.raiz

}
func (n *nodoAbb[K, V]) Clave() K {
	return n.clave

}

func (n *nodoAbb[K, V]) Der() *nodoAbb[K, V] {
	return n.der
}
func (n *nodoAbb[K, V]) Izq() *nodoAbb[K, V] {
	return n.izq
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	if a < b {
		return b
	}
	return a
}

func (a *abb[K, V]) Borrar(clave K) V {
	var dato V
	a.borrar(clave, &a.raiz, &dato)
	return dato

}

func (a *abb[K, V]) borrar(clave K, z **nodoAbb[K, V], dato *V) int {
	if *z == nil {
		panic("La clave no pertenece al diccionario")
	}
	comparacion := a.comparar(clave, (*z).clave)
	if comparacion > 0 {
		vDer := a.borrar(clave, &(*z).der, dato) //borrando en derecha
		(*z).vDer += vDer
		if (*z).vDer-(*z).vIzq < -1 {
			a.equilibrarDerecha(z, (*z).izq.izq == nil || (*z).izq.vIzq < (*z).izq.vDer)
		}
		if (*z).vIzq == (*z).vDer { //no se borro sobre el camino mas largo
			return vDer
		}
		return 0
	} else if comparacion < 0 {
		vIzq := a.borrar(clave, &(*z).izq, dato) //borrando en izquierda
		(*z).vIzq += vIzq
		if (*z).vIzq-(*z).vDer < -1 {
			a.equilibrarIzquierda(z, (*z).der.der == nil || (*z).der.vDer < (*z).der.vIzq)
		}
		if (*z).vDer == (*z).vIzq { // no se borro el camino mas largo
			return vIzq
		}
		return 0
	} else {
		*dato = (*z).dato
		n := a._borrar(z)
		if *z == nil { //borrado sin hijos
			return -1
		}
		if (*z).vDer-(*z).vIzq < -1 { // desequilibrio en izquierda
			a.equilibrarDerecha(z, (*z).izq.izq == nil || (*z).izq.vIzq < (*z).izq.vDer)
		}

		if (*z).vIzq == (*z).vDer { //no se borro el camino mas largo
			return n
		}
		return 0
	}
}

func (a *abb[K, V]) _borrar(actual **nodoAbb[K, V]) int {
	hayIzq := (*actual).izq != nil
	hayDer := (*actual).der != nil
	a.cant--

	if hayIzq && hayDer {
		derecho := &(*actual).der
		reemplazo, n := a.borrarDosHIjos(derecho) //rama de mayores

		(*actual).clave = (*reemplazo).clave
		(*actual).dato = (*reemplazo).dato
		(*actual).vDer += n //solo se altera en derecha

		*reemplazo = (*reemplazo).der
		if *derecho != nil && (*derecho).vIzq-(*derecho).vDer < -1 {
			a.equilibrarIzquierda(derecho, (*derecho).der.der == nil || (*derecho).der.vDer < (*derecho).der.vIzq)
		}
		return n
	} else if !hayIzq && !hayDer {
		*actual = nil
		return -1
	} else if !hayDer {
		*actual = (*actual).izq
		return -1
	} else if !hayIzq {
		*actual = (*actual).der
		return -1
	}
	a.cant--
	return 0

}

func (a *abb[K, V]) borrarDosHIjos(z **nodoAbb[K, V]) (**nodoAbb[K, V], int) {
	if (*z).izq == nil {
		zCopia := *z
		(*z) = (*z).der
		return &zCopia, -1
	}
	nodo, vDir := a.borrarDosHIjos(&(*z).izq)
	(*z).vIzq += vDir
	if (*z).vIzq-(*z).vDer < -1 {
		a.equilibrarIzquierda(z, (*z).der.der == nil || (*z).der.vDer < (*z).der.vIzq)
	}
	if (*z).vDer == (*z).vIzq { // no se borro el camino mas largo
		return nodo, vDir
	}
	return nodo, 0
}

func (a *abb[K, V]) equilibrarDerecha(z **nodoAbb[K, V], doble bool) {
	if doble {
		a.reordenarDer(&(*z).izq, &(*z).izq.der)
	}
	a.reordenarIzq(z, &(*z).izq)
}

func (a *abb[K, V]) equilibrarIzquierda(z **nodoAbb[K, V], doble bool) {
	if doble {
		a.reordenarIzq(&(*z).der, &(*z).der.izq)
	}
	a.reordenarDer(z, &(*z).der)
}

func (a *abb[K, V]) reordenarDer(z **nodoAbb[K, V], y **nodoAbb[K, V]) {
	yCopia := (*z).der
	(*z).der = (*y).izq
	zCopia := *z
	*z = yCopia
	(*z).izq = zCopia
	zCopia.vDer = (*z).vIzq
	(*z).vIzq = max(zCopia.vDer, zCopia.vIzq) + 1
}

func (a *abb[K, V]) reordenarIzq(z **nodoAbb[K, V], y **nodoAbb[K, V]) {
	yCopia := (*z).izq
	(*z).izq = (*y).der
	zCopia := *z
	*z = yCopia
	(*z).der = zCopia
	zCopia.vIzq = (*z).vDer
	(*z).vDer = max(zCopia.vDer, zCopia.vIzq) + 1
}

func (a *abb[K, V]) ObtenerNodo(clave K) *nodoAbb[K, V] {
	actual := a.raiz
	comparacion := 0
	for actual != nil {
		comparacion = a.comparar(clave, actual.clave)
		if comparacion > 0 {
			actual = actual.der
		} else if comparacion < 0 {
			actual = actual.izq
		} else {
			return actual
		}
	}
	return nil
}
