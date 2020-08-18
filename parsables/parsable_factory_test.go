package parsables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetParsable(t *testing.T) {
	parsableNames := [3]string{
		"aws_iam_user",
		"onelogin_users",
		"onelogin_apps",
		// add parsables by provider_resourceName here
	}
	tests := map[string]struct {
		Parsables *ParsableList
		ShouldErr bool
	}{
		"It creates and returns parsables": {
			Parsables: New(),
		},
		"It logs an error and returns empty parsable if parsable not defined": {
			Parsables: New(),
			ShouldErr: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if !test.ShouldErr {
				for _, name := range parsableNames {
					parsable := test.Parsables.GetParsable(name)
					memoizedParsable := test.Parsables.GetParsable(name)
					assert.Equal(t, test.Parsables.parsables[name], parsable)
					assert.Equal(t, test.Parsables.parsables[name], memoizedParsable)
				}
			} else {
				parsable := test.Parsables.GetParsable("missing")
				assert.Equal(t, nil, parsable)
			}
		})
	}
}
