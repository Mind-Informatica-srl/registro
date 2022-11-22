package registro

import (
	"testing"
	"time"
)

func TestModifyMovimentoToRegistro(t *testing.T) {
	// ci aggiungo due movimenti
	movimento1 := Movimento{
		ID:              1,
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        1500,
		TipoMovimento:   "carico",
		CodiceCer:       "123456",
	}
	movimento2 := Movimento{
		ID:              2,
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        1500,
		TipoMovimento:   "carico",
		CodiceCer:       "123456",
	}
	// istanzio un registro con già dei movimenti
	testRegistro := Registro{
		ID: 1,
	}
	testRegistro.AddMovimentoToRegistro(&movimento1)
	testRegistro.AddMovimentoToRegistro(&movimento2)
	// creo nuovo movimento che scambierò con il movimento 1
	movimento3 := Movimento{
		Numero:        3,
		DataMovimento: time.Date(2019, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:      3000,
		TipoMovimento: "scarico",
		CodiceCer:     "123456",
	}
	err := testRegistro.ModifyMovimentoToRegistro(2, &movimento3)
	if err != nil {
		t.Error(err)
	}
	// mi aspetto che nella posizione 1 della lista elementi i dati del movimento siano tornati alla versione di partenza dato che aggiungendo lo scarico il saldo tornerebbe inferiore allo scarico che volevamo inserire
	if testRegistro.ListaMovimenti[1].Quantita != movimento2.Quantita && testRegistro.ListaMovimenti[1].TipoMovimento != movimento2.TipoMovimento && testRegistro.ListaMovimenti[1].DataMovimento != movimento2.DataMovimento {
		t.Error("")
	}
}
