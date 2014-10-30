package common

import (
	. "framework/api/common/types"
	"testing"
)

func TestReadConfig(t *testing.T) {
	var cfg LaserConfig

	newcfg := ReadConfig(Lasercfg, cfg)
	Logging(newcfg)
	Logging(cfg)
}
