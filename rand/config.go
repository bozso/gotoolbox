package rand

import (
    "golang.org/x/exp/rand"
)

type Source rand.Source

func NewSource(seed uint64) (s Source) {
    s = Source(rand.NewSource(seed))
    return
}

type Config struct {
    Mean float64
    Std float64
}

func (c Config) ScaleInt(i int) (ri int) {
    return int(c.Std) * i + int(c.Mean)
}

func (c Config) ScaleFloat32(f float32) (rf float32) {
    return float32(c.Std) * f + float32(c.Mean)
}

func (c Config) ScaleFloat64(f float64) (rf float64) {
    return c.Std * f + c.Mean
}

type Rand struct {
    config Config
    *rand.Rand
}

func (c Config) New(s rand.Source) (r Rand) {
    r.config, r.Rand = c, rand.New(s)
    return
}

func (r Rand) Float32() (f float32) {
    return r.config.ScaleFloat32(r.Rand.Float32())
}

func (r Rand) Float64() (f float64) {
    return r.config.ScaleFloat64(r.Rand.Float64())
}

func (r Rand) Int() (i int) {
    return r.config.ScaleInt(r.Rand.Int())
}

func DefaultConfig() (c Config) {
    return defaultConfig
}

var defaultConfig = Config{
    Mean: 0.0,
    Std: 1.0,
}

func Default(s rand.Source) (r Rand) {
    return defaultConfig.New(s)
}
