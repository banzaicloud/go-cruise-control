# go-cruise-control

It's client library (written in Golang) for interacting with
[Linkedin Cruise Control](https://github.com/linkedin/cruise-control) using its HTTP API.

Supported _Cruise Control_ versions: **2.5.x-2.5.93** (tested with v2.5.86)

## How to use it

```shell
go get github.com/banzaicloud/go-cruise-control@latest
```

```go
package main

import (
	"fmt"

	"github.com/banzaicloud/go-cruise-control/client"
	"github.com/banzaicloud/go-cruise-control/api"
)

func main() {
    // Initializing Cruise Control client with default configuration
    cruisecontrol, err := client.NewDefaultClient()
    if err != nil {
        panic(err)
    }
    
    // Assembling the request using default values
    req := api.StateRequestWithDefaults()
    // Sending the request to the State API
    resp, err := cruisecontrol.State(req)
    if err != nil {
        panic(err)
    }
    // Getting the state of the Executor
    fmt.Println(resp.Result.ExecutorState.State)
}
```

## Development

### Prerequisites

* golang `1.18`
* docker
* docker-compose

### Development environment

Starting local development environment

```shell
make env-up
```

Tearing down the local development environment

```shell
make env-clean
```

### Testing

Running integration tests creates a new test environment prior executing the test suite and tears down at the end.

```shell
make integration-test
```

Set the `USE_EXISTING=true` environment variable if reusing the existing development/test environment
for integration testing is the desired behaviour.

```shell
export USE_EXISTING=true
make integration-test
```

## Support

If you encounter problems while using this project than open an issue or talk to us
on `#kafka-operator` channel on the [Banzai Cloud Community Slack](http://community-banzaicloud.slack.com).

## Contributing

If you find this project useful, help us:

* Support the development of this project and star this repo! ⭐
* Help new users with issues they may encounter 💪
* Send a pull request with your new features and bug fixes 🚀

## License

Copyright © 2021 Cisco and/or its affiliates. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

## Trademarks

The Banzai Cloud name, the Banzai Cloud logo, and all Banzai Cloud trademarks and logos are registered trademarks of Cisco.
