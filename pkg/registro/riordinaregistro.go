package registro

import "sort"

func (re *Registro) riordinaRegistro() (err error) {
	sort.Sort(re.ListaMovimenti)
	// primo movimento registro deve avere numero 1
	// quando numero anno diverso da quello precedente numero torna 1
	re.ListaMovimenti[0].Numero = 1
	first := re.ListaMovimenti[0]
	var numero int
	for i, v := range re.ListaMovimenti {
		if i == 0 {
			numero = 1
		} else if i > 0 {
			if v.DataMovimento.Year() != first.DataMovimento.Year() {
				numero = 1
				re.ListaMovimenti[i].Numero = numero
				first = v
			} else {
				numero++
				re.ListaMovimenti[i].Numero = numero
			}
		}

	}
	return
}
