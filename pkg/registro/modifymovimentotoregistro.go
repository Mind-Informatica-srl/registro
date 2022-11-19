package registro

// ModifyMovimentoToRegistro modifica movimento del registro nella lista dei movimenti e riordina i movimenti
func (re *Registro) ModifyMovimentoToRegistro(id int, movimento *Movimento) (err error) {
	// cerco l'indice del movimento con id corrispondente a quello passato
	indice := trovaIndice(id, re.ListaMovimenti)
	// copio nel movimento con indice trovato i campi tipo,datamovimento,quantita, dal movimento passato come argomento
	re.ListaMovimenti[indice].TipoMovimento = movimento.TipoMovimento
	re.ListaMovimenti[indice].DataMovimento = movimento.DataMovimento
	re.ListaMovimenti[indice].Quantita = movimento.Quantita
	// dato che elemento modificato è stato inserito riordino lista elementi del registro
	err = re.riordinaRegistro()
	return
}
