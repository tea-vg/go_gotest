package tdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func NewSafeClosedSection(left, right int) ClosedSection {
  section, _ := NewClosedSection(left, right)

  return section
}

func TestEqualは対象の閉区間と等しいかどうか検査する(t *testing.T) {
	tests := map[string]struct {
		section  ClosedSection
		section_ ClosedSection
		want     bool
	}{
		"下端も上端も等しい": {
			section:  NewSafeClosedSection(0, 1),
			section_: NewSafeClosedSection(0, 1),
			want:     true,
		},
		"下端のみ等しい": {
			section:  NewSafeClosedSection(0, 1),
			section_: NewSafeClosedSection(0, 2),
			want:     false,
		},
		"上端のみ等しい": {
			section:  NewSafeClosedSection(0, 1),
			section_: NewSafeClosedSection(-1, 1),
			want:     false,
		},
		"下端も上端も等しくない": {
			section:  NewSafeClosedSection(0, 1),
			section_: NewSafeClosedSection(-1, 2),
			want:     false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.section.Equal(test.section_)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestContainNumberは閉区間がある整数を含むかを検査する(t *testing.T) {
	tests := map[string]struct {
		section ClosedSection
		n       int
		want    bool
	}{
		"[0, 5]が3を含む": {
			section: NewSafeClosedSection(0, 5),
			n:       3,
			want:    true,
		},
		"[0, 5]が0を含む": {
			section: NewSafeClosedSection(0, 5),
			n:       0,
			want:    true,
		},
		"[0, 5]が5を含む": {
			section: NewSafeClosedSection(0, 5),
			n:       5,
			want:    true,
		},
		"[0, 5]が-5を含まない": {
			section: NewSafeClosedSection(0, 5),
			n:       -5,
			want:    false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.section.ContainNum(test.n)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestContainSectionは閉区間が別の閉区間を含むかを検査する(t *testing.T) {
	tests := map[string]struct {
		section ClosedSection
    section_ ClosedSection
		want    bool
	}{
		"[-2, 2]が[-2, -2]を含む": {
			section: NewSafeClosedSection(-2, 2),
      section_: NewSafeClosedSection(-2, 2),
			want:    true,
		},
		"[-2, 2]が[-1, -1]を含む": {
			section: NewSafeClosedSection(-2, 2),
      section_: NewSafeClosedSection(-1, 1),
			want:    true,
		},
		"[-2, 2]が[-2, 1]を含む": {
			section: NewSafeClosedSection(-2, 2),
      section_: NewSafeClosedSection(-2, 1),
			want:    true,
		},
		"[-2, 2]が[-1, 2]を含む": {
			section: NewSafeClosedSection(-2, 2),
      section_: NewSafeClosedSection(-1, 2),
			want:    true,
		},
		"[-2, 2]が[-2, 3]を含まない": {
			section: NewSafeClosedSection(-2, 2),
      section_: NewSafeClosedSection(-2, 3),
			want:    false,
		},
		"[-2, 2]が[-3, 2]を含まない": {
			section: NewSafeClosedSection(-2, 2),
      section_: NewSafeClosedSection(-3, 2),
			want:    false,
		},
		"[-2, 2]が[-3, 3]を含まない": {
			section: NewSafeClosedSection(-2, 2),
      section_: NewSafeClosedSection(-3, 3),
			want:    false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.section.ContainSection(test.section_)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestToStringは閉区間の文字列表現が正しいかを検査する(t *testing.T) {
  section := NewSafeClosedSection(0, 1)
  want := "[0, 1]"

  got := section.ToString()

  assert.Equal(t, want, got)
}

func TestCorrectSectionは正しい閉区間を生成できることを検査する(t *testing.T) {
  tests := map[string]struct {
    left int
    right int
    want ClosedSection
  } {
    "[0, 0]": {left: 0, right: 0, want: NewSafeClosedSection(0, 0)},
    "[0, 1]": {left: 0, right: 1, want: NewSafeClosedSection(0, 1)},
  }

  for name, test := range(tests) {
    t.Run(name, func(t *testing.T) {
      got, err := NewClosedSection(test.left, test.right)

      assert.Equal(t, test.want, got)
      assert.NoError(t, err)
    })
  }
}

func TestInCorrectSectionは誤った閉区間を検知できることを検査する(t *testing.T) {
  tests := map[string]struct {
    left int
    right int
  } {
    "[2, 1]": {left: 2, right: 1},
    "[-1, -2]": {left: -1, right: -2},
  }

  for name, test := range(tests) {
    t.Run(name, func(t *testing.T) {
      _, err := NewClosedSection(test.left, test.right)

      assert.Error(t, err)
    })
  }
}
