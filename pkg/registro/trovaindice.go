package registro

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
