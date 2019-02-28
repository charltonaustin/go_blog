package interfaces

import "io"

type Executor interface {
	Execute(wr io.Writer, data interface{}) error
}
