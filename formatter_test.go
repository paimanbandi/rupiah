package rupiah

import "testing"

//-- FormatFloat64ToRp (positive)
func TestFormatFloat64ToRp_Ones_POS(t *testing.T) {
	want := "Rp1,20"
	if got := FormatFloat64ToRp(1.20456); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Teens_POS(t *testing.T) {
	want := "Rp17,45"
	if got := FormatFloat64ToRp(17.454); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Tens_POS(t *testing.T) {
	want := "Rp32,96"
	if got := FormatFloat64ToRp(32.96); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Hundreds_POS(t *testing.T) {
	want := "Rp289,11"
	if got := FormatFloat64ToRp(289.1176); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Thousands_POS(t *testing.T) {
	want := "Rp4.200,00"
	if got := FormatFloat64ToRp(4200.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_TensThousands_POS(t *testing.T) {
	want := "Rp98.360,20"
	if got := FormatFloat64ToRp(98360.20); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_HundredsThousands_POS(t *testing.T) {
	want := "Rp615.510,30"
	if got := FormatFloat64ToRp(615510.30); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Millions_POS(t *testing.T) {
	want := "Rp5.786.000,00"
	if got := FormatFloat64ToRp(5786000.00235); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_TensMillions_POS(t *testing.T) {
	want := "Rp82.712.090,01"
	if got := FormatFloat64ToRp(82712090.0156); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_HundredsMillions_POS(t *testing.T) {
	want := "Rp377.475.000,50"
	if got := FormatFloat64ToRp(377475000.50345); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Billions_POS(t *testing.T) {
	want := "Rp6.934.550.700,00"
	if got := FormatFloat64ToRp(6934550700.008); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_TensBillions_POS(t *testing.T) {
	want := "Rp88.345.350.700,10"
	if got := FormatFloat64ToRp(88345350700.10); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_HundressBillions_POS(t *testing.T) {
	want := "Rp782.341.350.900,80"
	if got := FormatFloat64ToRp(782341350900.80); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Trillions_POS(t *testing.T) {
	want := "Rp3.232.190.050.000,00"
	if got := FormatFloat64ToRp(3232190050000.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_TensTrillions_POS(t *testing.T) {
	want := "Rp55.226.900.005.000,00"
	if got := FormatFloat64ToRp(55226900005000.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_HundredsTrillions_POS(t *testing.T) {
	want := "Rp341.117.004.000.000,00"
	if got := FormatFloat64ToRp(341117004000000.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Quadrillions_POS(t *testing.T) {
	want := "Rp9.011.150.000.000.000,00"
	if got := FormatFloat64ToRp(9011150000000000.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

//-- FormatFloat64ToRp (negative)
func TestFormatFloat64ToRp_Ones_NEG(t *testing.T) {
	want := "-Rp1,20"
	if got := FormatFloat64ToRp(-1.20); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Teens_NEG(t *testing.T) {
	want := "-Rp17,45"
	if got := FormatFloat64ToRp(-17.45); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Tens_NEG(t *testing.T) {
	want := "-Rp32,96"
	if got := FormatFloat64ToRp(-32.96); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Hundreds_NEG(t *testing.T) {
	want := "-Rp289,11"
	if got := FormatFloat64ToRp(-289.11); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Thousands_NEG(t *testing.T) {
	want := "-Rp4.200,00"
	if got := FormatFloat64ToRp(-4200.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_TensThousands_NEG(t *testing.T) {
	want := "-Rp98.360,20"
	if got := FormatFloat64ToRp(-98360.20); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_HundredsThousands_NEG(t *testing.T) {
	want := "-Rp615.510,30"
	if got := FormatFloat64ToRp(-615510.30); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Millions_NEG(t *testing.T) {
	want := "-Rp5.786.000,00"
	if got := FormatFloat64ToRp(-5786000.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_TensMillions_NEG(t *testing.T) {
	want := "-Rp82.712.090,01"
	if got := FormatFloat64ToRp(-82712090.01); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_HundredsMillions_NEG(t *testing.T) {
	want := "-Rp377.475.000,50"
	if got := FormatFloat64ToRp(-377475000.50); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Billions_NEG(t *testing.T) {
	want := "-Rp6.934.550.700,00"
	if got := FormatFloat64ToRp(-6934550700.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_TensBillions_NEG(t *testing.T) {
	want := "-Rp88.345.350.700,10"
	if got := FormatFloat64ToRp(-88345350700.10); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_HundressBillions_NEG(t *testing.T) {
	want := "-Rp782.341.350.900,80"
	if got := FormatFloat64ToRp(-782341350900.80); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Trillions_NEG(t *testing.T) {
	want := "-Rp3.232.190.050.000,00"
	if got := FormatFloat64ToRp(-3232190050000.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_TensTrillions_NEG(t *testing.T) {
	want := "-Rp55.226.900.005.000,00"
	if got := FormatFloat64ToRp(-55226900005000.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_HundredsTrillions_NEG(t *testing.T) {
	want := "-Rp341.117.004.000.000,00"
	if got := FormatFloat64ToRp(-341117004000000.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToRp_Quadrillions_NEG(t *testing.T) {
	want := "-Rp9.011.150.000.000.000,00"
	if got := FormatFloat64ToRp(-9011150000000000.00); got != want {
		t.Errorf("FormatFloat64ToRp() = %q, want %q", got, want)
	}
}

//-- FormatInt64ToRp (positive)
func TestFormatInt64ToRp_Ones_POS(t *testing.T) {
	want := "Rp1,00"
	if got := FormatInt64ToRp(1); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Teens_POS(t *testing.T) {
	want := "Rp17,00"
	if got := FormatInt64ToRp(17); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Tens_POS(t *testing.T) {
	want := "Rp33,00"
	if got := FormatInt64ToRp(33); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Hundreds_POS(t *testing.T) {
	want := "Rp289,00"
	if got := FormatInt64ToRp(289); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Thousands_POS(t *testing.T) {
	want := "Rp4.200,00"
	if got := FormatInt64ToRp(4200); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_TensThousands_POS(t *testing.T) {
	want := "Rp98.360,00"
	if got := FormatInt64ToRp(98360); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_HundredsThousands_POS(t *testing.T) {
	want := "Rp615.510,00"
	if got := FormatInt64ToRp(615510); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Millions_POS(t *testing.T) {
	want := "Rp5.786.000,00"
	if got := FormatInt64ToRp(5786000.00); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_TensMillions_POS(t *testing.T) {
	want := "Rp82.712.090,00"
	if got := FormatInt64ToRp(82712090.00); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_HundredsMillions_POS(t *testing.T) {
	want := "Rp377.475.000,00"
	if got := FormatInt64ToRp(377475000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Billions_POS(t *testing.T) {
	want := "Rp6.934.550.700,00"
	if got := FormatInt64ToRp(6934550700); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_TensBillions_POS(t *testing.T) {
	want := "Rp88.345.350.700,00"
	if got := FormatInt64ToRp(88345350700); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_HundressBillions_POS(t *testing.T) {
	want := "Rp782.341.350.900,00"
	if got := FormatInt64ToRp(782341350900); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Trillions_POS(t *testing.T) {
	want := "Rp3.232.190.050.000,00"
	if got := FormatInt64ToRp(3232190050000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_TensTrillions_POS(t *testing.T) {
	want := "Rp55.226.900.005.000,00"
	if got := FormatInt64ToRp(55226900005000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_HundredsTrillions_POS(t *testing.T) {
	want := "Rp341.117.004.000.000,00"
	if got := FormatInt64ToRp(341117004000000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Quadrillions_POS(t *testing.T) {
	want := "Rp9.011.150.000.000.000,00"
	if got := FormatInt64ToRp(9011150000000000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

//-- FormatInt64ToRp (negative)
func TestFormatInt64ToRp_Ones_NEG(t *testing.T) {
	want := "-Rp1,00"
	if got := FormatInt64ToRp(-1); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Teens_NEG(t *testing.T) {
	want := "-Rp17,00"
	if got := FormatInt64ToRp(-17); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Tens_NEG(t *testing.T) {
	want := "-Rp33,00"
	if got := FormatInt64ToRp(-33); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Hundreds_NEG(t *testing.T) {
	want := "-Rp289,00"
	if got := FormatInt64ToRp(-289); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Thousands_NEG(t *testing.T) {
	want := "-Rp4.200,00"
	if got := FormatInt64ToRp(-4200); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_TensThousands_NEG(t *testing.T) {
	want := "-Rp98.360,00"
	if got := FormatInt64ToRp(-98360); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_HundredsThousands_NEG(t *testing.T) {
	want := "-Rp615.510,00"
	if got := FormatInt64ToRp(-615510); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Millions_NEG(t *testing.T) {
	want := "-Rp5.786.000,00"
	if got := FormatInt64ToRp(-5786000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_TensMillions_NEG(t *testing.T) {
	want := "-Rp82.712.090,00"
	if got := FormatInt64ToRp(-82712090); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_HundredsMillions_NEG(t *testing.T) {
	want := "-Rp377.475.000,00"
	if got := FormatInt64ToRp(-377475000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Billions_NEG(t *testing.T) {
	want := "-Rp6.934.550.700,00"
	if got := FormatInt64ToRp(-6934550700); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_TensBillions_NEG(t *testing.T) {
	want := "-Rp88.345.350.700,00"
	if got := FormatInt64ToRp(-88345350700); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_HundressBillions_NEG(t *testing.T) {
	want := "-Rp782.341.350.900,00"
	if got := FormatInt64ToRp(-782341350900); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Trillions_NEG(t *testing.T) {
	want := "-Rp3.232.190.050.000,00"
	if got := FormatInt64ToRp(-3232190050000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_TensTrillions_NEG(t *testing.T) {
	want := "-Rp55.226.900.005.000,00"
	if got := FormatInt64ToRp(-55226900005000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_HundredsTrillions_NEG(t *testing.T) {
	want := "-Rp341.117.004.000.000,00"
	if got := FormatInt64ToRp(-341117004000000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToRp_Quadrillions_NEG(t *testing.T) {
	want := "-Rp9.011.150.000.000.000,00"
	if got := FormatInt64ToRp(-9011150000000000); got != want {
		t.Errorf("FormatInt64ToRp() = %q, want %q", got, want)
	}
}

//-- FormatFloat64ToWords (positive)
func TestFormatFloat64ToWords_Ones_POS(t *testing.T) {
	want := "satu rupiah dua puluh sen"
	if got := FormatFloat64ToWords(1.20); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Teens_POS(t *testing.T) {
	want := "tujuh belas rupiah empat puluh lima sen"
	if got := FormatFloat64ToWords(17.45); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Tens_POS(t *testing.T) {
	want := "tiga puluh dua rupiah sembilan puluh enam sen"
	if got := FormatFloat64ToWords(32.96); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Hundreds_POS(t *testing.T) {
	want := "dua ratus delapan puluh sembilan rupiah sebelas sen"
	if got := FormatFloat64ToWords(289.11); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Thousands_POS(t *testing.T) {
	want := "empat ribu dua ratus rupiah"
	if got := FormatFloat64ToWords(4200.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_TensThousands_POS(t *testing.T) {
	want := "sembilan puluh delapan ribu tiga ratus enam puluh rupiah dua puluh sen"
	if got := FormatFloat64ToWords(98360.20); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_HundredsThousands_POS(t *testing.T) {
	want := "enam ratus lima belas ribu lima ratus sepuluh rupiah tiga puluh sen"
	if got := FormatFloat64ToWords(615510.30); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Millions_POS(t *testing.T) {
	want := "lima juta tujuh ratus delapan puluh enam ribu rupiah"
	if got := FormatFloat64ToWords(5786000.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_TensMillions_POS(t *testing.T) {
	want := "delapan puluh dua juta tujuh ratus dua belas ribu sembilan puluh rupiah satu sen"
	if got := FormatFloat64ToWords(82712090.01); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_HundredsMillions_POS(t *testing.T) {
	want := "tiga ratus tujuh puluh tujuh juta empat ratus tujuh puluh lima ribu rupiah lima puluh sen"
	if got := FormatFloat64ToWords(377475000.50); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Billions_POS(t *testing.T) {
	want := "enam miliar sembilan ratus tiga puluh empat juta lima ratus lima puluh ribu tujuh ratus rupiah"
	if got := FormatFloat64ToWords(6934550700.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_TensBillions_POS(t *testing.T) {
	want := "delapan puluh delapan miliar tiga ratus empat puluh lima juta tiga ratus lima puluh ribu tujuh ratus rupiah sepuluh sen"
	if got := FormatFloat64ToWords(88345350700.10); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_HundressBillions_POS(t *testing.T) {
	want := "tujuh ratus delapan puluh dua miliar tiga ratus empat puluh satu juta tiga ratus lima puluh ribu sembilan ratus rupiah delapan puluh sen"
	if got := FormatFloat64ToWords(782341350900.80); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Trillions_POS(t *testing.T) {
	want := "tiga triliun dua ratus tiga puluh dua miliar seratus sembilan puluh juta lima puluh ribu rupiah"
	if got := FormatFloat64ToWords(3232190050000.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_TensTrillions_POS(t *testing.T) {
	want := "lima puluh lima triliun dua ratus dua puluh enam miliar sembilan ratus juta lima ribu rupiah"
	if got := FormatFloat64ToWords(55226900005000.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_HundredsTrillions_POS(t *testing.T) {
	want := "tiga ratus empat puluh satu triliun seratus tujuh belas miliar empat juta rupiah"
	if got := FormatFloat64ToWords(341117004000000.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Quadrillions_POS(t *testing.T) {
	want := "sembilan kuadriliun sebelas triliun seratus lima puluh miliar rupiah"
	if got := FormatFloat64ToWords(9011150000000000.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

//-- FormatFloat64ToWords (negative)
func TestFormatFloat64ToWords_Ones_NEG(t *testing.T) {
	want := "minus satu rupiah dua puluh sen"
	if got := FormatFloat64ToWords(-1.20); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Teens_NEG(t *testing.T) {
	want := "minus tujuh belas rupiah empat puluh lima sen"
	if got := FormatFloat64ToWords(-17.45); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Tens_NEG(t *testing.T) {
	want := "minus tiga puluh dua rupiah sembilan puluh enam sen"
	if got := FormatFloat64ToWords(-32.96); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Hundreds_NEG(t *testing.T) {
	want := "minus dua ratus delapan puluh sembilan rupiah sebelas sen"
	if got := FormatFloat64ToWords(-289.11); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Thousands_NEG(t *testing.T) {
	want := "minus empat ribu dua ratus rupiah"
	if got := FormatFloat64ToWords(-4200.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_TensThousands_NEG(t *testing.T) {
	want := "minus sembilan puluh delapan ribu tiga ratus enam puluh rupiah dua puluh sen"
	if got := FormatFloat64ToWords(-98360.20); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_HundredsThousands_NEG(t *testing.T) {
	want := "minus enam ratus lima belas ribu lima ratus sepuluh rupiah tiga puluh sen"
	if got := FormatFloat64ToWords(-615510.30); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Millions_NEG(t *testing.T) {
	want := "minus lima juta tujuh ratus delapan puluh enam ribu rupiah"
	if got := FormatFloat64ToWords(-5786000.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_TensMillions_NEG(t *testing.T) {
	want := "minus delapan puluh dua juta tujuh ratus dua belas ribu sembilan puluh rupiah satu sen"
	if got := FormatFloat64ToWords(-82712090.01); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_HundredsMillions_NEG(t *testing.T) {
	want := "minus tiga ratus tujuh puluh tujuh juta empat ratus tujuh puluh lima ribu rupiah lima puluh sen"
	if got := FormatFloat64ToWords(-377475000.50); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Billions_NEG(t *testing.T) {
	want := "minus enam miliar sembilan ratus tiga puluh empat juta lima ratus lima puluh ribu tujuh ratus rupiah"
	if got := FormatFloat64ToWords(-6934550700.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_TensBillions_NEG(t *testing.T) {
	want := "minus delapan puluh delapan miliar tiga ratus empat puluh lima juta tiga ratus lima puluh ribu tujuh ratus rupiah sepuluh sen"
	if got := FormatFloat64ToWords(-88345350700.10); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_HundressBillions_NEG(t *testing.T) {
	want := "minus tujuh ratus delapan puluh dua miliar tiga ratus empat puluh satu juta tiga ratus lima puluh ribu sembilan ratus rupiah delapan puluh sen"
	if got := FormatFloat64ToWords(-782341350900.80); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Trillions_NEG(t *testing.T) {
	want := "minus tiga triliun dua ratus tiga puluh dua miliar seratus sembilan puluh juta lima puluh ribu rupiah"
	if got := FormatFloat64ToWords(-3232190050000.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_TensTrillions_NEG(t *testing.T) {
	want := "minus lima puluh lima triliun dua ratus dua puluh enam miliar sembilan ratus juta lima ribu rupiah"
	if got := FormatFloat64ToWords(-55226900005000.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_HundredsTrillions_NEG(t *testing.T) {
	want := "minus tiga ratus empat puluh satu triliun seratus tujuh belas miliar empat juta rupiah"
	if got := FormatFloat64ToWords(-341117004000000.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatFloat64ToWords_Quadrillions_NEG(t *testing.T) {
	want := "minus sembilan kuadriliun sebelas triliun seratus lima puluh miliar rupiah"
	if got := FormatFloat64ToWords(-9011150000000000.00); got != want {
		t.Errorf("FormatFloat64ToWords() = %q, want %q", got, want)
	}
}

//-- FormatInt64ToWords (positive)
func TestFormatInt64ToWords_Ones_POS(t *testing.T) {
	want := "satu rupiah"
	if got := FormatInt64ToWords(1); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Teens_POS(t *testing.T) {
	want := "tujuh belas rupiah"
	if got := FormatInt64ToWords(17); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Tens_POS(t *testing.T) {
	want := "tiga puluh tiga rupiah"
	if got := FormatInt64ToWords(33); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Hundreds_POS(t *testing.T) {
	want := "dua ratus delapan puluh sembilan rupiah"
	if got := FormatInt64ToWords(289); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Thousands_POS(t *testing.T) {
	want := "empat ribu dua ratus rupiah"
	if got := FormatInt64ToWords(4200); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_TensThousands_POS(t *testing.T) {
	want := "sembilan puluh delapan ribu tiga ratus enam puluh rupiah"
	if got := FormatInt64ToWords(98360); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_HundredsThousands_POS(t *testing.T) {
	want := "enam ratus lima belas ribu lima ratus sepuluh rupiah"
	if got := FormatInt64ToWords(615510); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Millions_POS(t *testing.T) {
	want := "lima juta tujuh ratus delapan puluh enam ribu rupiah"
	if got := FormatInt64ToWords(5786000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_TensMillions_POS(t *testing.T) {
	want := "delapan puluh dua juta tujuh ratus dua belas ribu sembilan puluh rupiah"
	if got := FormatInt64ToWords(82712090); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_HundredsMillions_POS(t *testing.T) {
	want := "tiga ratus tujuh puluh tujuh juta empat ratus tujuh puluh lima ribu rupiah"
	if got := FormatInt64ToWords(377475000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Billions_POS(t *testing.T) {
	want := "enam miliar sembilan ratus tiga puluh empat juta lima ratus lima puluh ribu tujuh ratus rupiah"
	if got := FormatInt64ToWords(6934550700); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_TensBillions_POS(t *testing.T) {
	want := "delapan puluh delapan miliar tiga ratus empat puluh lima juta tiga ratus lima puluh ribu tujuh ratus rupiah"
	if got := FormatInt64ToWords(88345350700); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_HundressBillions_POS(t *testing.T) {
	want := "tujuh ratus delapan puluh dua miliar tiga ratus empat puluh satu juta tiga ratus lima puluh ribu sembilan ratus rupiah"
	if got := FormatInt64ToWords(782341350900); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Trillions_POS(t *testing.T) {
	want := "tiga triliun dua ratus tiga puluh dua miliar seratus sembilan puluh juta lima puluh ribu rupiah"
	if got := FormatInt64ToWords(3232190050000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_TensTrillions_POS(t *testing.T) {
	want := "lima puluh lima triliun dua ratus dua puluh enam miliar sembilan ratus juta lima ribu rupiah"
	if got := FormatInt64ToWords(55226900005000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_HundredsTrillions_POS(t *testing.T) {
	want := "tiga ratus empat puluh satu triliun seratus tujuh belas miliar empat juta rupiah"
	if got := FormatInt64ToWords(341117004000000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Quadrillions_POS(t *testing.T) {
	want := "sembilan kuadriliun sebelas triliun seratus lima puluh miliar rupiah"
	if got := FormatInt64ToWords(9011150000000000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

//-- FormatInt64ToWords (negative)
func TestFormatInt64ToWords_Ones_NEG(t *testing.T) {
	want := "minus satu rupiah"
	if got := FormatInt64ToWords(-1); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Teens_NEG(t *testing.T) {
	want := "minus tujuh belas rupiah"
	if got := FormatInt64ToWords(-17); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Tens_NEG(t *testing.T) {
	want := "minus tiga puluh tiga rupiah"
	if got := FormatInt64ToWords(-33); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Hundreds_NEG(t *testing.T) {
	want := "minus dua ratus delapan puluh sembilan rupiah"
	if got := FormatInt64ToWords(-289); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Thousands_NEG(t *testing.T) {
	want := "minus empat ribu dua ratus rupiah"
	if got := FormatInt64ToWords(-4200); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_TensThousands_NEG(t *testing.T) {
	want := "minus sembilan puluh delapan ribu tiga ratus enam puluh rupiah"
	if got := FormatInt64ToWords(-98360); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_HundredsThousands_NEG(t *testing.T) {
	want := "minus enam ratus lima belas ribu lima ratus sepuluh rupiah"
	if got := FormatInt64ToWords(-615510); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Millions_NEG(t *testing.T) {
	want := "minus lima juta tujuh ratus delapan puluh enam ribu rupiah"
	if got := FormatInt64ToWords(-5786000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_TensMillions_NEG(t *testing.T) {
	want := "minus delapan puluh dua juta tujuh ratus dua belas ribu sembilan puluh rupiah"
	if got := FormatInt64ToWords(-82712090); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_HundredsMillions_NEG(t *testing.T) {
	want := "minus tiga ratus tujuh puluh tujuh juta empat ratus tujuh puluh lima ribu rupiah"
	if got := FormatInt64ToWords(-377475000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Billions_NEG(t *testing.T) {
	want := "minus enam miliar sembilan ratus tiga puluh empat juta lima ratus lima puluh ribu tujuh ratus rupiah"
	if got := FormatInt64ToWords(-6934550700); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_TensBillions_NEG(t *testing.T) {
	want := "minus delapan puluh delapan miliar tiga ratus empat puluh lima juta tiga ratus lima puluh ribu tujuh ratus rupiah"
	if got := FormatInt64ToWords(-88345350700); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_HundressBillions_NEG(t *testing.T) {
	want := "minus tujuh ratus delapan puluh dua miliar tiga ratus empat puluh satu juta tiga ratus lima puluh ribu sembilan ratus rupiah"
	if got := FormatInt64ToWords(-782341350900); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Trillions_NEG(t *testing.T) {
	want := "minus tiga triliun dua ratus tiga puluh dua miliar seratus sembilan puluh juta lima puluh ribu rupiah"
	if got := FormatInt64ToWords(-3232190050000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_TensTrillions_NEG(t *testing.T) {
	want := "minus lima puluh lima triliun dua ratus dua puluh enam miliar sembilan ratus juta lima ribu rupiah"
	if got := FormatInt64ToWords(-55226900005000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_HundredsTrillions_NEG(t *testing.T) {
	want := "minus tiga ratus empat puluh satu triliun seratus tujuh belas miliar empat juta rupiah"
	if got := FormatInt64ToWords(-341117004000000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}

func TestFormatInt64ToWords_Quadrillions_NEG(t *testing.T) {
	want := "minus sembilan kuadriliun sebelas triliun seratus lima puluh miliar rupiah"
	if got := FormatInt64ToWords(-9011150000000000); got != want {
		t.Errorf("FormatInt64ToWords() = %q, want %q", got, want)
	}
}
