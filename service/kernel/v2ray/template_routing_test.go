package v2ray

import (
	"testing"

	"github.com/v2rayA/v2rayA/db/configure"
)

func TestAppendDokodemoTProxyListenAddress(t *testing.T) {
	for _, tt := range []struct {
		mode configure.TransparentType
		want string
	}{
		{configure.TransparentRedirect, "0.0.0.0"},
		{configure.TransparentTproxy, "127.0.0.1"},
	} {
		t.Run(string(tt.mode), func(t *testing.T) {
			tmpl := &Template{}
			tmpl.AppendDokodemoTProxy(string(tt.mode), 52345, "transparent")
			if len(tmpl.Inbounds) != 1 {
				t.Fatalf("got %d inbounds, want 1", len(tmpl.Inbounds))
			}
			ib := tmpl.Inbounds[0]
			if ib.Tag != "transparent" || ib.Port != 52345 || ib.Protocol != "dokodemo-door" {
				t.Fatalf("unexpected transparent inbound: %+v", ib)
			}
			if ib.Listen != tt.want {
				t.Fatalf("listen = %q, want %q", ib.Listen, tt.want)
			}
			if !ib.Settings.FollowRedirect || *ib.StreamSettings.Sockopt.Tproxy != string(tt.mode) {
				t.Fatalf("transparent mode settings changed: %+v", ib)
			}
		})
	}
}
