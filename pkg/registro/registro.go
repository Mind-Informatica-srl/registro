package registro

import (
	"errors"
)

type Registro struct {
	ID             int
	ListaMovimenti []Movimento
	counter        int
	countercsv     int
	sorter         func(Movimento, Movimento) bool
}

func (registro *Registro) Len() int { return len(registro.ListaMovimenti) }

func (registro *Registro) Less(i, j int) bool {
	// se sorter è definito lo uso
	if registro.sorter != nil {
		movimento1 := registro.ListaMovimenti[i]
		movimento2 := registro.ListaMovimenti[j]
		return registro.sorter(movimento1, movimento2)
	}
	// altrimenti uso l'ordinamento di default
	// se dataMovimento di i è uguale a data movimento di j
	if registro.ListaMovimenti[i].DataMovimento.Equal(registro.ListaMovimenti[j].DataMovimento) {
		// se tipo di i è uguale a tipo di j
		if registro.ListaMovimenti[i].TipoMovimento == registro.ListaMovimenti[j].TipoMovimento {
			// restituisco true se data inserimento di i è precedente a data inserimento j
			if registro.ListaMovimenti[i].DataInserimento.Before(registro.ListaMovimenti[j].DataInserimento) {
				return true
			} else {
				// altrimenti restituisco false
				return false
			}

		} else { // altrimenti
			// restituisco true se tipo di i viene prima in ordine alfabetico di j
			if registro.ListaMovimenti[i].TipoMovimento < registro.ListaMovimenti[j].TipoMovimento {
				return true
			} else {
				//altrimenti restituisco false}
				return false
			}

		}
	} else {
		// restituisco true se data movimento di i è precedente a data movimento di j
		if registro.ListaMovimenti[i].DataMovimento.Before(registro.ListaMovimenti[j].DataMovimento) {
			return true
		} else {
			// altrimenti restituisco false
			return false
		}

	}
}

func (registro *Registro) Swap(i, j int) {
	registro.ListaMovimenti[i], registro.ListaMovimenti[j] = registro.ListaMovimenti[j], registro.ListaMovimenti[i]
}

// Calcola Saldo
func (registro Registro) CalcolaSaldo(movimento Movimento) (int, error) {
	saldo := 0
	codiceCerMovimento := movimento.CodiceCer
	// parto da 0 per ogni movimento fino al movimneto -1
	for _, v := range registro.ListaMovimenti {
		if v.CodiceCer == codiceCerMovimento {
			if v.TipoMovimento == "carico" {
				saldo += v.Quantita
			} else {
				if saldo < v.Quantita {
					return 0, errors.New("Nuovo errore relativo al saldo")
				} else {
					saldo -= v.Quantita
				}
			}
		}
	}
	return saldo, nil
}
