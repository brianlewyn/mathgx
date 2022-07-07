package mathgx

import (
	"fmt"
	"strconv"
	"strings"
)

// ! My Custom Functions.

// Check if there is any empty input field.
func exists(x, gx string) error {
	if x == "" && gx == "" {
		return myErrGo("Fill in both input fields.")
	} else if x == "" && gx != "" {
		return myErrGo("Fill in the field f(a literal).")
	} else if x != "" && gx == "" {
		return myErrGo("Fill in the fild = (a math func).")
	}
	return nil
}

// Verify if the math func was written correctly.
func correctly(x, gx, sign string) error {
	slice := strings.Split("0123456789 .-+^"+x+sign, "")
	for _, item := range slice {
		gx = strings.ReplaceAll(gx, item, "")
	}
	if gx != "" {
		return myErrGo("Syntax →" + gx)
	}
	return nil
}

// Remove blank spaces between a sign and a number.
func rmBlankSpace(gx *string) {
	*gx = strings.ReplaceAll(*gx, "- ", "-")
	*gx = strings.ReplaceAll(*gx, "+ ", "")
	*gx = strings.ReplaceAll(*gx, "+", "")
}

// Split the math func by its Kx^n.
func split(gx string) []string {
	if strings.Contains(gx, " ") {
		gx = strings.Trim(gx, " ")
		return strings.Split(gx, " ")
	}
	return []string{gx}
}

// Rebuild each element to the form kx^n.
func rebuild(x string, sgx []string) []string {
	for i := range sgx {
		if strings.Contains(sgx[i], x) {
			case1 := strings.HasPrefix(sgx[i], x)
			case2 := strings.HasPrefix(sgx[i], "-"+x)

			if case1 || case2 {
				sgx[i] = strings.Replace(sgx[i], x, "1"+x, 1)
			}

			if !strings.Contains(sgx[i], x+"^") {
				sgx[i] = strings.Replace(sgx[i], x, x+"^1", 1)
			}
		} else {
			sgx[i] += x + "^0"
		}
	}
	return sgx
}

// Store k and n in different sets.
func store(x string, sKxn []string) ([]float64, [][]float64) {
	sN := make([]float64, len(sKxn))
	ssKN := make([][]float64, len(sKxn))
	for i, item := range sKxn {
		k, n, _ := strings.Cut(item, x+"^")
		kF, _ := strconv.ParseFloat(k, 64)
		nF, _ := strconv.ParseFloat(n, 64)
		sN[i], ssKN[i] = nF, []float64{kF, nF}
	}
	return sN, ssKN
}

// From highest to lowest number.
func highlow(s []float64) []float64 {
	var temp float64
	for x := range s {
		for y := range s {
			if s[x] > s[y] {
				temp = s[x]
				s[x] = s[y]
				s[y] = temp
			}
		}
	}
	return s
}

// Remove all duplicate n in sn.
func rmDuplicate(sN *[]float64) {
	s, dicc := []float64{}, make(map[float64]bool)
	for _, n := range *sN {
		if _, key := dicc[n]; !key {
			dicc[n] = true
			s = append(s, n)
		}
	}
	*sN = highlow(s)
}

// Add all the k's with the same n.
func simplify_A(ln int, sN []float64, ssKN [][]float64) [][]float64 {
	t1, simplifykn := 0, make([][]float64, len(sN))
	for _, n := range sN {

		t2, ssN := 0, make([][]float64, ln)
		for i := range ssKN {
			if n == ssKN[i][1] {
				ssN[t2] = []float64{ssKN[i][0], ssKN[i][1]}
				t2++
			}
		}

		var temp float64
		for i := range ssN[:t2] {
			temp += ssN[i][0]
		}

		simplifykn[t1] = []float64{temp, n}
		t1++
	}
	return simplifykn
}

// Add blank spaces between a sign and a number.
func addBlackSpace(kxn string) string {
	if kxn != "" {
		if strings.HasPrefix(kxn, "-") {
			kxn = " " + kxn
		} else {
			kxn = " +" + kxn
		}
	}
	return kxn
}

// Create new gz from what was gx=kxn's.
func create(x string, ssgz [][]float64, simple bool) string {
	var gz string
	for i := range ssgz {
		k := fmt.Sprint(ssgz[i][0])
		n := fmt.Sprint(ssgz[i][1])
		kxn := k + x + "^" + n

		if simple {
			if k == "0" {
				kxn = "0"
			} else if k == "1" && n == "0" {
				kxn = k
			} else if k == "1" && n == "1" {
				kxn = x
			} else if k != "1" && n == "1" {
				kxn = k + x
			} else if k != "1" && n == "0" {
				kxn = k
			} else if k == "1" && n != "0" && n != "1" {
				kxn = x + "^" + n
			}
		}

		if i == 0 {
			gz += kxn
		} else {
			gz += addBlackSpace(kxn)
		}
	}
	return gz
}

// ! Add: Addition of mathematical functions.
func A(x, gx string, simple bool) (string, error) {

	err := exists(x, gx)
	if err != nil {
		return "", err
	}

	err = correctly(x, gx, "")
	if err != nil {
		return "", err
	}

	rmBlankSpace(&gx)
	sgx := split(gx)
	sKxn := rebuild(x, sgx)
	sN, ssKN := store(x, sKxn)

	rmDuplicate(&sN)
	ssgz := simplify_A(len(sKxn), sN, ssKN)
	gz := create(x, ssgz, simple)

	return gz, nil
}
