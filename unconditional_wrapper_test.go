package wrap

import "testing"

func TestUnconditonalWrappper(t *testing.T) {
	testcases := []struct {
		name  string
		text  string
		limit int
		want  string
	}{
		{
			name:  "testcase 1",
			text:  "",
			limit: 3,
			want:  "",
		},
		{
			name:  "testcase 2",
			text:  " ",
			limit: 3,
			want:  " ",
		},
		{
			name:  "testcase 3",
			text:  "          ",
			limit: 3,
			want:  "   \n   \n   \n ",
		},
		{
			name:  "testcase 4",
			text:  "foo bar baz",
			limit: 3,
			want:  "foo\n ba\nr b\naz",
		},
		{
			name:  "testcase 5",
			text:  "foo bar baz",
			limit: 2,
			want:  "fo\no \nba\nr \nba\nz",
		},
		{
			name:  "testcase 6",
			text:  "foo bar baz",
			limit: 1,
			want:  "f\no\no\n \nb\na\nr\n \nb\na\nz",
		},
		{
			name:  "testcase 7",
			text:  "foo bar baz",
			limit: 0,
			want:  "",
		},
		{
			name:  "testcase 8",
			text:  " foo bar ",
			limit: 3,
			want:  " fo\no b\nar ",
		},
		{
			name:  "testcase 9",
			text:  " foo bar ",
			limit: 2,
			want:  " f\noo\n b\nar\n ",
		},
		{
			name:  "testcase 10",
			text:  " foo bar ",
			limit: 1,
			want:  " \nf\no\no\n \nb\na\nr\n ",
		},
		{
			name:  "testcase 11",
			text:  " foo bar ",
			limit: 0,
			want:  "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := UnconditionallyWrap(tc.text, tc.limit)
			if got != tc.want {
				t.Errorf("wanted %q, got %q", tc.want, got)
			}
		})
	}
}
