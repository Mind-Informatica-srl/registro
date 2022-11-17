package registro

import "time"

type Movimento struct {
	ID              int
	Numero          int
	DataMovimento   time.Time
	DataInserimento time.Time
	Quantita        int
	TipoMovimento   string
}
