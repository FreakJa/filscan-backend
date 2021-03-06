package main

import (
	"context"
	"filscan_lotus/filscaner"
	"filscan_lotus/utils"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func TestFilscaner(t *testing.T) {
	filscaner := filscanor_instance()
	filscaner.ChainHeadTest()
}

func new_notify(ctx context.Context) (<-chan []int, error) {
	c := make(chan []int, 10)
	go func() {
		report_ticker := time.NewTicker(time.Second * time.Duration(3))
		for {
			select {
			case <-report_ticker.C:
				{
					var x = []int{rand.Int(), rand.Int()}
					c <- x
				}
			case <-ctx.Done():
				{
					utils.Printf("", "ctx.done()!!!!!")
					return
				}
			}

		}
	}()

	return c, nil
}

func truncateNaive(f float64, unit float64) float64 {
	return math.Trunc(f/unit) * unit
}

func to_xsize(power *big.Int, x *big.Int) float64 {
	fw := big.NewFloat(0)
	fw.SetString(power.String())

	ftbsize := big.NewFloat(0)
	ftbsize.SetString(x.String())

	v1, _ := fw.Float64()
	v2, _ := ftbsize.Float64()

	return truncateNaive(v1/v2, 0.01)
}

func TestX(t *testing.T) {
	powr := big.NewInt(0)
	powr.SetString("24704651886592", 10)
	fmt.Println(to_xsize(powr, filscaner.GB))
	return

	ctx, cancel := context.WithCancel(context.TODO())

	c, _ := new_notify(ctx)

	for data := range c {
		fmt.Printf("%v\n", data)
	}
	return

	after := time.After(time.Second * 10)

forlabel:
	for {
		select {
		case data, _ := <-c:
			{
				for v := range data {
					utils.Printf("", "data is:%d\n", v)
				}
			}
		case <-after:
			{
				cancel()
				break forlabel
			}
		}
	}
}
