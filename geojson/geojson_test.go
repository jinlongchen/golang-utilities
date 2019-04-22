package geojson

import (
	"fmt"
	"testing"
)

func TestPreload(t *testing.T) {
	Preload()
}
func TestFindGeoFeatureByAddCode(t *testing.T) {
	Preload()
	x := FindGeoFeatureByAddCode(510000)
	fmt.Println(x)
}