package main

import "github.com/multitheftauto/build-tools/internal/rescheck"

type dupeRecommendation struct {
	Version      string
	Line         int
	ShouldDelete bool
}

// TODO(qaisjp): we recommend to keep the first one, but we could do a better job of suggesting which one to remove
func findDuplicates(items []rescheck.VersionItem) map[string][]dupeRecommendation {
	funcRecs := map[string][]dupeRecommendation{}

	for i, item := range items {
		recs, ok := funcRecs[item.FunctionName]

		// If the func already exists, append the version
		if ok {
			funcRecs[item.FunctionName] = append(recs, dupeRecommendation{item.MinMtaVersion, i, true})
		} else {
			funcRecs[item.FunctionName] = []dupeRecommendation{dupeRecommendation{item.MinMtaVersion, i, false}}
		}
	}

	// Now filter out the single items
	for fname, recs := range funcRecs {
		if len(recs) == 1 {
			delete(funcRecs, fname)
		}
	}

	return funcRecs
}
