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
	saldo, errore := re.CalcolaSaldo(movimento.CodiceCer, *movimento)
	err = errore
	if err != nil {
		panic(err)
	}
	fmt.Println(saldo)
	// appendo il nuovo movimento alla lista dei movimenti del registro
	re.ListaMovimenti = append(re.ListaMovimenti, *movimento)

	// dato che nuovo elemento è stato inserito riordino lista elementi del registro
	err = re.riordinaRegistro()
	return
}
