package apprulesschema

import (
	"fmt"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/apps/app_rules"
	"github.com/stretchr/testify/assert"
)

func TestRulesSchema(t *testing.T) {
	t.Run("creates and returns a map of a Rules Schema", func(t *testing.T) {
		provSchema := Schema()
		assert.NotNil(t, provSchema["app_id"])
		assert.NotNil(t, provSchema["name"])
		assert.NotNil(t, provSchema["match"])
		assert.NotNil(t, provSchema["position"])
		assert.NotNil(t, provSchema["conditions"])
		assert.NotNil(t, provSchema["actions"])
	})
}

func TestInflate(t *testing.T) {
	tests := map[string]struct {
		ResourceData   map[string]interface{}
		ExpectedOutput apprules.AppRule
	}{
		"creates and returns the address of a Rule struct": {
			ResourceData: map[string]interface{}{
				"id":       "123",
				"app_id":   "123",
				"name":     "test",
				"match":    "test",
				"enabled":  true,
				"position": 1,
			},
			ExpectedOutput: apprules.AppRule{
				ID:       oltypes.Int32(int32(123)),
				AppID:    oltypes.Int32(int32(123)),
				Name:     oltypes.String("test"),
				Match:    oltypes.String("test"),
				Enabled:  oltypes.Bool(true),
				Position: oltypes.Int32(int32(1)),
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			subj := Inflate(test.ResourceData)
			assert.Equal(t, test.ExpectedOutput, subj)
		})
	}
}

func TestValidMatch(t *testing.T) {
	tests := map[string]struct {
		InputKey       string
		InputValue     string
		ExpectedOutput []error
	}{
		"no errors on valid input": {
			InputKey:       "match",
			InputValue:     "all",
			ExpectedOutput: nil,
		},
		"errors on invalid input": {
			InputKey:       "match",
			InputValue:     "asdf",
			ExpectedOutput: []error{fmt.Errorf("match must be one of [all any], got: asdf")},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, errs := validMatch(test.InputValue, test.InputKey)
			assert.Equal(t, test.ExpectedOutput, errs)
		})
	}
}
