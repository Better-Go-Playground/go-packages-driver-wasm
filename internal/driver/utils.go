package driver

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func uniqueStrings(in []string) []string {
	if len(in) == 0 {
		return nil
	}

	sort.Strings(in)

	out := make([]string, 0, len(in))
	var prev string
	for i, item := range in {
		if i == 0 || item != prev {
			out = append(out, item)
			prev = item
		}
	}
	return out
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func relativeToBase(baseDir, target string) (string, bool) {
	rel, err := filepath.Rel(baseDir, target)
	if err != nil {
		return "", false
	}
	if rel == "." {
		return ".", true
	}
	if strings.HasPrefix(rel, ".."+string(filepath.Separator)) || rel == ".." {
		return "", false
	}
	return rel, true
}

func keysSorted[T any](m map[string]T) []string {
	out := make([]string, 0, len(m))
	for key := range m {
		out = append(out, key)
	}
	sort.Strings(out)
	return out
}

