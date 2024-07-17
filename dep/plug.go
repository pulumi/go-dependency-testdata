package dep

import indirect "github.com/pulumi/go-dependency-testdata/indirect-dep/v2"

func Bar() string {
	indirect.Baz()
	return "bar"
}
