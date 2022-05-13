package utils

import (
	"io"
	"log"
)

func HandleClose(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			log.Panicf("unable to close %T: %s\n", closer, err.Error())
		}
	}
}
