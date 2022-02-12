package resume_test

import (
	"os"
	"testing"

	"github.com/antonve/portfolio-api/test"
)

func TestMain(m *testing.M) {
	test.SetupDatabase(m)
	code := m.Run()
	defer os.Exit(code)
}
