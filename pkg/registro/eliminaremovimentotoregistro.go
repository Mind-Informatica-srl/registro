package registro

import (
	"fmt"
)

// EliminareMovimentoToRegistro elimina movimento dal registro e riodina i movimenti
func (re *Registro) EliminareMovimentoToRegistro(id int) (err error) {
	// cerco l'indice del movimento con id corrispondente a quello passato
	indice := trovaIndice(id, re.ListaMovimenti)

	// prendo movimento
	movimentoDaEliminare := re.ListaMovimenti[indice]
	// elimino dalla lista dei movimenti il movimento specificato
	re.ListaMovimenti = append(re.ListaMovimenti[:indice], re.ListaMovimenti[indice+1:]...)
	var saldo int
	if saldo, err = re.CalcolaSaldo(movimentoDaEliminare); err != nil {
		re.ListaMovimenti = append(re.ListaMovimenti, movimentoDaEliminare)
		re.riordinaRegistro()
		// se rifaccio calcola saldo e mi torna errore, rimettendo vecchio movimento significa che è presente un errore
		if saldo, err = re.CalcolaSaldo(movimentoDaEliminare); err != nil {
			panic(err)
		}
	}
	fmt.Println(saldo)
	// dato che elemento eliminato riordino lista movimenti del registro
	err = re.riordinaRegistro()
	if err != nil {
		panic(err)
	}

	return
}
