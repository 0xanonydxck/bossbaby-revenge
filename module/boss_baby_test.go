package module_test

import (
	"bossbaby/module"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string
	Input    string
	Expected string
	Panic    bool
}

var cases []TestCase = []TestCase{
	{
		Name:     "[SRSSRRR]GoodBoy",
		Input:    "SRSSRRR",
		Expected: module.GOOD_BOY,
		Panic:    false,
	},
	{
		Name:     "[SSRSRRR]GoodBoy",
		Input:    "SSRSRRR",
		Expected: module.GOOD_BOY,
		Panic:    false,
	},
	{
		Name:     "[RSSRR]GoodBoy",
		Input:    "RSSRR",
		Expected: module.BAD_BOY,
		Panic:    false,
	},
	{
		Name:     "[SSSRRRRS]BadBoy",
		Input:    "SSSRRRRS",
		Expected: module.BAD_BOY,
		Panic:    false,
	},
	{
		Name:     "[SRRSSR]BadBoy",
		Input:    "SRRSSR",
		Expected: module.BAD_BOY,
		Panic:    false,
	},
	{
		Name:     "[SRRSSR]BadBoy",
		Input:    "SRRSSR",
		Expected: module.BAD_BOY,
		Panic:    false,
	},
	{
		Name:     "[]Panic",
		Input:    "",
		Expected: module.BAD_BOY,
		Panic:    true,
	},
	{
		Name:     "[ABC]Panic",
		Input:    "",
		Expected: module.BAD_BOY,
		Panic:    true,
	},
}

func TestBossBabyRevenge(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			panicAssert := assert.NotPanics
			if test.Panic {
				panicAssert = assert.Panics
			}

			panicAssert(t, func() {
				output := module.BossBabyRevenge(test.Input)
				assert.Equal(t, test.Expected, output)
			})
		})
	}
}

func BenchmarkBossBabyRevenge(b *testing.B) {
	input := "SRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRSRS"
	for i := 0; i < b.N; i++ {
		module.BossBabyRevenge(input)
	}
}
