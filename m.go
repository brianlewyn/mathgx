package mathgx

import (
	"fmt"
	"strconv"
	"strings"
)

// ! My Custom Functions.

// Check if there is the same length of parentheses.
func parentheses(gx string) error {
	opn, cls := strings.Count(gx, "("), strings.Count(gx, ")")
	if opn != cls {
		if opn > cls {
			text := strconv.Itoa(opn - cls)
			return myErrGo("Necessary → " + text + " close parentheses.")
		} else {
			text := strconv.Itoa(cls - opn)
			return myErrGo("Necessary → " + text + " open parentheses.")
		}
	}
	return nil
}

// Split right where the parentheses meet.
func splitIndex(gx string) []string {
	sgx, n := strings.Split(gx, ""), strings.Count(gx, "(")
	opn, cls := make([]int, n), make([]int, n)
	var t1, t2, t3 int

	// Located index.
	for i, item := range sgx {
		if item == "(" {
			opn[t1] = i + 1
			t1++
		}
		if item == ")" {
			cls[t2] = i
			t2++
		}
	}

	// Split by index.
	sMain := make([]string, n)
	for _, p := range opn {
		for _, c := range cls {
			if c > p {
				sMain[t3] = gx[p:c]
				t3++
				break
			}
		}
	}

	return sMain
}

// Getting the pair of k & n and storing in different sets.
func getKN(x string, sgx []string) (sq1, sq2 []float64) {
	for i := range sgx {
		before, after, _ := strings.Cut(sgx[i], x+"^")
		kF, _ := strconv.ParseFloat(before, 64)
		nF, _ := strconv.ParseFloat(after, 64)
		sq1, sq2 = append(sq1, kF), append(sq2, nF)
	}
	sq1 = highlow(sq1)
	return
}

// Add or Multiply between two sets.
func twoSets(AM string, sA, sB []float64) (c []float64) {
	for _, a := range sA {
		for _, b := range sB {
			if AM == "A" {
				c = append(c, a+b)
			}
			if AM == "M" {
				c = append(c, a*b)
			}
		}
	}
	return
}

// Add or Multiply between two or more sets.
func sets(AM string, s ...[]float64) []float64 {
	n := len(s) - 1
	for i, j := 0, 1; i < n; i, j = i+1, j+1 {
		s[j] = twoSets(AM, s[i], s[j])
	}
	return s[n]
}

// Add and multiply all the kxn of a section.
func simplify_M(x string, sMain []string) string {
	// Apply the addition method.
	for i, gx := range sMain {
		sMain[i], _ = A(x, gx, false)
	}

	// Obtaining the two largest set of k & n.
	ssq1, ssq2 := [][]float64{}, [][]float64{}
	for _, gx := range sMain {
		sgx := strings.Split(gx, " ")
		sq1, sq2 := getKN(x, sgx)
		ssq1, ssq2 = append(ssq1, sq1), append(ssq2, sq2)
	}

	gz, sK, sN := "", sets("M", ssq1...), sets("A", ssq2...)
	for i, n := range sN {
		kStr, nStr := fmt.Sprint(sK[i]), fmt.Sprint(n)
		if i == 0 {
			gz += kStr + x + "^" + nStr
		} else {
			temp := kStr + x + "^" + nStr
			if strings.HasPrefix(temp, "-") {
				gz += " " + temp
			} else {
				gz += " +" + temp
			}
		}
	}

	return gz
}

// ! Multiply: Multiplication of mathematical functions.
func M(x, gx string, simple bool) (string, error) {

	err := exists(x, gx)
	if err != nil {
		return "", err
	}

	err = correctly(x, gx, "()")
	if err != nil {
		return "", err
	}

	err = parentheses(gx)
	if err != nil {
		return "", err
	}

	sMain := splitIndex(gx)
	gz := simplify_M(x, sMain)
	gz, _ = A(x, gz, simple)

	return gz, nil
}
