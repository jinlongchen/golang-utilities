package log

import "testing"

func TestDumpKeyValue(t *testing.T) {
    Config("", LevelDebug, true, "", 0, 0, 0)
    Infof("err: %s", "ok")
}
