package writers

import "github.com/easm-toolbox/gowitness/pkg/models"

// Writer is a results writer
type Writer interface {
	Write(*models.Result) error
}
