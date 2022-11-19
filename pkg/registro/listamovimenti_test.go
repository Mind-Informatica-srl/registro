package registro

import (
	"testing"
	"time"
)

func TestLenListaMovimenti(t *testing.T) {
	// creo 2 instanze di 2 movimenti
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
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	var listaMovimenti ListaMovimenti
	// aggiungo/appendo all'instanza di listaMovimenti i movimenti precedentemente inseriti
	listaMovimenti = append(listaMovimenti, movimento1, movimento2)
	// mi aspetto che la lista movimenti sia uguale a 2 dato che ho inserito 2 elementi
	if listaMovimenti.Len() != 2 {
		t.Errorf("Valore attesso della lunghezza della lista %d but got %d", 2, listaMovimenti.Len())
	}
}

func TestLessListaMovimenti(t *testing.T) {
	// creo 2 instanze di 2 movimenti
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
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	var listaMovimenti ListaMovimenti
	listaMovimenti = append(listaMovimenti, movimento1, movimento2)
	// Scambio primo elemento della lista con secondo elemento della lista
	listaMovimenti.Swap(0, 1)
	// mi aspetto che il secondo movimento dopo lo swap sia il primo della lista dato che l'ho inserito per ultimo
	checkSecondoMovimentoPrimoLista := listaMovimenti[0].ID == movimento2.ID
	if checkSecondoMovimentoPrimoLista != true {
		t.Errorf("Attesso movimento 2 come primo della lista ma ho %v", listaMovimenti[0])
	}
}

func TestSwapListaMovimenti(t *testing.T) {
	// creo 2 instanze di 2 movimenti per effettuare il primo test
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
		Numero:          2,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	movimento3 := Movimento{
		ID:              3,
		Numero:          4,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local), // stessa DataMovimento del movimento4
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	movimento4 := Movimento{
		ID:              4,
		Numero:          5,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local), // stessa DataMovimento del movimento3
		DataInserimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	movimento5 := Movimento{
		ID:              5,
		Numero:          6,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local), // stessa DataMovimento del movimento4
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "scarico",
	}
	movimento6 := Movimento{
		ID:              6,
		Numero:          7,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local), // stessa DataMovimento del movimento3
		DataInserimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	// creo instanza lista movimenti
	var listaMovimenti ListaMovimenti
	listaMovimenti = append(listaMovimenti, movimento1, movimento2, movimento3, movimento4, movimento5, movimento6)
	// testo tramite data movimento
	// mi aspetto che il primo movimento della lista(index=0) abbia data di movimento che sia dopo di quella del secondo movimento e quindi ritorni false
	if listaMovimenti.Less(0, 1) == true {
		t.Errorf("Data di movimento del secondo elemento della lista viene dopo di quella del primo movimento della lista ")
	}
	// testo tramite data inserimento
	// adesso mi aspetto che il terzo movimento della lista(index=2) abbia data di inserimento che dopo di quella del movimento e che quinndi ritorni false
	if listaMovimenti.Less(2, 3) == true {
		t.Error("Data di inserimento del quarto elemento della lista viene dopo di quella del terzo movimento della lista")
	}
	// testo tramite  tipoMovimento del Movimento
	// mi aspetto che alfabeticamente lo scarico del quinto elemento della lista(index=4) venga alfabeticamente dopo il carico del sesto elemento della lista
	if listaMovimenti.Less(4, 5) == true {
		t.Error("Il tipo di movimento scarico del sesto elemento della lista è venuto alfabeticamente prima del quinto elemento della lista chè un tipo movimento carico")
	}
}
