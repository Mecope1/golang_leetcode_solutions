package kata

import(
 "strconv"
 "strings"
)

func CountBits(u_int uint) int {
 x := int64(u_int)

 num_bits := 0
 bin_x := strconv.FormatInt(x, 2)
 num_bits = strings.Count(bin_x, "1")

 return num_bits
}