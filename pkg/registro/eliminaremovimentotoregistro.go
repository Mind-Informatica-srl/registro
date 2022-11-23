package registro

// EliminareMovimentoToRegistro elimina movimento dal registro e riodina i movimenti
func (re *Registro) EliminareMovimentoToRegistro(id int) (err error) {
	// cerco l'indice del movimento con id corrispondente a quello passato
	indice := trovaIndice(id, re.ListaMovimenti)

	// prendo movimento
	movimentoDaEliminare := re.ListaMovimenti[indice]
	// elimino dalla lista dei movimenti il movimento specificato
	re.ListaMovimenti = append(re.ListaMovimenti[:indice], re.ListaMovimenti[indice+1:]...)
	if err = re.CheckRegistro(); err != nil {
		re.ListaMovimenti = append(re.ListaMovimenti, movimentoDaEliminare)
		re.riordinaRegistro()
	}
	// dato che elemento eliminato riordino lista movimenti del registro
	err = re.riordinaRegistro()
	if err != nil {
		panic(err)
	}

	return
}
