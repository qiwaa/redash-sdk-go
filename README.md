[![CircleCI](https://circleci.com/gh/koooge/redash-sdk-go/tree/master.svg?style=svg)](https://circleci.com/gh/koooge/redash-sdk-go/tree/master)

# redash-sdk-go
Redash client in go. Many APIs are still WIP.

## Prerequisites
- go1.11+

## Usage
```
package main

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
)

const queryId = 1

func main() {
	config := &redash.Config{
		EndpointUrl: os.Getenv("REDASH_ENDPOINT_URL"),
		ApiKey:      os.Getenv("REDASH_API_KEY"),
	}
	client := redash.NewClient(config)

	input := &redash.GetQueryInput{
		QueryId: queryId,
	}

	output := client.GetQuery(input)
	fmt.Println(output.Body)
	fmt.Println(output.StatusCode)
}
```

See more about [sample](https://github.com/koooge/redash-sdk-go/tree/master/sample) and [doc](https://github.com/koooge/redash-sdk-go/tree/master/doc).
