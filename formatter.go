package rupiah

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func absFloat64(i float64) float64 {
	if i < 0 {
		return -i
	}
	return i
}

func absInt64(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

func FormatFloat64ToRp(amount float64) string {

	f := fmt.Sprintf("%f", amount)
	s := strings.Split(f, ".")

	sa := s[0]
	sb := s[1]

	sb = sb[0:2]

	r := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != sa; {
		n = sa
		sa = r.ReplaceAllString(sa, "$1.$2")
	}

	if amount < 0 {
		sa = strings.Replace(sa, "-", "", -1)

		return fmt.Sprintf("-Rp%s,%s", sa, sb)
	}

	return fmt.Sprintf("Rp%s,%s", sa, sb)

}

func FormatInt64ToRp(amount int64) string {

	s := fmt.Sprintf("%d", amount)

	r := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != s; {
		n = s
		s = r.ReplaceAllString(s, "$1.$2")
	}

	if amount < 0 {
		s = strings.Replace(s, "-", "", -1)

		return fmt.Sprintf("-Rp%s,00", s)
	}

	return fmt.Sprintf("Rp%s,00", s)

}

func FormatFloat64ToWords(amount float64) string {

	var minus = ""
	if amount < 0 {
		amount = absFloat64(amount)
		minus = "minus "
	}

	f := fmt.Sprintf("%f", amount)
	s := strings.Split(f, ".")

	sa := s[0]
	sb := s[1]

	sb = sb[0:2]

	isa, err := strconv.Atoi(sa)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	isb, err := strconv.Atoi(sb)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	ones := []string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}

	result := ""
	sen := " "
	if isa < 12 {
		result = result + ones[isa]
	} else if isa < 20 {
		result = result + FormatFloat64ToWords(float64(isa-10)) + " belas"
	} else if isa < 100 {
		result = result + FormatFloat64ToWords(float64(isa/10)) + " puluh " + FormatFloat64ToWords(float64(isa%10))
	} else if isa < 200 {
		result = result + "seratus " + FormatFloat64ToWords(float64(isa-100))
	} else if isa < 1000 {
		result = result + FormatFloat64ToWords(float64(isa/100)) + " ratus " + FormatFloat64ToWords(float64(isa%100))
	} else if isa < 2000 {
		result = result + "seribu " + FormatFloat64ToWords(float64(isa-1000))
	} else if isa < 1000000 {
		result = result + FormatFloat64ToWords(float64(isa/1000)) + " ribu " + FormatFloat64ToWords(float64(isa%1000))
	} else if isa < 1000000000 {
		result = result + FormatFloat64ToWords(float64(isa/1000000)) + " juta " + FormatFloat64ToWords(float64(isa%1000000))
	} else if isa < 1000000000000 {
		result = result + FormatFloat64ToWords(float64(isa/1000000000)) + " miliar " + FormatFloat64ToWords(float64(isa%1000000000))
	} else if isa < 1000000000000000 {
		result = result + FormatFloat64ToWords(float64(isa/1000000000000)) + " triliun " + FormatFloat64ToWords(float64(isa%1000000000000))
	} else if isa < 1000000000000000000 {
		result = result + FormatFloat64ToWords(float64(isa/1000000000000000)) + " kuadriliun " + FormatFloat64ToWords(float64(isa%1000000000000000))
	}

	if isb > 0 {
		if isb < 12 {
			sen = sen + ones[isb]
		} else if isb < 20 {
			sen = sen + FormatFloat64ToWords(float64(isb-10)) + " belas"
		} else if isb < 100 {
			sen = sen + FormatFloat64ToWords(float64(isb/10)) + " puluh " + FormatFloat64ToWords(float64(isb%10))
		}

		result = strings.Replace(result, "rupiah", "", -1)
		result = minus + result + " rupiah" + sen + " sen"
		result = strings.Replace(result, "  ", " ", -1)
		return result
	}

	result = minus + result

	result = strings.Replace(result, "rupiah", "", -1)
	result = strings.Replace(result, "  ", " ", -1)

	if isa >= 12 {
		result = result + "rupiah"
	}

	return result

}

func FormatInt64ToWords(amount int64) string {

	var minus = ""
	if amount < 0 {
		amount = absInt64(amount)
		minus = "minus "
	}

	ones := []string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}

	result := ""
	if amount < 12 {
		result = result + ones[amount]
	} else if amount < 20 {
		result = result + FormatInt64ToWords(amount-10) + " belas"
	} else if amount < 100 {
		result = result + FormatInt64ToWords(amount/10) + " puluh " + FormatInt64ToWords(amount%10)
	} else if amount < 200 {
		result = result + "seratus " + FormatInt64ToWords(amount-100)
	} else if amount < 1000 {
		result = result + FormatInt64ToWords(amount/100) + " ratus " + FormatInt64ToWords(amount%100)
	} else if amount < 2000 {
		result = result + "seribu " + FormatInt64ToWords(amount-1000)
	} else if amount < 1000000 {
		result = result + FormatInt64ToWords(amount/1000) + " ribu " + FormatInt64ToWords(amount%1000)
	} else if amount < 1000000000 {
		result = result + FormatInt64ToWords(amount/1000000) + " juta " + FormatInt64ToWords(amount%1000000)
	} else if amount < 1000000000000 {
		result = result + FormatInt64ToWords(amount/1000000000) + " miliar " + FormatInt64ToWords(amount%1000000000)
	} else if amount < 1000000000000000 {
		result = result + FormatInt64ToWords(amount/1000000000000) + " triliun " + FormatInt64ToWords(amount%1000000000000)
	} else if amount < 1000000000000000000 {
		result = result + FormatInt64ToWords(amount/1000000000000000) + " kuadriliun " + FormatInt64ToWords(amount%1000000000000000)
	}

	result = minus + result

	result = strings.Replace(result, "rupiah", "", -1)
	result = strings.Replace(result, "  ", " ", -1)
	result = result + " rupiah"
	result = strings.Replace(result, "  ", " ", -1)
	return result
}
