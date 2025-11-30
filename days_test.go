package test

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectations = map[string]string{
	"1-1":  "800139",
	"1-2":  "59885340",
	"2-1":  "582",
	"2-2":  "729",
	"3-1":  "299",
	"3-2":  "3621285278",
	"4-1":  "226",
	"4-2":  "160",
	"5-1":  "976",
	"5-2":  "685",
	"6-1":  "6714",
	"6-2":  "3435",
	"7-1":  "278",
	"7-2":  "45157",
	"8-1":  "1331",
	"8-2":  "1121",
	"9-1":  "29221323",
	"9-2":  "4389369",
	"10-1": "1625",
	"10-2": "3100448333024",
	"11-1": "2166",
	"11-2": "1955",
	"12-1": "582",
	"12-2": "52069",
	"13-1": "174",
	"13-2": "780601154795940",
	"14-1": "13476250121721",
	"14-2": "4463708436768",
	"15-1": "1025",
	"15-2": "129262",
}

func TestDays(t *testing.T) {
	for day, expect := range expectations {
		t.Run(day, func(t *testing.T) {
			t.Parallel()
			runCmd := exec.Command("go", "run", ".")
			runCmd.Dir = filepath.Join("days", day)
			output, err := runCmd.CombinedOutput()
			if err != nil {
				fmt.Println(string(output))
			}

			assert.NoError(t, err)
			assert.Equal(t, expect, strings.TrimRight(string(output), "\n"), fmt.Sprintf("Day %s", day))
		})
	}
}
