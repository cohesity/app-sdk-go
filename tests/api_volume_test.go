package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Volume struct {
	suite.Suite
	athena AthenaStruct
}

func (vol *Volume) SetupTest() {
	vol.athena.setmeUp()
}
func (vol *Volume) Test01CreateVolume() {

}

func (vol *Volume) Test02GetVolume() {

}

func (vol *Volume) Test03DeleteVolume() {

}

func (vol *Volume) Test04DeleteNonExistantVolume() {

}

func (vol * Volume) Test05BogusVolume() {

}

func TestVolume(t *testing.T) {
	// This is what actually runs our suite
	a := new(Volume)
	suite.Run(t, a)
}
