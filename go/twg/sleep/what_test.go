package sleep

import (
	"testing"
	"time"
)

// Cuando se lanzan los tests, Go crea un archivo binario temporal que lanza en
// el momento y que es destruido cuando terminan los tests

// Poninendo esta funcion sleep podemos buscar y ver que dicho binario existe
// durante el minuto que dura la espera, incluso podemos ejecutarlo

// A la hora de nombrar los archivos de los tests tenemos que tener en cuenta
// tres tipos de nomenclaturas, las normales, son aquellas que pertenecen a los
// archivos de ejecución que estmos testeando y se escriben con el mismo nombre
// del archivo y _test al final para diferenciarlo, los internos, con
// _internal_test y por último example_xxx_test para ejemplos

// Las funciones de los test se nombran empezando con "Test" en mayúscula
// seguido del nombre que vamos a testear en camelCase (por ejemplo TestDog)

// En el caso de que tengamos un test que deba probar diferentes escenarios los
// especificaremos usando _ (por ejemplo TesDog_Bark)
func TestTmpExecutable(t *testing.T) {
	time.Sleep(time.Minute)
}

// Otro aspecto importante en los tests, es la forma en la que llamamos a las
// variables, en Go tenemos principalmente tres posibles variables en los tests,
// arg para definir el argumento que pasamos a la funcion a testear (en caso de
// haber alguno), want para definir el valor esperado y got para definir el
// valor que tenemos
/* func TestColor(t *testing.T) {
	arg := "blue"
	want := "#00000F"
	got := Color(arg)

	if got != want {
		t.Errorf("Color(%q) = %q; want = %q", arg, got, want)
	}
} */

// Cuando queremos logear información en caso de que el test falle, tenemos que
// usar el método t.Log (t.Logf para el caso en que queramos parsear argumentos
// en el string) de la librería de testing. Estos string solo se mostrarán en
// caso de que fallen los tests o en caso de que lancemos go test -v (verbose
// mode)

// Si queremos que el test falle usaremos t.Fail en el caso de que queramos que
// continue con la ejecución de los tests que falten, o t.FailNow si queremos
// que falle y se detenga la ejecución de los tests

// Y después de ver esto, tenemos la combinación de ambos:
// Error = Log + Fail
// Errorf = Logf + Fail
// Fatal = Log + FailNow
// Fatalf = Logf + FailNow

// Cuando usar uno u otro depende de la información que queramos obtener de los
// tests, si creemos que aunque un test falle podemos seguir obteniendo
// información, entonces usaremos Error o Errorf, en caso de que pensemos que no
// vamos a obtener nada útil si uno de los test falla, usaremos Fatal o Fatalf
