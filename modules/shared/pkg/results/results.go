package results

import (
	"os"
	"path/filepath"

	"github.com/kubevirt/kubevirt-tekton-tasks/modules/shared/pkg/env"
)

func RecordResults(results map[string]string) error {
	return RecordResultsIn(env.GetTektonResultsDir(), results)
}

func RecordResultsIn(destination string, results map[string]string) error {
	if results == nil || len(results) == 0 {
		return nil
	}

	for resKey, resVal := range results {
		filename := filepath.Join(destination, resKey)
		err := os.WriteFile(filename, []byte(resVal), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func TektonResultDirExists() error {
	_, err := os.Stat(env.GetTektonResultsDir())

	return err
}
