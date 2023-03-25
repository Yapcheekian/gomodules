package mod5

import (
	"fmt"

	"github.com/google/uuid"
)

func New() {
	fmt.Println("mod5 v25.3.0")
	id := uuid.New()
	fmt.Println(id)
}
