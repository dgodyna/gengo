// c imports non-prod code. It shouldn't.
package d

import (
	_ "github.com/dgodyna/gengo/examples/import-boss/tests/inverse/lib/nonprod"
)
