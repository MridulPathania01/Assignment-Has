package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
)

// Point represents a (x, y) pair
type Point struct {
	X int64
	Y *big.Int
}

// Root represents each root entry in the JSON
type Root struct {
	Base  string `json:"base"`
	Value string `json:"value"`
}

type TestCase struct {
	Keys struct {
		N int `json:"n"`
		K int `json:"k"`
	} `json:"keys"`
	Data map[string]Root
}

func readTestCase(filename string) (*TestCase, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var fullMap map[string]json.RawMessage
	if err := json.Unmarshal(file, &fullMap); err != nil {
		return nil, err
	}

	var tc TestCase
	tc.Data = make(map[string]Root)

	for key, value := range fullMap {
		if key == "keys" {
			if err := json.Unmarshal(value, &tc.Keys); err != nil {
				return nil, err
			}
		} else {
			var r Root
			if err := json.Unmarshal(value, &r); err != nil {
				return nil, err
			}
			tc.Data[key] = r
		}
	}
	return &tc, nil
}

// Convert and decode points
func decodePoints(data map[string]Root) []Point {
	points := []Point{}
	for k, v := range data {
		x, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			log.Fatalf("Invalid key: %v", k)
		}
		base, err := strconv.Atoi(v.Base)
		if err != nil {
			log.Fatalf("Invalid base: %v", v.Base)
		}
		y := new(big.Int)
		y.SetString(v.Value, base)
		points = append(points, Point{X: x, Y: y})
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})
	return points
}

// Lagrange interpolation at x = 0 to find constant term
func lagrangeInterpolationAtZero(points []Point, k int) *big.Int {
	result := big.NewInt(0)

	for j := 0; j < k; j++ {
		numerator := big.NewInt(1)
		denominator := big.NewInt(1)

		for m := 0; m < k; m++ {
			if m != j {
				xm := big.NewInt(points[m].X)
				xj := big.NewInt(points[j].X)

				num := new(big.Int).Neg(xm)       // -xm
				den := new(big.Int).Sub(xj, xm)   // xj - xm
				numerator.Mul(numerator, num)     // numerator *= -xm
				denominator.Mul(denominator, den) // denominator *= (xj - xm)
			}
		}

		// term = yj * numerator / denominator
		yj := new(big.Int).Set(points[j].Y)
		term := new(big.Int).Mul(yj, numerator)
		term.Div(term, denominator)

		result.Add(result, term)
	}

	return result
}

func solve(filename string) *big.Int {
	tc, err := readTestCase(filename)
	if err != nil {
		log.Fatalf("Failed to read test case: %v", err)
	}
	points := decodePoints(tc.Data)
	return lagrangeInterpolationAtZero(points, tc.Keys.K)
}

func main() {
	secret1 := solve("testcase1.json")
	secret2 := solve("testcase2.json")

	fmt.Println("Secret from Test Case 1:", secret1)
	fmt.Println("Secret from Test Case 2:", secret2)
}
