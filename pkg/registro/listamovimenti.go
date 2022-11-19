package registro

type ListaMovimenti []Movimento

func (listaMovimenti ListaMovimenti) Len() int { return len(listaMovimenti) }
func (listaMovimenti ListaMovimenti) Less(i, j int) bool {
	// se dataMovimento di i è uguale a data movimento di j
	if listaMovimenti[i].DataMovimento.Equal(listaMovimenti[j].DataMovimento) {
		// se tipo di i è uguale a tipo di j
		if listaMovimenti[i].TipoMovimento == listaMovimenti[j].TipoMovimento {
			// restituisco true se data inserimento di i è precedente a data inserimento j
			if listaMovimenti[i].DataInserimento.Before(listaMovimenti[j].DataInserimento) {
				return true
			} else {
				// altrimenti restituisco false
				return false
			}

		} else { // altrimenti
			// restituisco true se tipo di i viene prima in ordine alfabetico di j
			if listaMovimenti[i].TipoMovimento < listaMovimenti[j].TipoMovimento {
				return true
			} else {
				//altrimenti restituisco false}
				return false
			}

		}
	} else {
		// restituisco true se data movimento di i è precedente a data movimento di j
		if listaMovimenti[i].DataMovimento.Before(listaMovimenti[j].DataMovimento) {
			return true
		} else {
			// altrimenti restituisco false
			return false
		}

	}
}
func (listaMovimenti ListaMovimenti) Swap(i, j int) {
	listaMovimenti[i], listaMovimenti[j] = listaMovimenti[j], listaMovimenti[i]
}
