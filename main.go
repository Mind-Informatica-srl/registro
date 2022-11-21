package main

import (
	"goProject/pkg/registro"
	"time"
)

func main() {
	movimento1 := registro.Movimento{
		ID:              1,
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        1500,
		TipoMovimento:   "carico",
		CodiceCer:       "123456",
	}
	movimento2 := registro.Movimento{
		ID:              1,
		Numero:          2,
		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        1500,
		TipoMovimento:   "scarico",
		CodiceCer:       "123456",
	}
	movimento3 := registro.Movimento{
		ID:              2,
		Numero:          3,
		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        500,
		TipoMovimento:   "carico",
		CodiceCer:       "123456",
	}
	movimento4 := registro.Movimento{
		ID:              3,
		Numero:          4,
		DataMovimento:   time.Date(2022, 8, 15, 14, 30, 45, 100, time.Local),
		DataInserimento: time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
		Quantita:        3000,
		TipoMovimento:   "scarico",
		CodiceCer:       "123456",
	}
	newRegistro := registro.Registro{
		ID:             2,
		ListaMovimenti: []registro.Movimento{},
	}
	err := newRegistro.AddMovimentoToRegistro(&movimento1)
	if err != nil {
		panic(err)
	}
	err = newRegistro.AddMovimentoToRegistro(&movimento2)
	if err != nil {
		panic(err)
	}
	err = newRegistro.AddMovimentoToRegistro(&movimento3)
	if err != nil {
		panic(err)
	}
	err = newRegistro.AddMovimentoToRegistro(&movimento4)
	if err != nil {
		panic(err)
	}
}
