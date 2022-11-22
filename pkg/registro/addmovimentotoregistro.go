package registro

import (
	"errors"
	"fmt"
	"time"
)

// AddMovimentoToRegistro aggiunge un movimento al registro, gli assegna un id univoco e la data di inserimento attuale e rioridina i movimenti
func (re *Registro) AddMovimentoToRegistro(movimento *Movimento) (err error) {
	// assegno l'id al movimento
	re.counter++
	movimento.ID = re.counter
	movimento.DataInserimento = time.Now()
	if len(re.ListaMovimenti) == 0 {
		if movimento.TipoMovimento != "carico" {
			err = errors.New("Il primo elemento dell'array non può essere uno scarico")
			panic(err)
		}

	}
	// appendo il nuovo movimento alla lista dei movimenti del registro
	re.ListaMovimenti = append(re.ListaMovimenti, *movimento)
	// dato che nuovo elemento è stato inserito riordino lista elementi del registro
	err = re.riordinaRegistro()
	saldo, errore := re.CalcolaSaldo(*movimento)
	err = errore
	if err != nil {
		re.EliminareMovimentoToRegistro(movimento.ID)
		stringaerrore := fmt.Sprintf("Movimento %d non è stato possibile inserirlo dato che la quantita del saldo è minore di quella dello scarico", movimento.ID)
		panic(errors.New(stringaerrore))
	}
	fmt.Println(saldo)
	return
}
