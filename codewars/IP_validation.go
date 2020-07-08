package kata

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {

	is_valid := true
	octet_str_arr := strings.Split(ip, ".")

	if len(octet_str_arr) != 4 {
		is_valid = false
		return is_valid
	}

	for _, val := range octet_str_arr {
		if val[0:1] == "0" && len(val) > 1 {
			is_valid = false
		}

		i_val, err := strconv.Atoi(val)

		if err != nil {
			is_valid = false
		} else if 0 > i_val || i_val > 255 {
			is_valid = false
		}
	}

	return is_valid
}