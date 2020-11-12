#publicip

This package returns the public facing IP address of the calling client (a la https://icanhazip.com, but from Go!)

 [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg)](http://godoc.org/github.com/dikman/publicip) [![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/dikman/publicip/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/dikman/publicip)](https://goreportcard.com/report/github.com/dikman/publicip) [![codecov](https://codecov.io/gh/dikman/publicip/branch/master/graph/badge.svg)](https://codecov.io/gh/dikman/publicip)

Changes
==
Removed output to a log directly from this package (Dmitriy Kuznetsov <dmitriy.v.kuznetsov@gmail.com>).

Author of the original package
==
James Polera <james@uncryptic.com>

Dependencies
==
publicip uses Glide for dependency management.  After cloning this package, run:
```bash
glide up
```

Credits
==
This package was inspired by both:

[public-ip (nodejs)](https://github.com/sindresorhus/public-ip/blob/master/index.js)

[OpenDNS::MyIP (Perl)](https://metacpan.org/pod/OpenDNS::MyIP)

Example
==
```go
package main

import (
  "fmt"
  "github.com/dikman/publicip"
)

func main() {

  myIpAddr, err := publicip.GetIP()
  if err != nil {
    fmt.Printf("Error getting IP address: %s\n", err)
  } else {
    fmt.Printf("Public IP address is: %s", myIpAddr)
  }

}

```


