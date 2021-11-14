// b only imports public and private packages. The latter it shouldn't.
package b

import (
	_ "github.com/dgodyna/gengo/examples/import-boss/tests/inverse/lib/private"
	_ "github.com/dgodyna/gengo/examples/import-boss/tests/inverse/lib/public"
)
