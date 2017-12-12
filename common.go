package vld

import (
	"errors"
	"reflect"

	"fmt"
)

// vld:maxint64
func MaxInt64(limit, value int64, name string) ValidationFunc {
	return func() error {
		if value > limit {
			return fmt.Errorf("vld: %s: %d greater than limit %d", name, value, limit)
		}
		return nil
	}
}

// vld:minint64
func MinInt64(limit, value int64, name string) ValidationFunc {
	return func() error {
		if value < limit {
			return fmt.Errorf("vld: %s: %d less than limit %d", name, value, limit)
		}
		return nil
	}
}

// vld:maxint
func MaxInt(limit, value int, name string) ValidationFunc {
	return func() error {
		if value > limit {
			return fmt.Errorf("vld: %s: %d greater than limit %d", name, value, limit)
		}
		return nil
	}
}

// vld:minint
func MinInt(limit, value int, name string) ValidationFunc {
	return func() error {
		if value < limit {
			return fmt.Errorf("vld: %s: %d less than limit %d", name, value, limit)
		}
		return nil
	}
}

// vld:maxfloat32
func MaxFloat32(limit, value float32, name string) ValidationFunc {
	return func() error {
		if value > limit {
			return fmt.Errorf("vld: %s: %f greater than limit %f", name, value, limit)
		}
		return nil
	}
}

// vld:minfloat32
func MinFloat32(limit, value float32, name string) ValidationFunc {
	return func() error {
		if value < limit {
			return fmt.Errorf("vld: %s: %f less than limit %f", name, value, limit)
		}
		return nil
	}
}

// vld:maxfloat64
func MaxFloat64(limit, value float32, name string) ValidationFunc {
	return func() error {
		if value > limit {
			return fmt.Errorf("vld: %s: %f greater than limit %f", name, value, limit)
		}
		return nil
	}
}

// vld:minfloat64
func MinFloat64(limit, value float32, name string) ValidationFunc {
	return func() error {
		if value < limit {
			return fmt.Errorf("vld: %s: %f less than limit %f", name, value, limit)
		}
		return nil
	}
}

// vld:maxlen
func MaxLength(length int, value, name string) ValidationFunc {
	return func() error {
		if len(value) > length {
			return fmt.Errorf("vld: Too long: %s: %s", name, value)
		}
		return nil
	}
}

// vld:minlen
func MinLength(length int, value, name string) ValidationFunc {
	return func() error {
		if len(value) < length {
			return fmt.Errorf("vld: Too short: %s: %s", name, value)
		}
		return nil
	}
}

// vld:req
func Required(value interface{}, name string) ValidationFunc {
	return func() error {
		if reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface()) {
			return fmt.Errorf("vld: Required: %s", name)
		}
		return nil
	}
}

func Message(f ValidationFunc, message string) ValidationFunc {
	return func() error {
		err := f()
		if err != nil {
			return errors.New(message)
		}
		return nil
	}
}
