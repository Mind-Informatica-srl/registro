package main

import (
	"fmt"
	"goProject/pkg/registro"
	"time"
)

func main() {

	// creare instanza registro
	instanzaRegistro := registro.Registro{
		ID:             1,
		ListaMovimenti: []registro.Movimento{},
	}
	movimento1 := registro.Movimento{
		ID:              1,
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        400,
		TipoMovimento:   "carico",
	}
	movimento2 := registro.Movimento{
		ID:              2,
		Numero:          3,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "carico",
	}
	movimento3 := registro.Movimento{
		ID:              3,
		Numero:          4,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        600,
		TipoMovimento:   "scarico",
	}
	movimento4 := registro.Movimento{
		ID:              4,
		Numero:          5,
		DataMovimento:   time.Date(2022, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2022, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        700,
		TipoMovimento:   "carico",
	}
	// aggiungere movimenti al registro
	instanzaRegistro.AddMovimentoToRegistro(&movimento1)
	instanzaRegistro.AddMovimentoToRegistro(&movimento2)
	instanzaRegistro.AddMovimentoToRegistro(&movimento3)
	instanzaRegistro.AddMovimentoToRegistro(&movimento4)
	instanzaRegistro.CreaCsvFile()

	registro1 := registro.Registro{
		ID: 1,
	}
	if path1, err := registro1.CreaCsvFile(); err == nil {
		fmt.Println(path1)
	}
	registro2 := registro.Registro{
		ID: 2,
	}
	if path2, err := registro2.CreaCsvFile(); err == nil {
		fmt.Println(path2)
	}
	if path2, err := registro2.CreaCsvFile(); err == nil {
		fmt.Println(path2)
	}

}

func returnaNumero2() int {
	return 2
}
