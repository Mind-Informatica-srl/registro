package registro

import (
	"testing"
	"time"
)

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
