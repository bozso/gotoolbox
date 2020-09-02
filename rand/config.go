package rand

import (
    "golang.org/x/exp/rand"
)

// type alias for the rand package's rand.Source type
type Source rand.Source

// wrapper for the rand package NewSource function
func NewSource(seed uint64) (s Source) {
    s = Source(rand.NewSource(seed))
    return
}

/*
 * Configuration for shifting the mean value and scaling the standard
 * deviation of random values.
 */
type Config struct {
    Mean float64
    Std float64
}

// scale an integer number
func (c Config) ScaleInt(i int) (ri int) {
    return int(c.Std) * i + int(c.Mean)
}

// scale a float32 number
func (c Config) ScaleFloat32(f float32) (rf float32) {
    return float32(c.Std) * f + float32(c.Mean)
}

// scale a float64 number
func (c Config) ScaleFloat64(f float64) (rf float64) {
    return c.Std * f + c.Mean
}

/*
 * Wrapper for the rand.Rand struct. Applies scaling defined by a Config
 * struct.
 */
type Rand struct {
    config Config
    *rand.Rand
}

// Create a new Rand struct with configuration defined by Config.
func (c Config) New(s rand.Source) (r Rand) {
    r.config, r.Rand = c, rand.New(s)
    return
}

// Get a random Float32.
func (r Rand) Float32() (f float32) {
    return r.config.ScaleFloat32(r.Rand.Float32())
}

// Get a random Float64.
func (r Rand) Float64() (f float64) {
    return r.config.ScaleFloat64(r.Rand.Float64())
}

// Get a random Int.
func (r Rand) Int() (i int) {
    return r.config.ScaleInt(r.Rand.Int())
}

// Returns a copy of the default configuration with zero mean and 1.0
// standard deviation.
func DefaultConfig() (c Config) {
    return defaultConfig
}

var defaultConfig = Config{
    Mean: 0.0,
    Std: 1.0,
}

// Create a new random number generator with the default configuration.
func Default(s rand.Source) (r Rand) {
    return defaultConfig.New(s)
}
