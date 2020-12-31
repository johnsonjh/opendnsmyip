// Copyright 2020 Jeffrey H. Johnson.
// Copyright 2020 Gridfinity, LLC.
// Copyright 2019 The Go Authors.
// All rights reserved.
// Use of this source code is governed by the BSD-style
// license that can be found in the LICENSE file.

package opendnsmyip_test

import (
	"testing"

	u "go.gridfinity.dev/leaktestfe"
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
