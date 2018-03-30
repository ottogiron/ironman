package strvals

import (
	"github.com/ironman-project/ironman/template/values"
	"github.com/pkg/errors"
	"k8s.io/helm/pkg/strvals"
)

var _ values.Reader = (*reader)(nil)

type reader struct {
	vals string
}

//New returns a new instance of a flags values reader
//flags in the form of key=value, key=value1
//
func New(vals string) values.Reader {
	return &reader{
		vals,
	}
}

func (r *reader) Read() (values.Values, error) {
	valsMap, err := strvals.Parse(r.vals)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to read values from string %s", r.vals)
	}
	return values.Values(valsMap), nil
}
