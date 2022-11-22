package registro

// func TestModifyMovimentoToRegistro(t *testing.T) {
// 	// ci aggiungo due movimenti
// 	movimento1 := Movimento{
// 		ID:              1,
// 		Numero:          2,
// 		DataMovimento:   time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
// 		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
// 		Quantita:        400,
// 		TipoMovimento:   "carico",
// 	}
// 	movimento2 := Movimento{
// 		ID:              2,
// 		Numero:          3,
// 		DataMovimento:   time.Date(2020, 8, 15, 14, 30, 45, 100, time.Local),
// 		DataInserimento: time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local),
// 		Quantita:        500,
// 		TipoMovimento:   "carico",
// 	}
// 	// istanzio un registro con già dei movimenti
// 	testRegistro := Registro{
// 		ID:             1,
// 		ListaMovimenti: []Movimento{movimento1, movimento2},
// 	}
// 	// creo nuovo movimento che scambierò con il movimento 1
// 	movimento3 := Movimento{
// 		Numero:        3,
// 		DataMovimento: time.Date(2019, 8, 15, 14, 30, 45, 100, time.Local),
// 		Quantita:      800,
// 		TipoMovimento: "scarico",
// 	}
// 	// attuo modifica
// 	err := testRegistro.ModifyMovimentoToRegistro(movimento1.ID, &movimento3)
// 	// controllo che non abbia restituito errore
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	// mi aspetto che il movimento con id 1 sia cambiato e che sia uguale a movimento3
// 	for _, v := range testRegistro.ListaMovimenti {
// 		if v.ID == movimento1.ID {
// 			if v.TipoMovimento != movimento3.TipoMovimento && v.Quantita != movimento3.Quantita && v.DataMovimento != movimento3.DataMovimento {
// 				t.Error("I movimenti sono diversi")
// 			}
// 		}
// 	}
// 	// controllo che il primo in ordine nel registro sia il movimento modificato dato che data movimento 2019
// 	if testRegistro.ListaMovimenti[0].ID != movimento1.ID {
// 		t.Error("Le date sono diverse quindi registro non oridinato")
// 	}
// }
