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
		panic(err)
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
		panic(err)
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

func TestModifyMovimentoToRegistro(t *testing.T) {
	// ci aggiungo due movimenti
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
		TipoMovimento:   "carico",
	}
	// istanzio un registro con già dei movimenti
	testRegistro := Registro{
		ID:             1,
		ListaMovimenti: ListaMovimenti{movimento1, movimento2},
	}
	// creo nuovo movimento che scambierò con il movimento 1
	movimento3 := Movimento{
		Numero:        3,
		DataMovimento: time.Date(2019, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:      800,
		TipoMovimento: "scarico",
	}
	// attuo modifica
	err := testRegistro.ModifyMovimentoToRegistro(movimento1.ID, &movimento3)
	// controllo che non abbia restituito errore
	if err != nil {
		panic(err)
	}
	// mi aspetto che il movimento con id 1 sia cambiato e che sia uguale a movimento3
	for _, v := range testRegistro.ListaMovimenti {
		if v.ID == movimento1.ID {
			if v.TipoMovimento != movimento3.TipoMovimento && v.Quantita != movimento3.Quantita && v.DataMovimento != movimento3.DataMovimento {
				t.Error("I movimenti sono diversi")
			}
		}
	}
	// controllo che il movimento modificato sia il primo in ordine nel registro dato che data movimento 2019
	if testRegistro.ListaMovimenti[0].DataMovimento != movimento3.DataMovimento {
		t.Error("Le date sono diverse quindi registro non oridinato")
	}
}

func TestEliminareMovimentoToRegistro(t *testing.T) {
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
	// istanzio un registro con già dei movimenti
	testRegistro := Registro{
		ID:             1,
		ListaMovimenti: ListaMovimenti{movimento1, movimento2, movimento3},
	}
	// elimino dal registro il movimento1
	testRegistro.EliminareMovimentoToRegistro(movimento1.ID)
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

func TestTrovaIndiceConId(t *testing.T) {
	var sliceMovimento []Movimento
	movimento1 := Movimento{
		ID:              1,
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	movimento2 := Movimento{
		ID:              24,
		Numero:          3,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "carico",
	}
	sliceMovimento = append(sliceMovimento, movimento1, movimento2)
	// cerco lista id presente nella lista
	result := trovaIndice(24, sliceMovimento)
	if result == -1 {
		t.Error("Il movimento con l'id inserito non è presente nella lista")
	}
}

func TestTrovaIndicesenzaId(t *testing.T) {
	var sliceMovimento []Movimento
	movimento1 := Movimento{
		ID:              1,
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	movimento2 := Movimento{
		ID:              24,
		Numero:          3,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "carico",
	}
	sliceMovimento = append(sliceMovimento, movimento1, movimento2)
	// cerco lista id presente nella lista
	result := trovaIndice(25, sliceMovimento)
	if result != -1 {
		t.Error("Il movimento con l'id inserito è presente nella lista")
	}
}
