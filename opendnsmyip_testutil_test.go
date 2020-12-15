// Copyright 2020 Gridfinity, LLC.
// Copyright 2019 The Go Authors.
// All rights reserved.
// Use of this source code is governed by the BSD-style
// license that can be found in the LICENSE file.

package opendnsmyip_test

import (
	"fmt"
	"testing"

	u "github.com/johnsonjh/opendnsmyip/testutil"
)

func TestLeakVerifyNoneDisabled(
	t *testing.T,
) {
	err := u.LeakVerifyNone(
		t,
	)
	if err != nil {
		t.Fatal(
			fmt.Sprintf(
				"\nopendnsmyip_testutil_test.TestLeakVerifyNoneDisabled.LeakVerifyNone FAILURE:\n	%v",
				err,
			),
		)
	}
}

func TestLeakVerifyNoneEnabled(
	t *testing.T,
) {
	err := u.LeakVerifyNone(
		t,
	)
	if err != nil {
		t.Fatal(
			fmt.Sprintf(
				"\nopendnsmyip_testutil_test.TestLeakVerifyNoneEnabled.LeakVerifyNone FAILURE:\n	%v",
				err,
			),
		)
	}
}
