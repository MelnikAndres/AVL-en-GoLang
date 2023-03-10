package diccionario_test

import (
	TDAABB "diccionario"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dicAbb := TDAABB.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dicAbb.Cantidad())
	require.False(t, dicAbb.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicAbb.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicAbb.Borrar("A") })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dicAbb := TDAABB.CrearABB[string, int](strings.Compare)
	dicAbb.Guardar("A", 10)
	require.EqualValues(t, 1, dicAbb.Cantidad())
	require.True(t, dicAbb.Pertenece("A"))
	require.False(t, dicAbb.Pertenece("B"))
	require.EqualValues(t, 10, dicAbb.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicAbb.Obtener("B") })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dicAbb := TDAABB.CrearABB[string, string](strings.Compare)
	require.False(t, dicAbb.Pertenece(claves[0]))
	dicAbb.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dicAbb.Cantidad())
	require.True(t, dicAbb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dicAbb.Obtener(claves[0]))

	require.False(t, dicAbb.Pertenece(claves[1]))
	require.False(t, dicAbb.Pertenece(claves[2]))
	dicAbb.Guardar(claves[1], valores[1])
	require.True(t, dicAbb.Pertenece(claves[0]))
	require.True(t, dicAbb.Pertenece(claves[1]))
	require.EqualValues(t, 2, dicAbb.Cantidad())
	require.EqualValues(t, valores[0], dicAbb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dicAbb.Obtener(claves[1]))

	require.False(t, dicAbb.Pertenece(claves[2]))
	dicAbb.Guardar(claves[2], valores[2])
	require.True(t, dicAbb.Pertenece(claves[0]))
	require.True(t, dicAbb.Pertenece(claves[1]))
	require.True(t, dicAbb.Pertenece(claves[2]))
	require.EqualValues(t, 3, dicAbb.Cantidad())
	require.EqualValues(t, valores[0], dicAbb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dicAbb.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dicAbb.Obtener(claves[2]))
}

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dicAbb := TDAABB.CrearABB[string, string](strings.Compare)
	dicAbb.Guardar(clave, "miau")
	dicAbb.Guardar(clave2, "guau")
	require.True(t, dicAbb.Pertenece(clave))
	require.True(t, dicAbb.Pertenece(clave2))
	require.EqualValues(t, "miau", dicAbb.Obtener(clave))
	require.EqualValues(t, "guau", dicAbb.Obtener(clave2))
	require.EqualValues(t, 2, dicAbb.Cantidad())

	dicAbb.Guardar(clave, "miu")
	dicAbb.Guardar(clave2, "baubau")
	require.True(t, dicAbb.Pertenece(clave))
	require.True(t, dicAbb.Pertenece(clave2))
	require.EqualValues(t, 2, dicAbb.Cantidad())
	require.EqualValues(t, "miu", dicAbb.Obtener(clave))
	require.EqualValues(t, "baubau", dicAbb.Obtener(clave2))
}

func TestDiccionarioBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dicAbb := TDAABB.CrearABB[string, string](strings.Compare)

	require.False(t, dicAbb.Pertenece(claves[0]))
	dicAbb.Guardar(claves[0], valores[0])
	dicAbb.Guardar(claves[1], valores[1])
	dicAbb.Guardar(claves[2], valores[2])

	require.True(t, dicAbb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dicAbb.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicAbb.Borrar(claves[2]) })
	require.EqualValues(t, 2, dicAbb.Cantidad())
	require.False(t, dicAbb.Pertenece(claves[2]))

	require.True(t, dicAbb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dicAbb.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicAbb.Borrar(claves[0]) })
	require.EqualValues(t, 1, dicAbb.Cantidad())
	require.False(t, dicAbb.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicAbb.Obtener(claves[0]) })

	require.True(t, dicAbb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dicAbb.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicAbb.Borrar(claves[1]) })
	require.EqualValues(t, 0, dicAbb.Cantidad())
	require.False(t, dicAbb.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicAbb.Obtener(claves[1]) })
}

func TestConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dicAbb := TDAABB.CrearABB[int, string](compararInt)
	clave := 10
	valor := "Gatito"

	dicAbb.Guardar(clave, valor)
	require.EqualValues(t, 1, dicAbb.Cantidad())
	require.True(t, dicAbb.Pertenece(clave))
	require.EqualValues(t, valor, dicAbb.Obtener(clave))
	require.EqualValues(t, valor, dicAbb.Borrar(clave))
	require.False(t, dicAbb.Pertenece(clave))
}

func TestConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	dicAbb := TDAABB.CrearABB[avanzado, int](func(a avanzado, b avanzado) int {
		if a == b {
			return 0
		} else if a.x.a > b.x.a {
			return 1
		} else {
			return -1
		}
	})

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dicAbb.Guardar(a1, 0)
	dicAbb.Guardar(a2, 1)
	dicAbb.Guardar(a3, 2)

	require.True(t, dicAbb.Pertenece(a1))
	require.True(t, dicAbb.Pertenece(a2))
	require.True(t, dicAbb.Pertenece(a3))
	require.EqualValues(t, 0, dicAbb.Obtener(a1))
	require.EqualValues(t, 1, dicAbb.Obtener(a2))
	require.EqualValues(t, 2, dicAbb.Obtener(a3))
	dicAbb.Guardar(a1, 5)
	require.EqualValues(t, 5, dicAbb.Obtener(a1))
	require.EqualValues(t, 2, dicAbb.Obtener(a3))
	require.EqualValues(t, 5, dicAbb.Borrar(a1))
	require.False(t, dicAbb.Pertenece(a1))
	require.EqualValues(t, 2, dicAbb.Obtener(a3))

}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dicAbb := TDAABB.CrearABB[string, string](strings.Compare)
	clave := ""
	dicAbb.Guardar(clave, clave)
	require.True(t, dicAbb.Pertenece(clave))
	require.EqualValues(t, 1, dicAbb.Cantidad())
	require.EqualValues(t, clave, dicAbb.Obtener(clave))
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dicAbb := TDAABB.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dicAbb.Guardar(clave, nil)
	require.True(t, dicAbb.Pertenece(clave))
	require.EqualValues(t, 1, dicAbb.Cantidad())
	require.EqualValues(t, (*int)(nil), dicAbb.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dicAbb.Borrar(clave))
	require.False(t, dicAbb.Pertenece(clave))
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dicAbb := TDAABB.CrearABB[string, *int](strings.Compare)
	dicAbb.Guardar(claves[0], nil)
	dicAbb.Guardar(claves[1], nil)
	dicAbb.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0

	dicAbb.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		cantidad += 1
		return true
	})

	require.EqualValues(t, cantidad, dicAbb.Cantidad())
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func compararInt(dato1 int, dato2 int) int {
	if dato1 > dato2 {
		return 1
	} else if dato1 < dato2 {
		return -1
	}
	return 0
}

func TestIteradorInternoSinCorte(t *testing.T) {
	t.Log("Prueba el iterador interno sin ningun corte sumando los elementos del diccionario")
	dicAbb := TDAABB.CrearABB[int, int](compararInt)
	claves := [10]int{5, 7, 3, 8, 1, 4, 6, 2, 9, 10}
	for i, clave := range claves {
		dicAbb.Guardar(clave, i+1)
	}
	contador := 0
	dicAbb.Iterar(func(dato1 int, dato2 int) bool {
		contador += dato2
		return true
	})
	require.Equal(t, 55, contador)
}

func TestIteradorInternoConCorte(t *testing.T) {
	t.Log("Prueba el iterador interno con un corte en la suma de elementos del diccionario")
	dicAbb := TDAABB.CrearABB[int, int](compararInt)
	claves := [10]int{5, 7, 3, 8, 1, 4, 6, 2, 9, 10}
	for _, clave := range claves {
		dicAbb.Guardar(clave, clave)
	}

	contador := 0
	dicAbb.Iterar(func(dato1 int, dato2 int) bool {
		if dato1 == 4 {
			return false
		}
		contador += dato2
		return true
	})
	require.EqualValues(t, 6, contador)
}

func TestIteradorInternoConRangoSinCorte(t *testing.T) {
	t.Log("Prueba el iterador interno con rango entre 10 y 15 sumando sus datos sin ningun corte")
	dicAbb := TDAABB.CrearABB[int, int](compararInt)
	rangoInicial := 10
	rangoFinal := 15
	dicAbb.Guardar(12, 1)
	dicAbb.Guardar(10, 1)
	dicAbb.Guardar(5, 1)
	dicAbb.Guardar(15, 1)
	dicAbb.Guardar(9, 1)
	dicAbb.Guardar(13, 1)
	dicAbb.Guardar(16, 1)
	dicAbb.Guardar(14, 1)
	dicAbb.Guardar(11, 1)
	dicAbb.Guardar(21, 1)

	contador := 0
	dicAbb.IterarRango(&rangoInicial, &rangoFinal, func(dato1 int, dato2 int) bool {
		contador += dato2
		return true
	})
	require.EqualValues(t, 6, contador)
}

func TestIteradorInternoConRangoConCorte(t *testing.T) {
	t.Log("Prueba el iterador interno con rango entre 10 y 15 sumando sus datos pero con un corte")
	dicAbb := TDAABB.CrearABB[int, int](compararInt)
	rangoInicial := 10
	rangoFinal := 15
	dicAbb.Guardar(12, 1)
	dicAbb.Guardar(10, 1)
	dicAbb.Guardar(5, 1)
	dicAbb.Guardar(15, 1)
	dicAbb.Guardar(9, 1)
	dicAbb.Guardar(13, 1)
	dicAbb.Guardar(16, 1)
	dicAbb.Guardar(14, 1)
	dicAbb.Guardar(11, 1)
	dicAbb.Guardar(21, 1)

	contador := 0
	dicAbb.IterarRango(&rangoInicial, &rangoFinal, func(dato1 int, dato2 int) bool {
		if dato1 == 13 { //Corte si se llega al tercer elemento
			return false
		}
		contador += dato2
		return true
	})
	require.EqualValues(t, 3, contador)
}

