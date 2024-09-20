/*
 * Copyright (c) 2020. Jinlong Chen.
 */

package geo

import (
    "fmt"
    "testing"
)

func TestGCJ02toWGS84(t *testing.T) {
    lon, lat := GCJ02toWGS84(104.042389, 30.47767)
    fmt.Printf("%0.6f %0.6f\n", lon, lat)
}
