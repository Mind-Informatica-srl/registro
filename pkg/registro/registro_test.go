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
		t.Error(err)
	}
	// mi aspetto che il movimento con id 1 sia cambiato e che sia uguale a movimento3
	for _, v := range testRegistro.ListaMovimenti {
		if v.ID == movimento1.ID {
			if v.TipoMovimento != movimento3.TipoMovimento && v.Quantita != movimento3.Quantita && v.DataMovimento != movimento3.DataMovimento {
				t.Error("I movimenti sono diversi")
			}
		}
	}
	// controllo che il primo in ordine nel registro sia il movimento modificato dato che data movimento 2019
	if testRegistro.ListaMovimenti[0].ID != movimento1.ID {
		t.Error("Le date sono diverse quindi registro non oridinato")
	}
}

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
		ListaMovimenti: ListaMovimenti{movimento1, movimento2, movimento3},
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

func TestTrovaIndiceConId(t *testing.T) {
	// creo uno slice di Movimento
	var sliceMovimento []Movimento
	// creo due instanze di movimento
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
	// appendo/aggiungo i due movimenti precedentemente creati allo sliceMovimento
	sliceMovimento = append(sliceMovimento, movimento1, movimento2)
	// cerco indice relativo al movimento presente nella lista
	result := trovaIndice(24, sliceMovimento)
	// in questo caso mi aspetto di trovare index = 1 dato che ho ricercato l'id 24 ovvero quello del secondo movimento che è stato appesso allo slice
	if result != 1 {
		t.Error("Il movimento con l'id inserito non è presente nella lista")
	}
}

func TestTrovaIndicesenzaId(t *testing.T) {
	// creo uno slice di Movimento
	var sliceMovimento []Movimento
	// creo due instanze di movimento
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
	// appendo/aggiungo i due movimenti precedentemente creati allo sliceMovimento
	sliceMovimento = append(sliceMovimento, movimento1, movimento2)
	// cerco indice relativo al movimento presente nella lista che in questo caso non c'è nessun movimento con id 25
	result := trovaIndice(25, sliceMovimento)
	// il risultato che prevediamo che ritorni è -1 dato che abbiamo ricercato un id di un elemento non presente nello slice
	if result != -1 {
		t.Error("Il movimento con l'id inserito è presente nella lista")
	}
}

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
		ListaMovimenti: ListaMovimenti{movimento1, movimento2, movimento3},
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
