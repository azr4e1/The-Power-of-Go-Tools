package yamlparse_test

import (
	"os"
	"testing"
	"time"
	yp "yamlparse"

	"github.com/google/go-cmp/cmp"
)

func TestParseReturnsCorrectlyFormattedConfigStruct(t *testing.T) {
	t.Parallel()
	want := yp.Config{
		Global: yp.GlobalConfig{
			ScrapeInterval:     15 * time.Second,
			EvaluationInterval: 30 * time.Second,
			ScrapeTimeout:      10 * time.Second,
			ExternalLabels: map[string]string{
				"monitor": "codelab",
				"foo":     "bar",
			},
		},
	}

	filename := "testdata/yamlConfig.yml"
	got, err := yp.ConfigFromYAML(filename)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}

}

func TestParseReturnsErrorWhenReadingNonExistingFile(t *testing.T) {
	t.Parallel()
	filename := "bogus"
	_, err := yp.ConfigFromYAML(filename)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestParseReturnsErrorWhenItDoesntHavePermissions(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/myyaml.yml"
	_, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chmod(path, 0o000)
	if err != nil {
		t.Fatal(err)
	}

	_, err = yp.ConfigFromYAML(path)
	if err == nil {
		t.Error("expected error, got nil")
	}

}

func TestParseReturnsErrorWhenIncorrectConfiguration(t *testing.T) {
	t.Parallel()
	filename := "testdata/yamlConfigIncorrect.yml"
	_, err := yp.ConfigFromYAML(filename)
	if err == nil {
		t.Error("expected error, got nil")
	}
}
