package registro

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type Registro struct {
	ID             int
	ListaMovimenti ListaMovimenti
	counter        int
	countercsv     int
}

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

// AddMovimentoToRegistro aggiunge un movimento al registro, gli assegna un id univoco e la data di inserimento attuale e rioridina i movimenti
func (re *Registro) AddMovimentoToRegistro(movimento *Movimento) (err error) {
	// assegno l'id al movimento
	re.counter++
	movimento.ID = re.counter
	movimento.DataInserimento = time.Now()
	// appendo il nuovo movimento alla lista dei movimenti del registro
	re.ListaMovimenti = append(re.ListaMovimenti, *movimento)
	// dato che nuovo elemento è stato inserito riordino lista elementi del registro
	err = re.riordinaRegistro()
	return
}

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

// EliminareMovimentoToRegistro elimina movimento dal registro e riodina i movimenti
func (re *Registro) EliminareMovimentoToRegistro(id int) (err error) {
	// cerco l'indice del movimento con id corrispondente a quello passato
	indice := trovaIndice(id, re.ListaMovimenti)
	// elimino dalla lista dei movimenti il movimento specificato
	re.ListaMovimenti = append(re.ListaMovimenti[:indice], re.ListaMovimenti[indice+1:]...)
	// dato che elemento eliminato riordino lista movimenti del registro
	err = re.riordinaRegistro()
	return
}

// CreaCsvFile crea un file csv con la lista dei movimenti del registro e restituisce il percorso del file
func (re *Registro) CreaCsvFile() (path string, err error) {
	re.countercsv++
	contatore := re.countercsv
	fmt.Println(re.countercsv)
	now := time.Now().Format("2006-01-02 15:04:05")
	path = fmt.Sprintf("registro-%s-%d.csv", now, contatore)
	csvFile, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)
	column := []string{"Movimento", "Datainserimento", "DataMovimento", "Numeromovimento", "Quantita"}
	csvwriter.Write(column)
	for _, movimento := range re.ListaMovimenti {
		row := []string{movimento.TipoMovimento, movimento.DataInserimento.String(), movimento.DataMovimento.String(), strconv.Itoa(movimento.Numero), strconv.Itoa(movimento.Quantita)}
		csvwriter.Write(row)
	}
	csvwriter.Flush()
	csvFile.Close()
	return
}

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

// TrovaIndice cerca il movimento con id uguale a "id" nello slice "lista" e ne restituisce l'indice (cioè la posizione all'interno dello slice)
// o -1 se non lo trova
func trovaIndice(id int, lista []Movimento) int {
	for i, v := range lista {
		if v.ID == id {
			return i
		}
	}
	return -1
}
