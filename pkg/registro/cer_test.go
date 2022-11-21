package registro

import (
	"testing"
)

func TestCerValidator(t *testing.T) {
	var testCer Cer
	testCer = "123456"
	// mi aspetto che il risultato di questo cer validator sia nil dato che il cer ha 6 numeri
	if CerValidator(testCer) != nil {
		t.Errorf("Atteso cer con 6 numeri e invece non ce ne sono 6")
	}
	testCer = "12345shi"
	Expected := "Codice cer non ha 6 cifre"
	// mi aspetto che si verifichi errore dato che il cer non contiene 6 numeri
	if CerValidator(testCer).Error() != Expected {
		t.Errorf("Atteso errore che non è avvenuto")
	}
}
