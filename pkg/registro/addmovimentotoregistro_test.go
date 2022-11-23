package registro

import (
	"testing"
	"time"
)

func TestAddMovimentoToRegistro(t *testing.T) {
	// istanzio un registro
	testRegistro := Registro{
		ID:             1,
		ListaMovimenti: []Movimento{},
	}

	// ci aggiungo un movimento
	movimento1 := Movimento{
		Numero:        2,
		DataMovimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:      400,
		TipoMovimento: "scarico",
		CodiceCer:     "123456",
	}
	movimento2 := Movimento{
		Numero:        2,
		DataMovimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:      600,
		TipoMovimento: "scarico",
		CodiceCer:     "123456",
	}
	movimento3 := Movimento{
		Numero:        3,
		DataMovimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:      600,
		TipoMovimento: "scarico",
		CodiceCer:     "123456",
	}
	// mi aspetto che lista di movimenti sia vuota dato che primo movimento che volevo aggiungere alla lista dei movimenti era uno scarico e invece dovrebbe essere un carico
	if testRegistro.AddMovimentoToRegistro(&movimento1).Error() != "Il primo elemento dell'array non può essere uno scarico" {
		t.Error("Mi aspetto che lista dei movimenti sia vuota dato che è stato inserito uno scarico come primo elemenento della lista dei moviemnti")
	}
	movimento1.TipoMovimento = "carico"
	if err := testRegistro.AddMovimentoToRegistro(&movimento1); err != nil {
		t.Error(err)
	}
	// Mi aspetto che sia presente un errore dato che ho inserito un movimento scarico come primo dell lista
	if testRegistro.AddMovimentoToRegistro(&movimento2).Error() != "Movimento con id 2 non è stato possibile inserirlo dato che la quantita del saldo è minore di quella dello scarico" {
		t.Error("Mi aspetto che ci sia un errore dato che il movimento che voglio inserire è il primo nella lista ed è uno scarico")
	}

	// Ultimo test che verifico mi dià errore e se aggiungendo uno scarico maggiore della quantità del saldo attuale mi dia errore
	if testRegistro.AddMovimentoToRegistro(&movimento3).Error() != "Movimento con id 2 non è stato possibile inserirlo dato che la quantita del saldo è minore di quella dello scarico" {
		t.Error("Mi aspetto che ci sia un errore dato lo scarico è maggiore del saldo")
	}

}
