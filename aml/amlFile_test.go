package aml

import (
	"strings"
	"testing"

	"github.com/windler/awesm/aml/constants"

	"github.com/stretchr/testify/assert"

	"github.com/windler/awesm/aml/instructions"
)

func TestCreateDotGraph(t *testing.T) {
	ins1 := &instructions.AMLInstruction{
		Name:        "node1",
		NodeOptions: make(map[string]string),
		EdgeOptions: make(map[string]string),
	}
	ins2 := &instructions.AMLInstruction{
		Name: "node2",
		NodeOptions: map[string]string{
			"c": "d",
		},
		EdgeOptions: map[string]string{
			"a": "b",
		},
		Predecessors: []*instructions.AMLInstruction{ins1},
	}
	aml := &File{
		GraphType: "mytype",
		Instructions: []instructions.AMLInstruction{
			*ins1,
			*ins2,
		},
	}

	g := aml.CreateDotGraph(constants.TopDowm)

	assert.True(t, strings.Contains(g.String(), `"node1"->"node2"[a="b"]`))
	assert.True(t, strings.Contains(g.String(), `"node2"[c="d"]`))
}
