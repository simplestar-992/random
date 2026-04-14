package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

var (
	intRange = flag.Int64("i", 0, "Random integer (0 = unlimited)")
	float    = flag.Bool("f", false, "Random float 0-1")
	choice   = flag.String("c", "", "Pick random from comma-separated list")
	hex      = flag.Int("x", 0, "Random hex string of length N")
	bin      = flag.Int("b", 0, "Random binary string of length N")
	uuid     = flag.Bool("u", false, "Random UUID")
)

func main() {
	flag.Parse()
	
	switch {
	case *intRange > 0:
		n, _ := rand.Int(rand.Reader, big.NewInt(*intRange))
		fmt.Println(n)
	case *float:
		f, _ := rand.Float64()
		fmt.Println(f)
	case *choice != "":
		parts := strings.Split(*choice, ",")
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(parts))))
		fmt.Println(strings.TrimSpace(parts[n.Int64()]))
	case *hex > 0:
		b := make([]byte, (*hex+1)/2)
		rand.Read(b)
		fmt.Printf("%x\n", b[:*hex/2+1])
	case *bin > 0:
		b := make([]byte, *bin)
		rand.Read(b)
		for i := range b {
			if b[i] < 128 { b[i] = '0' } else { b[i] = '1' }
		}
		fmt.Println(string(b))
	default:
		b := make([]byte, 16)
		rand.Read(b)
		fmt.Printf("%x-%x-%x-%x-%x\n", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	}
}
