package registro

import "errors"

type Cer string

func CerValidator(cer Cer) error {
	count := 0
	for _, v := range cer {
		if 49 <= v && v <= 57 {
			count++
		}
	}
	if count != 6 {
		return errors.New("Codice cer non ha 6 cifre")
	} else {
		return nil
	}
}
