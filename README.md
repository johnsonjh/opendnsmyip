# opendnsmyip

A Go package that returns the public-facing IPv4 address of the client
by querying the Cisco OpenDNS servers.

[![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/johnsonjh/opendnsmyip/master/LICENSE)
[![GoVersion](https://img.shields.io/github/go-mod/go-version/johnsonjh/opendnsmyip.svg)](https://github.com/johnsonjh/opendnsmyip/blob/master/go.mod)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/johnsonjh/opendnsmyip)](https://pkg.go.dev/github.com/johnsonjh/opendnsmyip)
[![GoReportCard](https://goreportcard.com/badge/github.com/johnsonjh/opendnsmyip)](https://goreportcard.com/report/github.com/johnsonjh/opendnsmyip)
[![GitHubRelease](https://img.shields.io/github/release/johnsonjh/opendnsmyip.svg)](https://github.com/johnsonjh/opendnsmyip/releases/)
[![GitHubTag](https://img.shields.io/github/tag/johnsonjh/opendnsmyip.svg)](https://github.com/johnsonjh/opendnsmyip/tags/)
[![LocCount](https://img.shields.io/tokei/lines/github/johnsonjh/opendnsmyip.svg)](https://github.com/XAMPPRocky/tokei)
[![GitHubCodeSize](https://img.shields.io/github/languages/code-size/johnsonjh/opendnsmyip.svg)](https://github.com/johnsonjh/opendnsmyip)
[![CoverageStatus](https://coveralls.io/repos/github/johnsonjh/opendnsmyip/badge.svg?branch=master)](https://coveralls.io/github/johnsonjh/opendnsmyip?branch=master)
[![CodacyBadge](https://api.codacy.com/project/badge/Grade/c756d556a38842a5b82265e5f1bebcc1)](https://app.codacy.com/gh/johnsonjh/opendnsmyip?utm_source=github.com&utm_medium=referral&utm_content=johnsonjh/opendnsmyip&utm_campaign=Badge_Grade)
[![CodeBeat](https://codebeat.co/badges/f148c0a4-604b-449e-b450-02fb5074a131)](https://codebeat.co/projects/github-com-johnsonjh-opendnsmyip-master)
[![CodeClimateMaintainability](https://api.codeclimate.com/v1/badges/d8e0a5a40404d2153688/maintainability)](https://codeclimate.com/github/johnsonjh/opendnsmyip/maintainability)
[![TickgitTODOs](https://img.shields.io/endpoint?url=https://api.tickgit.com/badge?repo=github.com/johnsonjh/opendnsmyip)](https://www.tickgit.com/browse?repo=github.com/johnsonjh/opendnsmyip)
[![DeepSource](https://deepsource.io/gh/johnsonjh/opendnsmyip.svg/?label=active+issues)](https://deepsource.io/gh/johnsonjh/opendnsmyip/?ref=repository-badge)
[![FOSSAStatus](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fjohnsonjh%2Fopendnsmyip.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fjohnsonjh%2Fopendnsmyip?ref=badge_shield)

## Availability
- [Gridfinity GitLab](https://gitlab.gridfinity.com/jeff/go-opendns-myip)
- [GitHub](https://github.com/johnsonjh/opendnsmyip)

## Security
- [Security Policy](https://github.com/johnsonjh/opendnsmyip/blob/master/SECURITY.md)

## Original Authors

- [James Polera](https://github.com/polera/publicip) <mailto:james@uncryptic.com>
- [Dmitriy Kuznetsov](https://github.com/Dikman/publicip) <mailto:dmitriy.v.kuznetsov@gmail.com>

## Coverage Reports
- [GoCov](https://pktdist.gridfinity.com/coverage/opendnsmyip/)
- [Coveralls](https://coveralls.io/github/johnsonjh/opendnsmyip)

## Credits

This package was inspired by:
- [public-ip (nodejs)](https://github.com/sindresorhus/public-ip/blob/master/index.js)
- [OpenDNS::MyIP (Perl 5)](https://metacpan.org/pod/OpenDNS::MyIP)

## License
- [MIT License](https://tldrlegal.com/license/mit-license)
- [![FOSSASmallStatus](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fjohnsonjh%2Fopendnsmyip.svg?type=small)](https://app.fossa.com/projects/git%2Bgithub.com%2Fjohnsonjh%2Fopendnsmyip?ref=badge_small)

## Usage
```go
package main

import (
 "fmt"
 myip "github.com/johnsonjh/opendnsmyip"
)

func main() {
 myIpAddr, err := myip.GetIP()
 if err != nil {
  fmt.Errorf("Error getting IPv4 address: %v", err)
 } else {
  fmt.Printf("Public IPv4 address is: %s", myIpAddr)
 }
}
```
