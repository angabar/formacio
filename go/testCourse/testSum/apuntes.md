# Bases del testing

Los test son tan solo código de Go, usaremos la libreria `testing` para poder tener acceso a métodos y variables relacionadas con el test.

```go
package utils

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum([]int{1, 2, 3})
	if sum != 6 {
		t.Errorf("Wanted 11 but got %d", sum)
	}
}
```

### Convecnión de nombres de archivos

A la hora de nombrar archivos tendremos en cuenta los siguientes criterios:

-   `export_test.go` para poder acceder a la variables internas de los tests
-   `xxx_internal_test.go` para los tests internos
-   `example_xxx_test.go` para los ejemplos

### Convención de nombres de funciones

Las funciones de testing empezarán siempre por la palabra `Test` en mayúscula seguida por el nombre de la función a testear.

`TestSum`

`TestAdd`

En el caso de que tengamos una función reciever de un tipo determinado, entonces pondremos el nombre relacionando primero al tipo y después al reciever `TestType_Reciver`

`TestDog_Bark`

Cualquier otra cosa que pongamos para especificar algún aspecto relevante de la función, irá seguida de un guión bajo.

`TestDog_Bark_muzzled`

### Convención de nombres de las variables

Además de tener en cuenta lo explicado en el capítulo de las variables, tendremos en cuenta las siguientes reglas.

Guardaremos lo que esperamos en una variable llamda `want` y lo que tenemos en una variable llamada `got`

En el caso de que tengamos un método al que le pasamos un argumento podemos guardar este en una variable `arg`

```go
func TestColor(t *testing.T) {
    arg := "blue"
    want := "#000FE"
    got := Color(arg)

    if got != want {
        t.Errorf("Color(%q) = %s; want $s", arg, got, want)
    }
}
```

### Mostrado de mensajes

Cuando ocurre un error podemos mostrar un mensaje por pantalla, para esto podemos la librería `testing` nos presenta una serie de métodos a elegir.

`Log` `Logf`

Muestran un mensaje por pantalla ya sea de tipo simple `Log` o con variables concatenadas `Logf`

```go
func TestHandler(t *testing.T) {
    t.Log("here is my message")
    t.Logf("here is my message with number: %d", 123)

    // Sin un error, los Log no se muestran
    t.Error("failing test")
}
```

`Fail` `FailNow`

`Fail` hará que el test falle pero continue con la ejecución del código, mientras que `FailNow` hará que el test falle y se detenga la ejecución del código.

`Fatal` `Fatalf`

`Fatal` es la combinación de `Log` y `FailNow` mientras que `Fatalf` es la combinación de `Logf` y `FailNow`

`Error` `Errorf`

`Error` es la combinación de `Log` y `Fail` mientras que `Errorf` es la combinación de `Logf` y `Fail`

Cuando usar uno u otro depende de si queremos continuar la ejecución del test para obtener más información, o no es necesario continuar el test pues no vamos a obtener más información.

Por ejemplo, no tiene sentido continuar un test en el que a fallado una llamada al servidor, por lo tanto en este caso usaremos `Fatal` o `Fatalf`

```go
func TestHandler(t *testing.T) {
    w := httpRequest.NewRecorder()
    r, err := http.NewRequest(http.MethodGet, "", nil)
    if err != nil {
        t.Fatalf("...")
    }
}
```

También podriamos usar `Fatal` o `Fatalf` en el caso de que la respuesta del servidor no sea 200, ya que no tiene sentido seguir el test si no vamos a tener una respuesta válida.

```go
func TestHandler(t *testing.T) {
    ...

    resp := w.Result()
    if resp.StatusCode != 200 {
        t.Fatalf(...)
    }
}
```

En el caso por ejemplo de que el `header` no sea exactamente un `application/json` quizás no sea necesario lanzar un `Fatal` sino un `Error`

```go
func TestHandler(t *testing.T) {
    ...

    contentType := resp.Header.Get("Content-Type")
    if contentType != "application/json" {
        t.Errorf(...)
    }
}
```

En el caso de que no podamos recibir el objeto `data` de la llamada tendríamos que lanzar un `Fatal` ya que el resto del test queda inservible.

```go
func TestHandler(t *testing.T) {
    ...

    data, err := ioutil.ReadAll(resp.body)
    if err != nil {
        t.Fatalf(...)
    }
}
```

Lo mismo para el caso de que la conversión del `json` falle por algún motivo.

```go
func TestHandler(t *testing.T) {
    ...

    var p Person
    err = json.Unmarshal(data, &p)
    if err != nil {
        t.Fatalf(...)
    }
}
```

Sin embargo para hacer comprobaciones de si los atributos del objeto traducido tienen uno u otro valor, quizás sea suficiente con un `Error`

```go
func TestHandler(t *testing.T) {
    ...

    if p.Age != 21 {
        t.Errorf(...)
    }
}
```
