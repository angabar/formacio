package signal

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		// Lo que está claro es que no queremos que continuen los tests de
		// manera normal si nisiquiera podemos establecer la conexión con el
		// servidor, por lo que aquí tiene sentido que usemos Fatalf
		t.Fatalf("http.NewRequest() err = %s", err)
	}

	Handler(w, r)

	resp := w.Result()
	if resp.StatusCode != 200 {
		// Una situación similar pasa cuando no obtenemos una respuesta que
		// tiene un 200, ya que lo más seguro es que si no es un 200, es porque
		// hay un error y no tiene sentido seguir con los tests
		t.Fatalf("Handler() status = %d; want %d", resp.StatusCode, 200)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		// En el caso de que no tengamos exactamente el header de
		// application/json podemos continuar con los tests, ya que no es una
		// señal inequivoca de que la petición haya fallado
		t.Errorf("Header() Content-Type = %q; want %q", contentType, "application/json")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Si la lectura de los datos no nos permite extraer el objeto data, si
		// que tenemos un problema, ya que el resto de los tests se basan en
		// este para seguir haciendo comprobaciones, por lo tanto, en este caso
		// si que tiene sentio usar Fatal
		t.Fatalf("ioutil.ReadAll(resp.Body) err = %s", err)
	}

	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		// Lo mismo pasa si no podemos parsear el objeto json
		t.Fatalf("json.Unmashal(resp.Body) err = %s", err)
	}

	wantAge := 21
	if p.Age != wantAge {
		// No obstante, para el caso en que tengamos que comprobar que los
		// atributos del objeto tengan cierto valor, con un Error nos sobra, ya
		// que podmeos continuar sin problemas con el resto de los tests
		t.Errorf("person.Age = %d; want %d", p.Age, wantAge)
	}

	wantName := "John Calhoun"
	if p.Name != wantName {
		t.Errorf("person.Name = %s; want %s", p.Name, wantName)
	}
}

func TestPick(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arg := make([]int, 5)
	for i := 0; i < 5; i++ {
		arg[i] = r.Int()
	}

	got := Pick(arg)
	if got != arg[2] {
		t.Errorf("Pick(%v) = %d; want %d", arg, got, arg[2])
	}
}

// Cuando escribimos mensajes de error, tenemos que ser concisos, no escribir
// cosas como "crap broke" que no aportan nada, siempre hay que intentar dar
// toda la información posible para ayudar al desarrollador a solucionar el
// problema

// Es muy importante seguir la nomenclatura de, que estoy llamando, que es lo
// que me devuelve, y que es lo que quiero, el orden es importante, function -
// got - want

// t.Errorf("person.Name = %s; want %s", p.Name, wantName)

// En principio no es necesario poner los argumentos de las funciones a las que
// llamamos, pero esto es como todo, depende de lo que estmeos testeando y lo
// que busquemos, unas veces tendrá sentido ponerlos, otras no

// Se conciso, prueba e imprime atributos concretos en lugar de objetos enteros
