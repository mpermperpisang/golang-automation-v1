package supports

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/cucumber/godog"

	"github.com/golang-automation-v1/features/helper"
	"github.com/golang-automation-v1/features/helper/data"
	"github.com/golang-automation-v1/features/helper/formats"
	"github.com/golang-automation-v1/features/helper/messages"
	appsbase "github.com/golang-automation-v1/features/supports/base/apps"
	webbase "github.com/golang-automation-v1/features/supports/base/web"
	supports "github.com/golang-automation-v1/features/supports/drivers"
	"github.com/golang-automation-v1/features/supports/logs"
	"github.com/golang-automation-v1/features/supports/screenshots"
	"github.com/golang-automation-v1/features/supports/structs"
)

var (
	testCase               structs.ScenarioDetail
	scenarioTags, platform string
)

func InitializeTestSuite(s *godog.TestSuiteContext) {
	s.BeforeSuite(func() {
		recreateLogs()
		recreateReport()
		recreateSS()
		recreateXray()
	})
}

func InitializeScenario(s *godog.ScenarioContext) {
	s.Before(func(ctx context.Context, scenario *godog.Scenario) (context.Context, error) {
		scenarioDetail(scenario)

		for i := 0; i < len(testCase.Tags); i++ {
			scenarioTags += " " + testCase.Tags[i].Name
		}

		platformCheck(scenarioTags)

		return ctx, nil
	})

	s.After(func(ctx context.Context, scenario *godog.Scenario, log error) (context.Context, error) {
		scenarioTags = ""

		scenarioDetail(scenario)

		if log != nil {
			createSS()
			createLog(log)
		}

		return ctx, nil
	})
}

func scenarioDetail(scenario *godog.Scenario) {
	data, err := json.Marshal(scenario)
	helper.LogPanicln(err)

	json.Unmarshal(data, &testCase)
}

func platformCheck(tags string) {
	if strings.Contains(tags, data.API) {
		platform = data.API
	} else if strings.Contains(tags, data.DWEB) {
		platform = data.DWEB

		webStart(platform)
	} else if strings.Contains(tags, data.MWEB) {
		platform = data.MWEB

		webStart(platform)
	} else if strings.Contains(tags, data.ANDROID) {
		platform = data.ANDROID

		appsStart(platform)
	} else if strings.Contains(tags, data.IOS) {
		platform = data.IOS

		appsStart(platform)
	} else {
		helper.LogPanicln(messages.PlatformList())
	}
}

func webStart(platform string) {
	supports.WebCapabilities()
	webbase.OpenWebURL(platform, "/")
}

func appsStart(platform string) {
	switch platform {
	case data.ANDROID:
		supports.AndroidCapabilities()
	case data.IOS:
		supports.IOSCapabilities()
	default:
		helper.LogPanicln(messages.PlatformList())
	}

	appsbase.OpenApps(platform)
}

func createSS() {
	helper.SetFilename(data.SS, platform, testCase.Name)
	screenshots.SSWeb(platform)
	screenshots.SSAndroid()
	screenshots.SSIOS()
}

func createLog(log error) {
	helper.SetFilename(data.LOGS, platform, testCase.Name)
	logs.LogWeb(platform, log)
	logs.LogAndroid(log)
	logs.LogIOS(log)
}

func recreateLogs() {
	helper.RemoveAllContent(formats.TestPath(data.LOGS, data.WEB))
	helper.RemoveAllContent(formats.TestPath(data.LOGS, data.APPS))
	helper.DirectoryCheck(helper.GetPWD() + formats.TestPath(data.LOGS, data.WEB))
	helper.DirectoryCheck(helper.GetPWD() + formats.TestPath(data.LOGS, data.APPS))
}

func recreateSS() {
	helper.RemoveAllContent(formats.TestPath(data.SS, data.WEB))
	helper.RemoveAllContent(formats.TestPath(data.SS, data.APPS))
	helper.DirectoryCheck(helper.GetPWD() + formats.TestPath(data.SS, data.WEB))
	helper.DirectoryCheck(helper.GetPWD() + formats.TestPath(data.SS, data.APPS))
}

func recreateReport() {
	helper.RemoveAllContent(formats.TestPath(data.REPORT, data.WEB))
	helper.RemoveAllContent(formats.TestPath(data.REPORT, data.APPS))
	helper.DirectoryCheck(helper.GetPWD() + formats.TestPath(data.REPORT, data.WEB))
	helper.DirectoryCheck(helper.GetPWD() + formats.TestPath(data.REPORT, data.APPS))
}

func recreateXray() {
	helper.RemoveAllContent(formats.TestPath(data.XRAY, data.WEB))
	helper.RemoveAllContent(formats.TestPath(data.XRAY, data.APPS))
	helper.DirectoryCheck(helper.GetPWD() + formats.TestPath(data.XRAY, data.WEB))
	helper.DirectoryCheck(helper.GetPWD() + formats.TestPath(data.XRAY, data.APPS))
}
