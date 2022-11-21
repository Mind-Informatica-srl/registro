package registro

// CreaNuovoRegistro crea un nuovo registro con l'id e la funzione di ordinamento.
// Se sorter è nil il registro utilizza l'ordinamento di default (data movimento, tipo, data inserimento)
func CreaNuovoRegistro(id int, sorter func(mov1 Movimento, mov2 Movimento) bool) Registro {
	return Registro{
		ID:     id,
		sorter: sorter,
	}
}
