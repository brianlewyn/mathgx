package mathgx

import "fmt"

// ! My Custom Error Message.

type errGo struct {
	Happd string
}

func myErrGo(h string) *errGo {
	return &errGo{Happd: h}
}
func (e *errGo) Error() string {
	return fmt.Sprintf("\nError: %s", e.Happd)
}