func TestIteradorInternoConRangoNil(t *testing.T) {
	t.Log("Prueba el iterador interno con rango entre 10 y 15 sumando sus datos sin ningun corte")
	dicAbb := TDAABB.CrearABB[int, int](compararInt)
	var rangoInicial *int
	var rangoFinal *int
	dicAbb.Guardar(12, 1)
	dicAbb.Guardar(10, 1)
	dicAbb.Guardar(5, 1)
	dicAbb.Guardar(15, 1)
	dicAbb.Guardar(9, 1)
	dicAbb.Guardar(13, 1)
	dicAbb.Guardar(16, 1)
	dicAbb.Guardar(14, 1)
	dicAbb.Guardar(11, 1)
	dicAbb.Guardar(21, 1)

	contador := 0
	dicAbb.IterarRango(rangoInicial, rangoFinal, func(dato1 int, dato2 int) bool {
		contador += dato2
		return true
	})
	require.EqualValues(t, 10, contador)
}

func TestIteradorExternoSinRango(t *testing.T) {
	t.Log("Prueba el iterador externo comun sin rango")
	dicAbb := TDAABB.CrearABB[int, int](compararInt)
	claves := [13]int{1, 32, 5, 76, 2, 6, 48, 10, 14, 3, 29, 47, 82}
	clavesOrdenadas := [13]int{1, 2, 3, 5, 6, 10, 14, 29, 32, 47, 48, 76, 82}
	for i, clave := range claves {
		dicAbb.Guardar(clave, i)
	}
	iter := dicAbb.Iterador()
	i := 0
	for iter.HaySiguiente() {
		claveActual, _ := iter.VerActual()
		require.EqualValues(t, clavesOrdenadas[i], claveActual)
		i++
		iter.Siguiente()
	}
	require.EqualValues(t, false, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIteradorExternoConRango(t *testing.T) {
	t.Log("Prueba el iterador externo con un rango entre 10 y 30")
	dicAbb := TDAABB.CrearABB[int, int](compararInt)
	rangoInicial := 10
	rangoFinal := 30
	claves := [13]int{1, 32, 5, 76, 2, 6, 48, 10, 14, 3, 29, 47, 82}
	clavesOrdenadas := [13]int{1, 2, 3, 5, 6, 10, 14, 29, 32, 47, 48, 76, 82}
	for i, clave := range claves {
		dicAbb.Guardar(clave, i)
	}
	iter := dicAbb.IteradorRango(&rangoInicial, &rangoFinal)
	i := 5
	for iter.HaySiguiente() {
		claveActual, _ := iter.VerActual()
		require.EqualValues(t, clavesOrdenadas[i], claveActual)
		i++
		iter.Siguiente()
	}
	require.EqualValues(t, false, iter.HaySiguiente())
	require.EqualValues(t, 8, i)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIteradorExternoConRangoNil(t *testing.T) {
	t.Log("Prueba el iterador externo con un rango entre 10 y 30")
	dicAbb := TDAABB.CrearABB[int, int](compararInt)
	var rangoInicial *int
	var rangoFinal *int
	claves := [13]int{1, 32, 5, 76, 2, 6, 48, 10, 14, 3, 29, 47, 82}
	clavesOrdenadas := [13]int{1, 2, 3, 5, 6, 10, 14, 29, 32, 47, 48, 76, 82}
	for i, clave := range claves {
		dicAbb.Guardar(clave, i)
	}
	iter := dicAbb.IteradorRango(rangoInicial, rangoFinal)
	i := 0
	for iter.HaySiguiente() {
		claveActual, _ := iter.VerActual()
		require.EqualValues(t, clavesOrdenadas[i], claveActual)
		i++
		iter.Siguiente()
	}
	require.EqualValues(t, false, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIteradorExternoConRangoZigZag(t *testing.T) {
	t.Log("Prueba el iterador externo con un rango entre 111 y 112 e insertando de forma que el diccionario" +
		"inserte elementos en forma de zig zag")
	dicAbb := TDAABB.CrearABB[int, int](compararInt)
	rangoInicial := 111
	rangoFinal := 112
	dicAbb.Guardar(100, 0)
	dicAbb.Guardar(120, 0)
	dicAbb.Guardar(114, 0)
	dicAbb.Guardar(112, 0)
	dicAbb.Guardar(111, 0)

	iter := dicAbb.IteradorRango(&rangoInicial, &rangoFinal)
	i := 0
	for iter.HaySiguiente() {
		claveActual, _ := iter.VerActual()
		require.EqualValues(t, 111+i, claveActual)
		i++
		iter.Siguiente()
	}

}

func TestIteradoresArbolVacio(t *testing.T) {
	t.Log("Comprueba los iteradores con un arbol vacio")
	dicAbb := TDAABB.CrearABB[int, int](compararInt)
	require.EqualValues(t, 0, dicAbb.Cantidad())

	cantidad := 0
	dicAbb.Iterar(func(clave int, dato int) bool {
		cantidad++
		return true
	})
	require.EqualValues(t, 0, cantidad)

	iterAbb := dicAbb.Iterador()
	require.EqualValues(t, false, iterAbb.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterAbb.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterAbb.Siguiente() })
}
