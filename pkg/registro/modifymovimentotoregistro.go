package registro

import (
	"errors"
	"fmt"
)

// ModifyMovimentoToRegistro modifica movimento del registro nella lista dei movimenti e riordina i movimenti
func (re *Registro) ModifyMovimentoToRegistro(id int, movimento *Movimento) (err error) {
	// cerco l'indice del movimento con id corrispondente a quello passato
	indice := trovaIndice(id, re.ListaMovimenti)
	// copio nel movimento con indice trovato i campi tipo,datamovimento,quantita, dal movimento passato come argomento
	var oldMovimento *Movimento = &re.ListaMovimenti[indice]
	re.ListaMovimenti[indice].TipoMovimento = movimento.TipoMovimento
	re.ListaMovimenti[indice].DataMovimento = movimento.DataMovimento
	re.ListaMovimenti[indice].Quantita = movimento.Quantita
	var newMovimento *Movimento = &re.ListaMovimenti[indice]
	// dato che elemento modificato è stato inserito riordino lista elementi del registro
	err = re.riordinaRegistro()
	saldo, errore := re.CalcolaSaldo(*movimento)
	err = errore
	if err != nil {
		re.EliminareMovimentoToRegistro(newMovimento.ID)
		re.AddMovimentoToRegistro(oldMovimento)
		stringaerrore := fmt.Sprintf("Movimento %d non è stato possibile modificarlo dato che la quantita del saldo è minore di quella dello scarico", movimento.ID)
		panic(errors.New(stringaerrore))
	}
	fmt.Println(saldo)
	return
}
