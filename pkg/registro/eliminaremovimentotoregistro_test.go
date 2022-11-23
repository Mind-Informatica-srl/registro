package registro

// func TestEliminareMovimentoToRegistro(t *testing.T) {
// 	// creo tre instanze movimento
// 	movimento1 := Movimento{
// 		ID:              1,
// 		Numero:          2,
// 		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
// 		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
// 		Quantita:        400,
// 		TipoMovimento:   "carico",
// 		CodiceCer:       "123456",
// 	}
// 	movimento2 := Movimento{
// 		ID:              2,
// 		Numero:          3,
// 		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
// 		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
// 		Quantita:        500,
// 		TipoMovimento:   "carico",
// 		CodiceCer:       "123456",
// 	}
// 	movimento3 := Movimento{
// 		ID:              3,
// 		Numero:          4,
// 		DataMovimento:   time.Date(2020, 9, 15, 14, 30, 45, 100, time.Local),
// 		DataInserimento: time.Date(2021, 9, 15, 14, 30, 45, 100, time.Local),
// 		Quantita:        500,
// 		TipoMovimento:   "scarico",
// 		CodiceCer:       "123456",
// 	}
// 	movimento4 := Movimento{
// 		ID:              3,
// 		Numero:          4,
// 		DataMovimento:   time.Date(2021, 9, 15, 14, 30, 45, 100, time.Local),
// 		DataInserimento: time.Date(2021, 9, 15, 14, 30, 45, 100, time.Local),
// 		Quantita:        500,
// 		TipoMovimento:   "carico",
// 		CodiceCer:       "123456",
// 	}

// 	testRegistro := Registro{
// 		ID: 1,
// 	}
// 	err := testRegistro.AddMovimentoToRegistro(&movimento1)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = testRegistro.AddMovimentoToRegistro(&movimento2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = testRegistro.AddMovimentoToRegistro(&movimento3)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = testRegistro.AddMovimentoToRegistro(&movimento4)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = testRegistro.EliminareMovimentoToRegistro(2)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	// mi aspetto che alla seconda posizione nell'array di movimenti sia nuovamente presente il movimento 2 che prima era stato tolto
// 	// inoltre mi aspetto che avvenga anche il riordinamento dato che lista di movimenti dopo viene riordinata
// 	if testRegistro.ListaMovimenti[1].ID != movimento2.ID {
// 		t.Error("Mi aspettavo che fosse presente in questa posizione dell'array di movimenti il movimento2 avente id 2")
// 	}
// }
