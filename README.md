[Website](https://appscode.com) • [Slack](https://slack.appscode.com) • [Forum](https://discuss.appscode.com) • [Twitter](https://twitter.com/AppsCodeHQ)

# cloudid
Detect Cloud Provider from Instance Metadata

## Documentation
https://godoc.org/github.com/appscode/cloudid

## Example
```go
package main

import (
	"fmt"

	"github.com/appscode/cloudid"
)

func main() {
	fmt.Println("Cloud Provider:", cloudid.Detect())
}
```
