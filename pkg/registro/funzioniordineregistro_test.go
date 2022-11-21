package registro

import (
	"testing"
	"time"
)

func TestOrdinaPerDataMovimento(t *testing.T) {
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
		DataInserimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "carico",
	}
	// mi aspetto che la funzione restitusca false dato che DataMovimento del movimento 2 è più piccola DataMovimento del movimento 1
	if OrdinaPerDataMovimento(movimento1, movimento2) != false {
		t.Error("Attesso movimento 2 con data di movimento che viene prima di quella di movimento 1")
	}
	// testo ordinamento se data movimento del movimento 1 è  uguale a quella del movimento 2
	movimento2.DataMovimento = time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	// quindi si ordina per data inserimento
	// mi aspetto che la funzione restitusca false dato che data inserimento movimento 2 è più piccola di data di inserimento del movimento 1
	if OrdinaPerDataMovimento(movimento1, movimento2) != false {
		t.Error("Attesso movimento 2 con data di movimento che viene prima di quella di movimento 1 dato che data di inserimento viene prima di quella del movimento 1")
	}
}

func TestOrdinaPerDataInserimento(t *testing.T) {
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
		DataInserimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "carico",
	}
	// mi aspetto che la funzione restitusca false dato che data inserimento movimento 2 è più piccola di data di inserimento del movimento 1
	if OrdinaPerDataMovimento(movimento1, movimento2) != false {
		t.Error("Attesso movimento 2 con data di movimento che viene prima di quella di movimento 1")
	}
}

func TestOrdinaPerTipoMovimento(t *testing.T) {
	// ci aggiungo due movimenti
	movimento1 := Movimento{
		ID:              1,
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "scarico",
	}
	movimento2 := Movimento{
		ID:              2,
		Numero:          3,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "carico",
	}
	// mi aspetto che la funzione restitusca false dato che tipo movimento  "carico" del movimento 2 viene alfabeticamente di tipo movimento "scarico" del movimento 1
	if OrdinaPerDataMovimento(movimento1, movimento2) != false {
		t.Error("Attesso movimento 2 con tipo di movimento che viene prima di quello di movimento 1")
	}
	// testo tipoMovimento se tipo del movimento 1 è  tipo del movimento 2
	// cambio tipo movimento 2 in scarico
	movimento2.TipoMovimento = "scarico"
	// quindi si ordina per data inserimento
	// mi aspetto che la funzione restitusca false dato che data inserimento movimento 2 è più piccola di data di inserimento del movimento 1
	if OrdinaPerDataMovimento(movimento1, movimento2) != false {
		t.Error("Attesso movimento 2 con tipo di movimento che viene prima di quello di movimento 1")
	}
}
