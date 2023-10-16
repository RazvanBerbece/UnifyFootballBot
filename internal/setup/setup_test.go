package setup

import "testing"

func TestSetup(t *testing.T) {
	got := Setup("../../.env")
	if len(got.AppId) == 0 {
		t.Error("AppId variable after setup bot not initialised. len(AppId)=", len(got.AppId))
	}
}
