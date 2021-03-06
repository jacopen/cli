package user

import (
	"github.com/cloudfoundry/cli/cf/i18n"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
	"path/filepath"
)

var T goi18n.TranslateFunc

func init() {
	T = i18n.Init(filepath.Join("cf", "commands", "user"), i18n.GetResourcesPath())
}
