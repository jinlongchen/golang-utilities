package log

import "testing"

func TestDumpKeyValue(t *testing.T) {
    Config("", LevelDebug, true, "", LogFormatJSON, 0, 0, 1)
    Infof("err: %s", "ok")
}
