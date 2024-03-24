package godi

import "github.com/lvyahui8/goenum"

type InjectTag struct {
	goenum.Enum
}

var (
	TagDefault = goenum.NewEnum[InjectTag]("Default")
	TagPrivate = goenum.NewEnum[InjectTag]("Private")
	TagNamed   = goenum.NewEnum[InjectTag]("Named")
)
