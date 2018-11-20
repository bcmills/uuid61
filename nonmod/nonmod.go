package nonmod

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func PrintUUID() {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
