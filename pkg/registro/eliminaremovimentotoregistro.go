package registro

import (
	"errors"
	"fmt"
)

// EliminareMovimentoToRegistro elimina movimento dal registro e riodina i movimenti
func (re *Registro) EliminareMovimentoToRegistro(id int) (err error) {
	// cerco l'indice del movimento con id corrispondente a quello passato
	indice := trovaIndice(id, re.ListaMovimenti)

	// prendo codiceCer elemento
	movimentodaelimare := re.ListaMovimenti[indice]
	// elimino dalla lista dei movimenti il movimento specificato
	re.ListaMovimenti = append(re.ListaMovimenti[:indice], re.ListaMovimenti[indice+1:]...)
	// dato che elemento eliminato riordino lista movimenti del registro
	err = re.riordinaRegistro()
	saldo, errore := re.CalcolaSaldo(movimentodaelimare)
	err = errore
	if err != nil {
		re.ListaMovimenti = append(re.ListaMovimenti, movimentodaelimare)
		stringaerrore := fmt.Sprintf("Movimento %d non è stato possibile eliminarlo dato che la quantita del saldo è minore di quella dello scarico", movimentodaelimare.ID)
		panic(errors.New(stringaerrore))
	}
	fmt.Println(saldo)
	return
}

// elimino elemto verifico che scarico non scoppi se scoppia panic error non puoi eliminare questo movimento senno scoppia scarico
