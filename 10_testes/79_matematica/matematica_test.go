package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

/*
	Conventions:
	The TEST file must be endend with _test.go (e.g. foo.go --> foo_test.go)
	The TEST function must be started with Test (e.g. func Bar() --> func TestBar())
	The arguments from TEST function must be t *testing.T (e.g. func Bar() --> func TestBar(t *testing.T))
*/

// It must be ended with _test.go. To run this test, you need to click on "run test" or using command "go test" inside the dir or "go test ./..." outside the dir
func TestMedia(t *testing.T) {
	t.Parallel() // allow this test run in parallel with others tests
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
