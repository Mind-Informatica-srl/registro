package registro

import (
	"testing"
	"time"
)

func TestEliminareMovimentoToRegistro(t *testing.T) {
	// creo tre instanze movimento
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
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "scarico",
	}
	movimento3 := Movimento{
		ID:              3,
		Numero:          4,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "carico",
	}
	// istanzio un registro con i tre movimenti instanziati precedentemente
	testRegistro := Registro{
		ID:             1,
		ListaMovimenti: []Movimento{movimento1, movimento2, movimento3},
	}
	// elimino dal registro il movimento1
	if err := testRegistro.EliminareMovimentoToRegistro(movimento1.ID); err != nil {
		return
	}
	// verifico che nella lista dei movimenti non ci sia più movimento
	for _, elemento := range testRegistro.ListaMovimenti {
		if elemento.ID == movimento1.ID {
			t.Error("Elemento presente ancora nel db")
		}
	}
	// verifico riordinamento e che quindi primo elemento della lista movimenti sia movimento 3
	if testRegistro.ListaMovimenti[0].ID != movimento3.ID {
		t.Error("I movimenti sono diversi")
	}

}
