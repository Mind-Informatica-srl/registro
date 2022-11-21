package registro

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestCSVFile(t *testing.T) {
	// creo tre instanze movimento
	movimento1 := Movimento{
		ID:              1,
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	StringQuantita := strconv.Itoa(movimento1.Quantita)
	StringaNumero := strconv.Itoa(movimento1.Numero)
	StringaDataMovimento := movimento1.DataMovimento.String()
	StringaDataInserimento := movimento1.DataInserimento.String()

	// istanzio un registro con i tre movimenti instanziati precedentemente
	testRegistro := Registro{
		ID:             1,
		ListaMovimenti: []Movimento{movimento1},
	}
	path, err := testRegistro.CreaCsvFile()
	if err != nil {
		t.Error(err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		// se errore torna  ovvero l'errore che ritorna quando non ci sono più linee da leggere faccio un break
		if err == io.EOF {
			break
		}
		// se invece torna un errore diverso da quello da io.EOF e quindi torna errore
		if err != nil {
			t.Error(err)
		}
		// prendo record ovvero riga del csv che è uno slice
		recordpreso := record
		// se record che ho preso risulta come lo avevo previsto ovvero alla prima riga le stringhe delle informazioni del movimento e al secondo record che "itero" i dati relativi al movimento 1 in stringa allora non fare nulla
		if (recordpreso[0] == "Movimento" && recordpreso[1] == "Datainserimento" && recordpreso[2] == "DataMovimento" && recordpreso[3] == "Numeromovimento" && recordpreso[4] == "Quantita") || (recordpreso[0] == movimento1.TipoMovimento && recordpreso[1] == StringaDataInserimento && recordpreso[2] == StringaDataMovimento && recordpreso[3] == StringaNumero && recordpreso[4] == StringQuantita) {
			fmt.Println("")
		} else { // altrimenti
			// printare risultati delle previsioni che avevo fatto sono sbagliati
			t.Error("Record del csv diversi da quelli previsti")
		}
	}
	os.Remove(path)
}
