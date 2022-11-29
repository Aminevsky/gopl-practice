package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifyUser(t *testing.T) {
	saved := notifyUser
	defer func() { notifyUser = saved }()

	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	const user = "joe@example.com"
	usage[user] = 980000000
	CheckQuota(user)

	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called, %s, %s", notifiedUser, notifiedMsg)
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}

	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+"want substring %q", notifiedMsg, wantSubstring)
	}
}