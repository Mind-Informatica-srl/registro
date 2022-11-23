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
	if len(re.ListaMovimenti) == 0 && movimento.TipoMovimento != "carico" {
		err = errors.New("il primo elemento dell'array non può essere uno scarico")
		re.counter = 0
		return err

	}
	movimento.ID = re.counter
	movimento.DataInserimento = time.Now()
	// appendo il nuovo movimento alla lista dei movimenti del registro
	re.ListaMovimenti = append(re.ListaMovimenti, *movimento)
	// dato che nuovo elemento è stato inserito riordino lista elementi del registro
	err = re.riordinaRegistro()
	if err != nil {
		return
	}
	if err = re.CheckRegistro(); err != nil {
		if err = re.EliminareMovimentoToRegistro(movimento.ID); err != nil {
			return err
		}
		re.counter--
		stringaerrore := fmt.Sprintf("Movimento con id %d non è stato possibile inserirlo dato che la quantita del saldo è minore di quella dello scarico", movimento.ID)
		err = errors.New(stringaerrore)
		return err
	}
	return
}
