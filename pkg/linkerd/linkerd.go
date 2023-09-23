package linkerd

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

// CheckConfig holds the linkerd check output.
type CheckConfig struct {
	Success    bool       `json:"success,omitempty" yaml:"success,omitempty"`
	Categories []Category `json:"categories,omitempty" yaml:"categories,omitempty"`
}

// Category holds the Category part of linkerd check's output.
type Category struct {
	Name   string  `json:"categoryName,omitempty" yaml:"categoryName,omitempty"`
	Checks []Check `json:"checks,omitempty" yaml:"checks,omitempty"`
}

// Check holds the actual result linkerd check's output.
type Check struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Result      string `json:"result,omitempty" yaml:"result,omitempty"`
	Hint        string `json:"hint,omitempty" yaml:"hint,omitempty"`
	Error       string `json:"error,omitempty" yaml:"error,omitempty"`
}

func GetCheckConfig(input *json.Decoder) (*CheckConfig, error) {
	var data CheckConfig

	if err := input.Decode(&data); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &data, nil
}

func (analyse *Analyse) SetLogger(logger *logrus.Logger) {
	analyse.logger = logger
}
