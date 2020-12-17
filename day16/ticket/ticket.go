package ticket

import (
	"encoding/json"
	"fmt"
)

type Ticket []int

func (t *Ticket) Parse(s string) error {
	data := []byte(fmt.Sprintf("[%s]", s))
	err := json.Unmarshal(data, t)
	return err
}
