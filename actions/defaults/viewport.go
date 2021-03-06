package defaults

import (
	"strconv"

	"github.com/driusan/de/actions"
	"github.com/driusan/de/demodel"
	"github.com/driusan/de/viewer"
)

func init() {
	// These should logically be in the viewport package,
	// but that would result in a cyclical import, so they're
	// here instead.
	actions.RegisterAction("TermWidth", TermWidth)
	actions.RegisterAction("WarnAlpha", WarnAlpha)
	actions.RegisterAction("LineNumbers", LineNumberMode)
	actions.RegisterAction("BackgroundMode", BackgroundMode)
}

func TermWidth(args string, buff *demodel.CharBuffer, v demodel.Viewport) {
	i, err := strconv.Atoi(args)
	if err != nil {
		buff.AppendTag("\nTermWidth: " + err.Error())
		return
	}

	v.SetOption("TermWidth", int(i))
}

func WarnAlpha(args string, buff *demodel.CharBuffer, v demodel.Viewport) {
	i, err := strconv.Atoi(args)
	if err != nil {
		return
	}
	if i >= 256 || i < 0 {
		buff.AppendTag("\nWarnAlpha: argument must be between 0 and 255")
		return
	}
	v.SetOption("WarnAlpha", uint8(i))
}

func LineNumberMode(args string, buff *demodel.CharBuffer, v demodel.Viewport) {
	switch args {
	case "absolute":
		v.SetOption("LineNumbers", viewer.AbsoluteLineNumbers)
	case "relative":
		v.SetOption("LineNumbers", viewer.RelativeLineNumbers)
	case "off":
		v.SetOption("LineNumbers", viewer.NoLineNumbers)
	default:
		v.SetOption("RotateLineNumbers", nil)
	}
}

func BackgroundMode(args string, buff *demodel.CharBuffer, v demodel.Viewport) {
	switch args {
	case "stable":
		v.SetOption("BackgroundMode", viewer.StableBackground)
	default:
		v.SetOption("BackgroundMode", viewer.DefaultBackground)
	}
}
