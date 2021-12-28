package upgrader

import (
	"context"
	"fmt"
	"time"

	"github.com/abahmed/kwatch/constant"
	"github.com/abahmed/kwatch/util"
	"github.com/google/go-github/v41/github"
)

func CheckUpdates() {
	ticker := time.NewTicker(24 * time.Hour)

	for range ticker.C {
		client := github.NewClient(nil)
		r, _, err := client.Repositories.GetLatestRelease(context.TODO(), "abahmed", "kwatch")
		if err == nil {
			if constant.Version != *r.TagName {
				notifyNewVersion(*r.TagName)
			}
		}
	}
}

func notifyNewVersion(version string) {
	providers := util.GetProviders()
	for _, p := range providers {
		p.SendMessage(fmt.Sprintf(constant.KwatchUpdateMsg, version))
	}
}
