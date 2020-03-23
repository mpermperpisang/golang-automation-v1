package androidaction

import (
	"github.com/golang-automation/features/helper"
)

/*SendKeysByXpath input element by Xpath selector*/
func SendKeysByXpath(locator string, text string) error {
	element := Device.FindByXPath(locator).SendKeys(text)
	helper.LogPanicln(element)

	return nil
}

/*SendKeysByButton input element by button*/
func SendKeysByButton(locator string, text string) error {
	element := Device.FindByButton(locator).SendKeys(text)
	helper.LogPanicln(element)

	return nil
}

/*SendKeysByClass input element by class*/
func SendKeysByClass(locator string, text string) error {
	element := Device.FindByClass(locator).SendKeys(text)
	helper.LogPanicln(element)

	return nil
}

/*SendKeysByID input element by ID*/
func SendKeysByID(locator string, text string) error {
	element := Device.FindByID(locator).SendKeys(text)
	helper.LogPanicln(element)

	return nil
}

/*SendKeysByLabel input element by label*/
func SendKeysByLabel(locator string, text string) error {
	element := Device.FindByLabel(locator).SendKeys(text)
	helper.LogPanicln(element)

	return nil
}

/*SendKeysByLink input element by link*/
func SendKeysByLink(locator string, text string) error {
	element := Device.FindByLink(locator).SendKeys(text)
	helper.LogPanicln(element)

	return nil
}

/*SendKeysByName input element by class name*/
func SendKeysByName(locator string, text string) error {
	element := Device.FindByName(locator).SendKeys(text)
	helper.LogPanicln(element)

	return nil
}

/*SendKeysByText input element by text in xpath*/
func SendKeysByText(locator string, text string) error {
	element := Device.FindByXPath("//*[contains(@text, '" + locator + "')]").SendKeys(text)
	helper.LogPanicln(element)

	return nil
}
