package logcheck

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Razzle131/loglint/config"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func initEnv() {
	cfg := config.Load("")
	_ = NewAnalyzer(cfg)
}

func TestCheckFirstLetterCase(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "ok",
			input:   "good msg",
			wantErr: false,
		},
		{
			name:    "empty",
			input:   "",
			wantErr: false,
		},
		{
			name:    "other language ok",
			input:   "хорошее сообщение",
			wantErr: false,
		},
		{
			name:    "special symbol first",
			input:   "$msg",
			wantErr: false, // because here we check that first symbol is letter and it is lowered
		},
		{
			name:    "bad",
			input:   "Bad msg",
			wantErr: true,
		},
		{
			name:    "other language bad",
			input:   "Плохое сообщение",
			wantErr: true,
		},
	}

	initEnv()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := checkFirstLetterCase(testCase.input)

			if testCase.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestCheckEnglish(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "ok",
			input:   "good msg",
			wantErr: false,
		},
		{
			name:    "empty",
			input:   "",
			wantErr: false,
		},
		{
			name:    "english with special symbols",
			input:   "starting...✅",
			wantErr: false, // should not give error, it is different check
		},
		{
			name:    "not english",
			input:   "русское сообщение",
			wantErr: true,
		},
	}

	initEnv()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := checkEnglish(testCase.input)

			if testCase.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestCheckSpecialSymbols(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "ok",
			input:   "good msg",
			wantErr: false,
		},
		{
			name:    "empty",
			input:   "",
			wantErr: false,
		},
		{
			name:    "other language ok",
			input:   "хорошее сообщение",
			wantErr: false,
		},
		{
			name:    "bad",
			input:   "starting...✅",
			wantErr: true,
		},
		{
			name:    "only special",
			input:   ".",
			wantErr: true,
		},
		{
			name:    "only emoji",
			input:   "✅",
			wantErr: true,
		},
	}

	initEnv()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := checkSpecialSymbols(testCase.input)

			if testCase.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestCheckSensetive(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "ok",
			input:   "num",
			wantErr: false,
		},
		{
			name:    "empty",
			input:   "",
			wantErr: false,
		},
		{
			name:    "bad",
			input:   "apiKey",
			wantErr: true,
		},
	}

	initEnv()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			err := checkSensetive(testCase.input)

			if testCase.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	initEnv()

	testdata := filepath.Join(wd, "testdata")
	analysistest.Run(t, testdata, analyzer, "slog/")
}
