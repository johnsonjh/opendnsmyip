// Copyright © 2021 Jeffrey H. Johnson. <trnsz@pobox.com>
// Copyright © 2021 Gridfinity, LLC.
// Copyright © 2020 The Go Authors.
//
// All rights reserved.
//
// Use of this source code is governed by the BSD-style
// license that can be found in the LICENSE file.

package opendnsmyip_test

import (
	"testing"

	u "github.com/johnsonjh/leaktestfe"
)

func TestLeakVerifyNoneDisabled(
	t *testing.T,
) {
	u.Leakplug(
		t,
	)
}

func TestLeakVerifyNoneEnabled(
	t *testing.T,
) {
	u.Leakplug(
		t,
	)
}
