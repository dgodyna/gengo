// c imports the library root, which in turn imports the public and private packages. This is fine because
// the private package is not directly imported.
package c

import (
	_ "github.com/dgodyna/gengo/examples/import-boss/tests/inverse/lib"
)
