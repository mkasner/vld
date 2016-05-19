package ext

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/mkasner/vld"
)

// vld:email
func Email(value, name string) vld.ValidationFunc {
	return func() error {
		if !govalidator.IsEmail(value) && value != "" {
			return fmt.Errorf("vld: Not regular email: %s: %s", name, value)
		}
		return nil
	}
}

// vld:ip
func Ip(value, name string) vld.ValidationFunc {
	return func() error {
		if !govalidator.IsIP(value) && value != "" {
			return fmt.Errorf("vld: Not regular ip: %s: %s", name, value)
		}
		return nil
	}
}

// vld:mac
func Mac(value, name string) vld.ValidationFunc {
	return func() error {
		if !govalidator.IsMAC(value) && value != "" {
			return fmt.Errorf("vld: Not regular mac address: %s: %s", name, value)
		}
		return nil
	}
}
