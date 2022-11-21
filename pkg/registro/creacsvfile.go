package registro

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

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
