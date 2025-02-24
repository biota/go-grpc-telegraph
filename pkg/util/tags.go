package util

import (
	"reflect"
)

// Search for matching `tag` structure field tags.
// For example, for a struct like:
//
//	type Record struct {
//		Name    string `hashtag:"NAME"`
//		Address string `hashtag:"ADDRESS"`
//		Hidden  int
//		Private bool
//		Key     string `hashtag:"KEY"`
//		Empty   int ``
//	}
//
// searching for `hashtag` aka
//
//	tagSearch(reflect.ValueOf(Record{}), "hashtag") will return
//
// will return:
//
//	["NAME", "ADDRESS", "KEY"]
func tagSearch(value reflect.Value, tag string) []string {
	matches := make([]string, 0)
	kind := value.Type()

	for idx := 0; idx < value.NumField(); idx++ {
		v := value.Field(idx)

		// Check for inner/embedded structs.
		if v.IsValid() && v.Kind() == reflect.Struct {
			results := tagSearch(v, tag)
			matches = append(matches, results[:]...)
			continue
		}

		field := kind.Field(idx)
		if etag := field.Tag.Get(tag); etag != "" {
			matches = append(matches, etag)
		}
	}

	return matches

} // End of function  tagSearch.

// Search struct fields tags that match `tag` and return matched values.
func StructTags(record any, tag string) []string {
	return tagSearch(reflect.ValueOf(record), tag)

} // End of function  StructTags.
