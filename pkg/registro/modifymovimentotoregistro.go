package registro

import (
	"fmt"
)

// ModifyMovimentoToRegistro modifica movimento del registro nella lista dei movimenti e riordina i movimenti
func (re *Registro) ModifyMovimentoToRegistro(id int, movimento *Movimento) (err error) {
	// cerco l'indice del movimento con id corrispondente a quello passato
	indice := trovaIndice(id, re.ListaMovimenti)
	// copio nel movimento con indice trovato i campi tipo,datamovimento,quantita, dal movimento passato come argomento
	var oldMovimento Movimento = re.ListaMovimenti[indice]
	re.ListaMovimenti[indice].TipoMovimento = movimento.TipoMovimento
	re.ListaMovimenti[indice].DataMovimento = movimento.DataMovimento
	re.ListaMovimenti[indice].Quantita = movimento.Quantita
	var saldo int
	if saldo, err = re.CalcolaSaldo(re.ListaMovimenti[indice]); err != nil {
		re.ListaMovimenti[indice].Quantita = oldMovimento.Quantita
		re.ListaMovimenti[indice].TipoMovimento = oldMovimento.TipoMovimento
		re.ListaMovimenti[indice].DataMovimento = oldMovimento.DataMovimento
		// se rifaccio calcola saldo e mi torna errore, rimettendo vecchio movimento significa che è presente un errore
		if saldo, err = re.CalcolaSaldo(re.ListaMovimenti[indice]); err != nil {
			panic(err)
		}
	}
	fmt.Println(saldo)
	// dato che elemento modificato è stato inserito riordino lista elementi del registro
	err = re.riordinaRegistro()
	if err != nil {
		panic(err)
	}
	return
}
