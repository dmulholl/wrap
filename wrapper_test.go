package wrap

import "testing"

func TestWrappper_DefaultBreakpoints(t *testing.T) {
	testcases := []struct {
		name  string
		text  string
		limit int
		want  string
	}{
		{
			name:  "testcase 1",
			text:  "foo bar baz",
			limit: 3,
			want:  "foo\nbar\nbaz",
		},
		{
			name:  "testcase 2",
			text:  "foo bar baz",
			limit: 2,
			want:  "fo\no\nba\nr\nba\nz",
		},
		{
			name:  "testcase 3",
			text:  "foobarbaz",
			limit: 3,
			want:  "foo\nbar\nbaz",
		},
		{
			name:  "testcase 4",
			text:  "abc foobarbaz xyz",
			limit: 3,
			want:  "abc\nfoo\nbar\nbaz\nxyz",
		},
		{
			name:  "testcase 5",
			text:  "x foobar x",
			limit: 3,
			want:  "x\nfoo\nbar\nx",
		},
		{
			name:  "testcase 6",
			text:  "x abcd x",
			limit: 3,
			want:  "x\nabc\nd x",
		},
		{
			name:  "testcase 7",
			text:  "xx abcd xx",
			limit: 3,
			want:  "xx\nabc\nd\nxx",
		},
		{
			name:  "testcase 8",
			text:  "",
			limit: 3,
			want:  "",
		},
		{
			name:  "testcase 9",
			text:  " ",
			limit: 3,
			want:  "",
		},
		{
			name:  "testcase 10",
			text:  "          ",
			limit: 3,
			want:  "",
		},
		{
			name:  "testcase 11",
			text:  " foo bar ",
			limit: 5,
			want:  "foo\nbar",
		},
		{
			name:  "testcase 12",
			text:  " foo bar ",
			limit: 4,
			want:  "foo\nbar",
		},
		{
			name:  "testcase 13",
			text:  " foo bar ",
			limit: 3,
			want:  "foo\nbar",
		},
		{
			name:  "testcase 14",
			text:  " foo bar ",
			limit: 2,
			want:  "fo\no\nba\nr",
		},
		{
			name:  "testcase 15",
			text:  " foo bar ",
			limit: 1,
			want:  "f\no\no\nb\na\nr",
		},
		{
			name:  "testcase 16",
			text:  " foo bar ",
			limit: 0,
			want:  "",
		},
		{
			name:  "testcase 17",
			text:  "foo   bar   \n\tbaz   bam",
			limit: 9,
			want:  "foo bar\nbaz bam",
		},
		{
			name:  "testcase 18",
			text:  "foo   bar   \n\tbaz   bam",
			limit: 8,
			want:  "foo bar\nbaz bam",
		},
		{
			name:  "testcase 19",
			text:  "foo   bar   \n\tbaz   bam",
			limit: 7,
			want:  "foo bar\nbaz bam",
		},
		{
			name:  "testcase 20",
			text:  "foo   bar   \n\tbaz   bam",
			limit: 6,
			want:  "foo\nbar\nbaz\nbam",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := Wrap(tc.text, tc.limit)
			if got != tc.want {
				t.Errorf("wanted %q, got %q", tc.want, got)
			}
		})
	}
}

func TestWrappper_CustomBreakpoints(t *testing.T) {
	testcases := []struct {
		name  string
		text  string
		limit int
		want  string
	}{
		{
			name:  "testcase 1",
			text:  "foo bar baz",
			limit: 3,
			want:  "foo\nbar\nbaz",
		},
		{
			name:  "testcase 2",
			text:  "foo bar baz",
			limit: 2,
			want:  "fo\no\nba\nr\nba\nz",
		},
		{
			name:  "testcase 3",
			text:  "foobarbaz",
			limit: 3,
			want:  "foo\nbar\nbaz",
		},
		{
			name:  "testcase 4",
			text:  "abc foobarbaz xyz",
			limit: 3,
			want:  "abc\nfoo\nbar\nbaz\nxyz",
		},
		{
			name:  "testcase 5",
			text:  "x foobar x",
			limit: 3,
			want:  "x\nfoo\nbar\nx",
		},
		{
			name:  "testcase 6",
			text:  "x abcd x",
			limit: 3,
			want:  "x\nabc\nd x",
		},
		{
			name:  "testcase 7",
			text:  "xx abcd xx",
			limit: 3,
			want:  "xx\nabc\nd\nxx",
		},
		{
			name:  "testcase 8",
			text:  "",
			limit: 3,
			want:  "",
		},
		{
			name:  "testcase 9",
			text:  " ",
			limit: 3,
			want:  "",
		},
		{
			name:  "testcase 10",
			text:  "          ",
			limit: 3,
			want:  "",
		},
		{
			name:  "testcase 11",
			text:  " foo bar ",
			limit: 5,
			want:  "foo\nbar",
		},
		{
			name:  "testcase 12",
			text:  " foo bar ",
			limit: 4,
			want:  "foo\nbar",
		},
		{
			name:  "testcase 13",
			text:  " foo bar ",
			limit: 3,
			want:  "foo\nbar",
		},
		{
			name:  "testcase 14",
			text:  " foo bar ",
			limit: 2,
			want:  "fo\no\nba\nr",
		},
		{
			name:  "testcase 15",
			text:  " foo bar ",
			limit: 1,
			want:  "f\no\no\nb\na\nr",
		},
		{
			name:  "testcase 16",
			text:  " foo bar ",
			limit: 0,
			want:  "",
		},
		{
			name:  "testcase 17",
			text:  " foo.bar ",
			limit: 5,
			want:  "foo.\nbar",
		},
		{
			name:  "testcase 18",
			text:  " foo.bar ",
			limit: 4,
			want:  "foo.\nbar",
		},
		{
			name:  "testcase 19",
			text:  " foo.bar ",
			limit: 3,
			want:  "foo\n.\nbar",
		},
		{
			name:  "testcase 20",
			text:  " foo.bar ",
			limit: 2,
			want:  "fo\no.\nba\nr",
		},
		{
			name:  "testcase 21",
			text:  " foo.bar ",
			limit: 1,
			want:  "f\no\no\n.\nb\na\nr",
		},
		{
			name:  "testcase 22",
			text:  " foo.bar ",
			limit: 0,
			want:  "",
		},
		{
			name:  "testcase 23",
			text:  " foo-bar ",
			limit: 5,
			want:  "foo\n-bar",
		},
		{
			name:  "testcase 24",
			text:  " foo-bar ",
			limit: 4,
			want:  "foo\n-bar",
		},
		{
			name:  "testcase 25",
			text:  " foo-bar ",
			limit: 3,
			want:  "foo\n-ba\nr",
		},
		{
			name:  "testcase 26",
			text:  " foo-bar ",
			limit: 2,
			want:  "fo\no\n-b\nar",
		},
		{
			name:  "testcase 27",
			text:  " foo-bar ",
			limit: 1,
			want:  "f\no\no\n-\nb\na\nr",
		},
		{
			name:  "testcase 28",
			text:  " foo-bar ",
			limit: 0,
			want:  "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			wrapper := NewWrapper()
			wrapper.BreakAfter = []rune{' ', '.'}
			wrapper.BreakBefore = []rune{' ', '-'}

			got := wrapper.Wrap(tc.text, tc.limit)
			if got != tc.want {
				t.Errorf("wanted %q, got %q", tc.want, got)
			}
		})
	}
}
