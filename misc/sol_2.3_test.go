package misc

import "testing"

func TestSol2_3(t *testing.T) {
	tt := []struct {
		Word   string
		K      int
		Output string
	}{
		{"abcdefghij klmnop", 4, "abcefgij lmnp"},
	}

	for _, tc := range tt {
		got := Sol2_3(tc.Word, tc.K)
		want := tc.Output
		if got != want {
			t.Errorf("Got (%v) != Want (%v)", got, want)
		}
	}

}
