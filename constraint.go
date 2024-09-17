package compatibility

import (
	"strings"

	"github.com/hashicorp/go-version"
)

type Constraint struct {
	Constraints             version.Constraints
	ConsiderPreReleaseLabel bool
	PermitAny               bool
}

func (o Constraint) Check(v *version.Version) bool {
	if !o.ConsiderPreReleaseLabel {
		// drop the pre-release label
		v = v.Core()
	}

	if !o.PermitAny {
		// v must satisfy all constraints
		return o.Constraints.Check(v)
	}

	// v can satisfy any constraint
	for _, constraint := range o.Constraints {
		if constraint.Check(v) {
			return true
		}
	}

	// v does not satisfy any constraint
	return false
}

func (o Constraint) String() string {
	result := make([]string, len(o.Constraints))
	for i, constraint := range o.Constraints {
		result[i] = constraint.String()
	}

	return strings.Join(result, ",")
}
