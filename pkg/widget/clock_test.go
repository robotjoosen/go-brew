package widget_test

import (
	"github.com/robotjoosen/go-brew/pkg/brew"
	"github.com/robotjoosen/go-brew/pkg/widget"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClockGenerate(t *testing.T) {
	testCases := map[string]struct {
		withSchema          []brew.Pour
		withPosition        int
		expectsResponse     string
		expectsError        bool
		expectsErrorMessage string
	}{
		"1st_zero_second": {
			withSchema: []brew.Pour{
				{Duration: 24 * time.Second},
				{Duration: 48 * time.Second},
				{Duration: 96 * time.Second},
				{Duration: 192 * time.Second},
			},
			withPosition: 0,
			expectsResponse: "    ○  *\n" +
				" ○        *\n" +
				"○   pour   *\n" +
				"○          *\n" +
				" ○        *\n" +
				"    ○  ○",
		},
		"1st_half_way": {
			withSchema: []brew.Pour{
				{Duration: 24 * time.Second},
				{Duration: 48 * time.Second},
				{Duration: 96 * time.Second},
				{Duration: 192 * time.Second},
			},
			withPosition: 12,
			expectsResponse: "    ○  ●\n" +
				" ○        ●\n" +
				"○          ●\n" +
				"○   wait   ●\n" +
				" ○        ●\n" +
				"    ○  ●",
		},
		"1st_full_circle": {
			withSchema: []brew.Pour{
				{Duration: 24 * time.Second},
				{Duration: 48 * time.Second},
				{Duration: 96 * time.Second},
				{Duration: 192 * time.Second},
			},
			withPosition: 24,
			expectsResponse: "    ●  ●\n" +
				" ●        ●\n" +
				"●          ●\n" +
				"●   wait   ●\n" +
				" ●        ●\n" +
				"    ●  ●",
		},
		"3rd_half_way": {
			withSchema: []brew.Pour{
				{Duration: 24 * time.Second},
				{Duration: 48 * time.Second},
				{Duration: 96 * time.Second},
				{Duration: 192 * time.Second},
			},
			withPosition: 120, // 24 + 48 + (96 / 2)
			expectsResponse: "    ○  ●\n" +
				" ○        ●\n" +
				"○          ●\n" +
				"○   wait   ●\n" +
				" ○        ●\n" +
				"    ○  ●",
		},
		"4th_half_way": {
			withSchema: []brew.Pour{
				{Duration: 24 * time.Second},
				{Duration: 48 * time.Second},
				{Duration: 96 * time.Second},
				{Duration: 192 * time.Second},
			},
			withPosition: 264, // 24 + 48 + 96 + (192 / 2)
			expectsResponse: "    ○  ●\n" +
				" ○        ●\n" +
				"○          ●\n" +
				"○   wait   ●\n" +
				" ○        ●\n" +
				"    ○  ●",
		},
		"end": {
			withSchema: []brew.Pour{
				{Duration: 24 * time.Second},
				{Duration: 48 * time.Second},
				{Duration: 96 * time.Second},
				{Duration: 192 * time.Second},
			},
			withPosition: 360, // 24 + 48 + 96 + 192
			expectsResponse: "    ●  ●\n" +
				" ●        ●\n" +
				"●          ●\n" +
				"●   wait   ●\n" +
				" ●        ●\n" +
				"    ●  ●",
		},
		"overflow": {
			withSchema: []brew.Pour{
				{Duration: 24 * time.Second},
				{Duration: 48 * time.Second},
				{Duration: 96 * time.Second},
				{Duration: 192 * time.Second},
			},
			withPosition:        360, // 24 + 48 + 96 + 192 + 1
			expectsResponse:     "",
			expectsError:        true,
			expectsErrorMessage: "no pour eligible",
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			d := widget.NewClock(tc.withSchema)
			response, err := d.Render(tc.withPosition)
			if tc.expectsError {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.expectsErrorMessage)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.expectsResponse, response)
		})
	}
}
