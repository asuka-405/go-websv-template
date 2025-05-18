package prettyprint

import (
	"fmt"
	"sort"
)

func MapCompact[K comparable, V any](m map[K]V, format string) {
	fmt.Printf(format+"\n", m)
}

// PrettyPrintMap prints map in a sorted key format (only works with string keys).
func Map[V any](m map[string]V) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %v\n", k, m[k])
	}
}
