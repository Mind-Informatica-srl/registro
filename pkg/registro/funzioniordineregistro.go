package registro

// OrdinaPerDataMovimento ritorna true se la la data del movimento 1 viene prima della data del movimento 2 se sono uguali invece i movimenti si testano per la data di inserimento dei movimenti che sarà sempre diversa
func OrdinaPerDataMovimento(mov1 Movimento, mov2 Movimento) bool {
	if mov1.DataMovimento.Before(mov2.DataMovimento) {
		return true
	} else if mov1.DataMovimento.Equal(mov2.DataMovimento) {
		return OrdinaPerDataInserimento(mov1, mov2)
	} else {
		return false
	}
}

// OrdinaPerDataInserimento ritorna true se la la data di inserimento 1 viene prima della data di inserimento 2 se sono uguali invece i movimenti si testano per la data di inserimento dei movimenti che sarà sempre diversa
func OrdinaPerDataInserimento(mov1 Movimento, mov2 Movimento) bool {
	return mov1.DataInserimento.Before(mov2.DataInserimento)
}

// OrdinaPerTipoMovimento ritorna true se il tipo di movimento 1 viene alfabeticamente prima del tipo di movimento 2 se sono uguali invece i movimenti si testano per la data di inserimento dei movimenti che sarà sempre diversa
func OrdinaPerTipoMovimento(mov1 Movimento, mov2 Movimento) bool {
	if mov1.TipoMovimento < mov2.TipoMovimento {
		return true
	} else if mov1.TipoMovimento == mov2.TipoMovimento {
		return OrdinaPerDataInserimento(mov1, mov2)
	} else {
		return false
	}
}
