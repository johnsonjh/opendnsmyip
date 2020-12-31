# opendnsmyip

A Go package that returns the public-facing IPv4 address of the client
by querying the Cisco OpenDNS servers.

[![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/gridfinity/opendnsmyip/master/LICENSE)
[![GoVersion](https://img.shields.io/github/go-mod/go-version/gridfinity/opendnsmyip.svg)](https://github.com/gridfinity/opendnsmyip/blob/master/go.mod)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gridfinity/opendnsmyip)](https://pkg.go.dev/github.com/gridfinity/opendnsmyip)
[![GoReportCard](https://goreportcard.com/badge/github.com/gridfinity/opendnsmyip)](https://goreportcard.com/report/github.com/gridfinity/opendnsmyip)
[![GitHubRelease](https://img.shields.io/github/release/gridfinity/opendnsmyip.svg)](https://github.com/gridfinity/opendnsmyip/releases/)
[![LocCount](https://img.shields.io/tokei/lines/github/gridfinity/opendnsmyip.svg)](https://github.com/XAMPPRocky/tokei)
[![GitHubCodeSize](https://img.shields.io/github/languages/code-size/gridfinity/opendnsmyip.svg)](https://github.com/gridfinity/opendnsmyip)
[![CoverageStatus](https://coveralls.io/repos/github/gridfinity/opendnsmyip/badge.svg?branch=master)](https://coveralls.io/github/gridfinity/opendnsmyip?branch=master)
[![CodacyBadge](https://api.codacy.com/project/badge/Grade/c756d556a38842a5b82265e5f1bebcc1)](https://app.codacy.com/gh/gridfinity/opendnsmyip?utm_source=github.com&utm_medium=referral&utm_content=gridfinity/opendnsmyip&utm_campaign=Badge_Grade)
[![CodebeatBadge](https://codebeat.co/badges/40937670-1418-4a95-92a0-e4c144aef861)](https://codebeat.co/projects/github-com-gridfinity-opendnsmyip-master-4ee973c4-e210-4be1-bb7a-45b89aa0cb70)
[![CodeClimateMaintainability](https://api.codeclimate.com/v1/badges/d8e0a5a40404d2153688/maintainability)](https://codeclimate.com/github/gridfinity/opendnsmyip/maintainability)
[![TickgitTODOs](https://img.shields.io/endpoint?url=https://api.tickgit.com/badge?repo=github.com/gridfinity/opendnsmyip)](https://www.tickgit.com/browse?repo=github.com/gridfinity/opendnsmyip)
[![DeepSource](https://deepsource.io/gh/gridfinity/opendnsmyip.svg/?label=active+issues)](https://deepsource.io/gh/gridfinity/opendnsmyip/?ref=repository-badge)
[![DeepScanGrade](https://deepscan.io/api/teams/12184/projects/15166/branches/299578/badge/grade.svg)](https://deepscan.io/dashboard#view=project&tid=12184&pid=15166&bid=299578)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fjohnsonjh%2Fopendnsmyip.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fjohnsonjh%2Fopendnsmyip?ref=badge_shield)

## Availability

### Go Modules

- [go.gridfinity.dev](https://go.gridfinity.dev/opendnsmyip)
- [go.gridfinity.com](https://go.gridfinity.com/)

### Source Code

- [Gridfinity GitLab](https://gitlab.gridfinity.com/go/opendnsmyip)
- [SourceHut Git](https://sr.ht/~trn/opendnsmyip/)
- [GitHub](https://github.com/gridfinity/opendnsmyip)

## Issue Tracking

- [Gridfinity GitLab](https://gitlab.gridfinity.com/go/opendnsmyip/-/issues)

## Security Policy

- [Security Policy and Vulnerability Reporting](https://gitlab.gridfinity.com/go/opendnsmyip/-/blob/master/SECURITY.md)

## Original Authors

- [James Polera](https://github.com/polera/publicip) \<[james@uncryptic.com](mailto:james@uncryptic.com)\>
- [Dmitriy Kuznetsov](https://github.com/Dikman/publicip) \<[dmitriy.v.kuznetsov@gmail.com](mailto:dmitriy.v.kuznetsov@gmail.com)\>

## Coverage Reports

- [GoCov](https://pktdist.gridfinity.com/coverage/opendnsmyip/)
- [Coveralls](https://coveralls.io/github/gridfinity/opendnsmyip)

## Credits

This package was inspired by:

- [public-ip (nodejs)](https://github.com/sindresorhus/public-ip/blob/master/index.js)
- [OpenDNS::MyIP (Perl 5)](https://metacpan.org/pod/OpenDNS::MyIP)

## License

- [MIT License](https://tldrlegal.com/license/mit-license)
- [![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fjohnsonjh%2Fopendnsmyip.svg?type=small)](https://app.fossa.com/projects/git%2Bgithub.com%2Fjohnsonjh%2Fopendnsmyip?ref=badge_small)

## Usage

```go
package main

import (
        "fmt"

        myip "go.gridfinity.dev/opendnsmyip"
)

func main() {
        myIpAddr, err := myip.GetIP()
        if err != nil {
                fmt.Errorf(
                           "Error getting IPv4 address: %v",
                           err,
                          )
        } else {
                fmt.Printf(
                           "Public IPv4 address is: %s",
                           myIpAddr,
                          )
        }
}
```
