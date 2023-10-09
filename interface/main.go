package main

import (
	"fmt"
	// 	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (i IPAddr) String() string {
	// ans := ""
	// for _, n := range i {
	// 	ans = ans + strconv.Itoa(int(n)) + "."
	// }
	// return fmt.Sprintf("%q", ans[:len(ans)-1])
	return fmt.Sprintf("\"%d.%d.%d.%d\"", i[0], i[1], i[2], i[3])

}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	i := float64(1)

	for j := 0; j < 10; j++ {
		i -= (i*i - x) / (2 * i)
	}
	return i, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
