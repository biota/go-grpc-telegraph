package util

import (
	"reflect"
	"sort"
	"testing"
)

// struct with no tags.
type noTagsTest struct {
	name   string
	flag   bool
	number int
	data   []byte
}

// struct with empty tags.
type emptyTagsTest struct {
	empty   string  ``
	flag    bool    ``
	number  int     ``
	floater float64 ``
	inter   int64   ``
	biter   []byte  ``
}

// struct with counter tags.
type counterTags struct {
	empty string  ``
	one   bool    `counter:"one"`
	two   int     `counter:"two"`
	three []byte  `counter:"three"`
	four  int64   `counter:"four"`
	five  float64 `counter:"five"`
}

// struct with env tags.
type envTags struct {
	empty          string ``
	name           bool   `env:"TEST_NAME"`
	cert           string `env:"CERT"`
	key            string `env:"KEY"`
	enable         bool   `env:"ENABLE_FLAG"`
	maxMessageSize uint32 `env:"MAX_MSG_SIZE"`
	maxRetries     uint32 `env:"MAX_RETRY_QSIZE"`
	debug          bool   `env:"DEBUG"`
}

// struct with lang tags.
type langTags struct {
	empty  string `lang:""`
	uno    bool   `lang:"espanol"`
	two    int    `lang:"english"`
	mittsu int    `lang:"nihongo"`
	quatre int    `lang:"french"`
	paanch int    `lang:"hindi"`
}

// tags at different levels.
type levelTags struct {
	level_1 struct {
		uno bool `env:"L1"`

		level_2 struct {
			deux int `env:"L2"`

			level_3 struct {
				mittsu float64 `env:"L3"`

				level_4 struct {
					vier int64 `env:"L4"`

					level_5 struct {
						paanch string `env:"L5"`
					}
				}
			}
		}
	}
}

// struct with mixed tags.
type mixedTagsTest struct {
	empty  string  `env:""`
	uno    bool    `env:"field-1"`
	deux   int     `env:"field-2"`
	mittsu float64 `env:"field-3"`
	vier   int64   `env:"field-4"`
	paanch string  `env:"field-5"`

	// no jo{el,y} tests!
	embid_1 struct {
		// no tags at this level.
		embid_2 struct {
			liu     []byte `env:"field-6"`
			chil    int    `env:"field-7"`
			embid_3 struct {
				levelTags // embedded struct

				zortzi  bool `env:"field-8"`
				embid_4 struct {
					tet   int64   `env:"field-9"`
					ashar float64 `env:"field-10"`
				}
			}
		}
	}

	enya         bool
	biz          float32 `bizzy:"biz"`
	otherPackage int     `foo:"env_other,omitempty"`
	otherTags    float64 `bizzy:"foo,name=bar,tag=t1"`
	badabing     bool    `bizzy:""`
	nunya        string  `:"mtv"`
	nada         string  `:"vh1"`
}

// Test StructTags function.
func TestStructTags(t *testing.T) {
	units := []struct {
		name     string
		tag      string
		obj      any
		expected []string
	}{
		{
			name:     "no tags env test",
			tag:      "env",
			obj:      noTagsTest{},
			expected: []string{},
		},
		{
			name:     "no tags foo test",
			tag:      "foo",
			obj:      noTagsTest{},
			expected: []string{},
		},
		{
			name:     "no tags empty test",
			tag:      "",
			obj:      noTagsTest{},
			expected: []string{},
		},
		{
			name:     "empty tags env test",
			tag:      "env",
			obj:      emptyTagsTest{},
			expected: []string{},
		},
		{
			name:     "empty tags empty test",
			tag:      "",
			obj:      emptyTagsTest{},
			expected: []string{},
		},
		{
			name: "counters tags test",
			tag:  "counter",
			obj:  counterTags{},
			expected: []string{"one", "two", "three", "four",
				"five",
			},
		},
		{
			name:     "counters tags env test",
			tag:      "env",
			obj:      counterTags{},
			expected: []string{},
		},
		{
			name: "env tags env test",
			tag:  "env",
			obj:  envTags{},
			expected: []string{"TEST_NAME", "CERT", "KEY",
				"ENABLE_FLAG", "MAX_MSG_SIZE",
				"MAX_RETRY_QSIZE", "DEBUG",
			},
		},
		{
			name:     "env tags missing test",
			tag:      "missing-404-not-found",
			obj:      envTags{},
			expected: []string{},
		},
		{
			name:     "env tags empty test",
			tag:      "",
			obj:      envTags{},
			expected: []string{},
		},
		{
			name: "lang tags test",
			tag:  "lang",
			obj:  langTags{},
			expected: []string{"espanol", "english", "nihongo",
				"french", "hindi",
			},
		},
		{
			name:     "lang tags env test",
			tag:      "env",
			obj:      langTags{},
			expected: []string{},
		},
		{
			name:     "level tags test",
			tag:      "level",
			obj:      levelTags{},
			expected: []string{},
		},
		{
			name:     "level tags env test",
			tag:      "env",
			obj:      levelTags{},
			expected: []string{"L1", "L2", "L3", "L4", "L5"},
		},
		{
			name: "mixed tags env test",
			tag:  "env",
			obj:  mixedTagsTest{},
			expected: []string{"field-1", "field-2", "field-3",
				"field-4", "field-5", "field-6", "field-7",
				"field-8", "field-9", "field-10",
				"L1", "L2", "L3", "L4", "L5",
			},
		},
		{
			name:     "mixed tags foo test",
			tag:      "foo",
			obj:      mixedTagsTest{},
			expected: []string{"env_other,omitempty"},
		},
		{
			name:     "mixed tags bizzy test",
			tag:      "bizzy",
			obj:      mixedTagsTest{},
			expected: []string{"biz", "foo,name=bar,tag=t1"},
		},
		{
			name:     "mixed tags empty test",
			tag:      "",
			obj:      mixedTagsTest{},
			expected: []string{},
		},
	}

	for _, step := range units {
		t.Logf("Running test '%v', tag=%v ...", step.name, step.tag)

		tags := StructTags(step.obj, step.tag)

		sort.Sort(sort.StringSlice(tags))
		sort.Sort(sort.StringSlice(step.expected))

		if !reflect.DeepEqual(tags, step.expected) {
			t.Errorf("test '%v' for tag %v expected %v, got %v",
				step.name, step.tag, tags, step.expected)
		}
	}

} //  End of function  TestStructTags.
