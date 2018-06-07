package unittest

import (
	"gopkg.in/check.v1"

	"github.com/dagdafrench/testing_osh_cli/core/clusters/probes"
)

func (s *OshinkoUnitTestSuite) TestNewHTTPGetProbe(c *check.C) {
	expectedPort := 8080
	probe := probes.NewHTTPGetProbe(expectedPort)
	c.Assert(probe.Handler.HTTPGet.Port.IntValue(), check.Equals, expectedPort)
}
