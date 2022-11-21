package registro

import (
	"testing"
	"time"
)

func TestRiordinaRegistro(t *testing.T) {
	// ci aggiungo tre movimenti
	movimento1 := Movimento{
		ID:              1,
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	movimento2 := Movimento{
		ID:              2,
		Numero:          3,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "carico",
	}
	movimento3 := Movimento{
		ID:              3,
		Numero:          4,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        1000,
		TipoMovimento:   "scarico",
	}
	// istanzio un registro con i tre movimenti
	testRegistro := Registro{
		ID:             1,
		ListaMovimenti: []Movimento{movimento1, movimento2, movimento3},
	}
	// se si verifica errore, loggare con fallimento l'errore
	err := testRegistro.riordinaRegistro()
	if err != nil {
		t.Error(err)
	}
	// verifico che il primo movimento della lista abbia lo stesso id del secondo movimento che mi aspetto dopo il riordine e che il suo Numero sia uguale a 1 dato che primo movimento dell'anno
	if testRegistro.ListaMovimenti[0].ID != movimento2.ID && testRegistro.ListaMovimenti[0].Numero == 1 {
		t.Error("Movimenti non corretti")
	}
	// verifico che il secondo movimento della lista abbia lo stesso id del terzo movimento che mi aspetto dopo il riordine e che il suo Numero sia uguale a 2 dato che secondo movimento dell'anno
	if testRegistro.ListaMovimenti[1].ID != movimento3.ID && testRegistro.ListaMovimenti[1].Numero == 2 {
		t.Error("Movimenti non corretti")
	}
	// verifico che il terzo movimento della lista abbia lo stesso id del primo movimento che mi aspetto dopo il riordine e che il suo Numero sia uguale a 1 dato che è primo elemento del 'nuovo anno'
	if testRegistro.ListaMovimenti[2].ID != movimento1.ID && testRegistro.ListaMovimenti[2].Numero == 1 {
		t.Error("")
	}
}
