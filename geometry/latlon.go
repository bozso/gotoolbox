package geometry

import (
    "fmt"

    "github.com/bozso/gotoolbox/errors"
    "github.com/bozso/gotoolbox/splitted"
)

type LatLon struct {
    Lat float64 `json:"lan"`
    Lon float64 `json:"lot"`
}

func (ll LatLon) ToPoint() (p Point2D) {
    p.X, p.Y = ll.Lon, ll.Lat
    return 
} 

func (ll LatLon) String() string {
    return fmt.Sprintf("%f,%f", ll.Lon, ll.Lat)
}

func (ll *LatLon) Set(s string) (err error) {
    const field errors.NotEmpty = "LatLon"
    if err = field.Check(s); err != nil {
        return
    }
    
    split, err := splitted.New(s, ",")
    if err != nil {
        return
    }
    
    ll.Lat, err = split.Float(0)
    if err != nil {
        return
    }

    ll.Lon, err = split.Float(1)
    if err != nil {
        return
    }
    
    return nil
}

