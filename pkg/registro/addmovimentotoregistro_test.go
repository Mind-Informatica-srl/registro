package registro

import (
	"testing"
	"time"
)

func TestAddMovimentoToRegistro(t *testing.T) {
	// istanzio un registro
	testRegistro := Registro{
		ID:             1,
		ListaMovimenti: ListaMovimenti{},
	}

	// ci aggiungo un movimento
	movimento1 := Movimento{
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	err := testRegistro.AddMovimentoToRegistro(&movimento1)
	// controllo che non abbia restituito errore
	if err != nil {
		t.Error(err)
	}
	// ci aggiungo un movimento con data precedente al movimento precedente
	movimento2 := Movimento{
		Numero:          3,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "carico",
	}
	err = testRegistro.AddMovimentoToRegistro(&movimento2)
	// controllo che non abbia restituito errore
	if err != nil {
		t.Error(err)
	}
	// confronto gli id dei due movimenti per vedere che siano diversi
	isDifferent := testRegistro.ListaMovimenti[0].ID != testRegistro.ListaMovimenti[1].ID
	if isDifferent != true {
		t.Error("Expected", true, "Got", false)
	}
	// controllo che il secondo movimento inserito sia il primo in ordine nel registro
	checkSecondoMovimentoPrimoLista := testRegistro.ListaMovimenti[0].ID == movimento2.ID // se uguale significa che è lo stesso movimento
	if checkSecondoMovimentoPrimoLista != true {
		t.Error("Expected that movimento2 equal to listaMovimenti[0]")
	}
}
