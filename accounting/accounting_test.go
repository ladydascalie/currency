package accounting_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	fuzz "github.com/google/gofuzz"
	"github.com/ladydascalie/currency"
	"github.com/ladydascalie/currency/accounting"
)

func ExampleFloat64ToAmount_gbp() {
	gbp := accounting.Float64ToAmount(currency.GBP, 32.32)
	fmt.Printf("minor value: %d, stringer: %s", gbp.MinorValue, gbp)
	// output: minor value: 3232, stringer: 32.32 GBP
}

func ExampleFloat64ToAmount_jpy() {
	// JPY doesn't allow for values after the decimal dot, since it has a
	// minor currency unit of 1.
	jpy := accounting.Float64ToAmount(currency.JPY, 32.32)
	fmt.Printf("minor value: %d, stringer: %s", jpy.MinorValue, jpy)
	// output: minor value: 32, stringer: 32 JPY
}

func ExampleExchange_usd_eur() {
	// 4 May 2020 08:00 UTC - 5 May 2020 08:01 UTC
	// EUR/USD close:1.08968 low:1.08871 high:1.09479
	usd := accounting.Float64ToAmount(currency.USD, 100.0)
	eur, err := accounting.Exchange(usd, currency.EUR, 1.08968)
	if err != nil {
		// handle...
	}
	fmt.Println(eur)
	// output: 91.77 EUR
}

func ExampleExchange_usd_jpy() {
	// 4 May 2020 07:40 UTC - 5 May 2020 07:40 UTC
	// JPY/USD close:0.00937 low:0.00934 high:0.00939
	usd := accounting.Float64ToAmount(currency.USD, 100.0)
	jpy, err := accounting.Exchange(usd, currency.JPY, 0.00937)
	if err != nil {
		// handle...
	}
	fmt.Println(jpy)
	// output: 10672 JPY
}

func TestAmount_String(t *testing.T) {
	t.Run("GBP", func(t *testing.T) {
		got := accounting.MakeAmount(currency.GBP, 1234).String()
		want := "12.34 GBP"
		if diff := cmp.Diff(got, want); diff != "" {
			t.Fatalf("failed (-got +want): %s", diff)
		}
	})
	t.Run("JPY", func(t *testing.T) {
		got := accounting.MakeAmount(currency.JPY, 1234).String()
		want := "1234 JPY"
		if diff := cmp.Diff(got, want); diff != "" {
			t.Fatalf("failed (-got +want): %s", diff)
		}
	})
}

func TestExchange(t *testing.T) {
	// factors of 1, 100, 1000 and 10000 are represented.
	//
	// As there is no point in duplicating tests for currencies with
	// the exact same minor unit, the better approach is to instead
	// bulk up the amount of unique values tested.
	//
	// The test is essentially:
	// value / factor / rate = want
	//
	// Each value in the want array has been calculated manually.
	tests := []struct {
		expectedErr error
		from        currency.Currency
		to          currency.Currency
		name        string
		amount      float64
		expected    float64
		rate        float64
	}{
		{
			name:     "100 USD in USD",
			expected: 100,
			rate:     1,
			from:     currency.USD,
			to:       currency.USD,
		},
		{
			name:     "100 USD in EUR",
			expected: 91.61,
			rate:     1.0916,
			from:     currency.USD,
			to:       currency.EUR,
		},
		{
			name:     "100 USD in GBP",
			expected: 80.51,
			rate:     1.2421,
			from:     currency.USD,
			to:       currency.GBP,
		},
		{
			name:     "100 USD in JPY",
			expected: 10678,
			rate:     0.00936545,
			from:     currency.USD,
			to:       currency.JPY,
		},
		{
			name:     "100 USD in TND",
			expected: 290.02,
			rate:     0.3448,
			from:     currency.USD,
			to:       currency.TND,
		},
		{
			name:     "100 EUR in USD",
			expected: 109.16,
			rate:     0.9161,
			from:     currency.EUR,
			to:       currency.USD,
		},
		{
			name:     "100 EUR in EUR",
			expected: 100,
			rate:     1,
			from:     currency.EUR,
			to:       currency.EUR,
		},
		{
			name:     "100 EUR in GBP",
			expected: 87.88,
			rate:     1.1379,
			from:     currency.EUR,
			to:       currency.GBP,
		},
		{
			name:     "100 EUR in JPY",
			expected: 11628,
			rate:     0.0086,
			from:     currency.EUR,
			to:       currency.JPY,
		},
		{
			name:     "100 EUR in TND",
			expected: 316.66,
			rate:     0.3158,
			from:     currency.EUR,
			to:       currency.TND,
		},
		{
			name:     "100 GBP in USD",
			expected: 124.21,
			rate:     0.8051,
			from:     currency.GBP,
			to:       currency.USD,
		},
		{
			name:     "100 GBP in EUR",
			expected: 113.79,
			rate:     0.8788,
			from:     currency.GBP,
			to:       currency.EUR,
		},
		{
			name:     "100 GBP in GBP",
			expected: 100,
			rate:     1,
			from:     currency.GBP,
			to:       currency.GBP,
		},
		{
			name:     "100 GBP in JPY",
			expected: 13333,
			rate:     0.0075,
			from:     currency.GBP,
			to:       currency.JPY,
		},
		{
			name:     "100 GBP in TND",
			expected: 360.23,
			rate:     0.2776,
			from:     currency.GBP,
			to:       currency.TND,
		},
		{
			name:     "100 JPY in USD",
			expected: 0.94,
			rate:     106.9404,
			from:     currency.JPY,
			to:       currency.USD,
		},
		{
			name:     "100 JPY in EUR",
			expected: 0.86,
			rate:     116.7394,
			from:     currency.JPY,
			to:       currency.EUR,
		},
		{
			name:     "100 JPY in GBP",
			expected: 0.75,
			rate:     132.8351,
			from:     currency.JPY,
			to:       currency.GBP,
		},
		{
			name:     "100 JPY in JPY",
			expected: 100,
			rate:     1,
			from:     currency.JPY,
			to:       currency.JPY,
		},
		{
			name:     "100 JPY in TND",
			expected: 2.71,
			rate:     36.8696,
			from:     currency.JPY,
			to:       currency.TND,
		},
		{
			name:     "100 TND in USD",
			expected: 34.48,
			rate:     2.9005,
			from:     currency.TND,
			to:       currency.USD,
		},
		{
			name:     "100 TND in EUR",
			expected: 31.58,
			rate:     3.1663,
			from:     currency.TND,
			to:       currency.EUR,
		},
		{
			name:     "100 TND in GBP",
			expected: 27.76,
			rate:     3.6028,
			from:     currency.TND,
			to:       currency.GBP,
		},
		{
			name:     "100 TND in JPY",
			expected: 3690,
			rate:     0.0271,
			from:     currency.TND,
			to:       currency.JPY,
		},
		{
			name:     "100 TND in TND",
			expected: 100,
			rate:     1,
			from:     currency.TND,
			to:       currency.TND,
		},
		{
			name:        "Subzero exchange rate must fail",
			expected:    100,
			rate:        -1,
			from:        currency.TND,
			to:          currency.TND,
			expectedErr: accounting.ErrSubZeroRate,
		},
		{
			name:     "Zero rate should return zero amount",
			expected: 0,
			rate:     0,
			from:     currency.TND,
			to:       currency.TND,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amount := accounting.Float64ToAmount(tt.from, 100)
			result, err := accounting.Exchange(amount, tt.to, tt.rate)
			if tt.expectedErr != err {
				t.Fatal("expected an error but got none")
			}
			if err == nil {
				if diff := cmp.Diff(tt.expected, accounting.AmountToFloat64(result)); diff != "" {
					t.Errorf("Exchange() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestRatNetAmount(t *testing.T) {
	br := func(f64 float64) *big.Rat {
		return new(big.Rat).SetFloat64(f64)
	}

	type args struct {
		value *big.Rat
		rate  *big.Rat
	}
	tests := []struct {
		args        args
		expectedErr error
		name        string
		want        string
	}{
		{
			name: "gross 10 vat 19",
			args: args{
				value: br(10),
				rate:  br(.19),
			},
			want: "8.403361345",
		},
		{
			name: "gross 19.99 vat 20",
			args: args{
				value: br(19.99),
				rate:  br(.20),
			},
			want: "16.658333333",
		},
		{
			name: "gross 123 vat 7",
			args: args{
				value: br(123),
				rate:  br(.07),
			},
			want: "114.953271028",
		},
		{
			name: "Subzero rate should fail",
			args: args{
				value: br(123),
				rate:  br(-1),
			},
			want:        "",
			expectedErr: accounting.ErrSubZeroRate,
		},
		{
			name: "Zero rate should return zero value",
			args: args{
				value: br(123),
				rate:  br(0),
			},
			want:        "123.000000000",
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accounting.RatNetAmount(tt.args.value, tt.args.rate)
			if err != tt.expectedErr {
				t.Fatal("expected an error but got none")
			}
			if err == nil {
				if diff := cmp.Diff(tt.want, got.Text('f', 9)); diff != "" {
					t.Errorf("RatNetAmount() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestToMinorUnit(t *testing.T) {
	for _, tt := range minorUnitTestsCases {
		t.Run(tt.name, func(t *testing.T) {
			amount := accounting.Float64ToAmount(currency.GBP, tt.f64)

			if diff := cmp.Diff(tt.i64, amount.MinorValue); diff != "" {
				t.Errorf("ToMinorUnit() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFromMinorUnit(t *testing.T) {
	for _, tt := range minorUnitTestsCases {
		t.Run(tt.name, func(t *testing.T) {
			amount := accounting.MakeAmount(currency.GBP, tt.i64)

			f := accounting.AmountToFloat64(amount)
			if diff := cmp.Diff(tt.f64, f); diff != "" {
				t.Errorf("FromMinorUnit() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestValidateFloatIsPrecise(t *testing.T) {
	tests := []struct {
		name      string
		amount    float64
		wantError bool
	}{
		{
			name:      "exact number",
			amount:    10,
			wantError: false,
		},
		{
			name:      "too long",
			amount:    12.123,
			wantError: true,
		},
		{
			name:      "valid amount",
			amount:    12.12,
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := accounting.ValidateFloatIsPrecise(tt.amount)
			if (err != nil) != tt.wantError {
				t.Fatalf(
					"ValidateFloatIsPrecise: failed on %v",
					tt.amount,
				)
			}
		})
	}
}

func TestValidateManyFloatsArePrecise(t *testing.T) {
	tests := []struct {
		name      string
		amounts   []float64
		wantError bool
	}{
		{
			name:      "exact numbers",
			amounts:   []float64{10, 11, 12},
			wantError: false,
		},
		{
			name:      "one is too long",
			amounts:   []float64{11, 12.123, 13},
			wantError: true,
		},
		{
			name:      "valid amounts",
			amounts:   []float64{11.11, 12.12, 13.13},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := accounting.ValidateManyFloatsArePrecise(tt.amounts...)
			if (err != nil) != tt.wantError {
				t.Fatalf("ValidateFloatIsPrecise: failed with %v", err)
			}
		})
	}
}

func TestNetAmount(t *testing.T) {
	type args struct {
		gross int64
		rate  float64
	}
	tests := []struct {
		args    args
		want    int64
		wantErr bool
	}{
		{
			args: args{
				gross: 1999,
				rate:  0.20,
			},
			want:    1666,
			wantErr: false,
		},
		{
			args: args{
				gross: 1234,
				rate:  0.19,
			},
			want:    1037,
			wantErr: false,
		},
		{
			args: args{
				gross: 9999,
				rate:  0.33,
			},
			want:    7518,
			wantErr: false,
		},
		{
			args: args{
				gross: 1234,
				rate:  0.0,
			},
			want:    1234,
			wantErr: false,
		},
		{
			args: args{
				gross: 1234,
				rate:  -1.0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("NetAmount_%d", i), func(t *testing.T) {
			got, err := accounting.NetAmount(tt.args.gross, tt.args.rate)
			if (err != nil) != tt.wantErr {
				t.Errorf("NetAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("NetAmount() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFuzzNetAmount(t *testing.T) {
	type args struct {
		Gross int64
		Rate  float64
	}

	fz := fuzz.New()

	// 1M iterations.
	for i := 0; i < 1e6; i++ {
		var a args
		fz.Fuzz(&a)
		// we don't really care about the values here.
		// we just want to prove that the function cannot panic.
		netAmount, err := accounting.NetAmount(a.Gross, a.Rate)
		useInt64(netAmount)
		useError(err)
	}
}

func TestFuzzTaxAmount(t *testing.T) {
	// One might call this test spurious.
	// And you'd be right.
	// But if we ever change the internals, we'll be glad to have it!
	type args struct {
		Gross int64
		Net   int64
	}

	fz := fuzz.New()

	// 1M iterations.
	for i := 0; i < 1e6; i++ {
		var a args
		fz.Fuzz(&a)

		// we don't really care about the values here.
		// we just want to prove that the function cannot panic.
		taxAmount, taxAmountError := accounting.TaxAmount(a.Gross, a.Net)
		useInt64(taxAmount)
		useError(taxAmountError)
	}
}

var minorUnitTestsCases = []struct {
	name string
	f64  float64
	i64  int64
}{
	{
		name: "known dangerous value",
		f64:  9.95,
		i64:  995,
	},
	{
		name: "known dangerous negative value",
		f64:  -9.95,
		i64:  -995,
	},
	{
		name: "f64:1.15",
		f64:  1.15,
		i64:  115,
	},
	{
		name: "f64:2.25",
		f64:  2.25,
		i64:  225,
	},
	{
		name: "f64:3.35",
		f64:  3.35,
		i64:  335,
	},
	{
		name: "f64:4.45",
		f64:  4.45,
		i64:  445,
	},
	{
		name: "f64:5.55",
		f64:  5.55,
		i64:  555,
	},
	{
		name: "f64:6.65",
		f64:  6.65,
		i64:  665,
	},
	{
		name: "f64:7.75",
		f64:  7.75,
		i64:  775,
	},
	{
		name: "f64:8.85",
		f64:  8.85,
		i64:  885,
	},
	{
		name: "f64:9.95",
		f64:  9.95,
		i64:  995,
	},
	{
		name: "f64:10.10",
		f64:  10.10,
		i64:  1010,
	},
	{
		name: "f64:11.11",
		f64:  11.11,
		i64:  1111,
	},
	{
		name: "f64:12.12",
		f64:  12.12,
		i64:  1212,
	},
	{
		name: "f64:13.13",
		f64:  13.13,
		i64:  1313,
	},
	{
		name: "f64:14.14",
		f64:  14.14,
		i64:  1414,
	},
	{
		name: "f64:15.15",
		f64:  15.15,
		i64:  1515,
	},
	{
		name: "f64:16.16",
		f64:  16.16,
		i64:  1616,
	},
	{
		name: "f64:17.17",
		f64:  17.17,
		i64:  1717,
	},
	{
		name: "f64:18.18",
		f64:  18.18,
		i64:  1818,
	},
	{
		name: "f64:19.19",
		f64:  19.19,
		i64:  1919,
	},
	{
		name: "f64:20.20",
		f64:  20.20,
		i64:  2020,
	},
	{
		name: "f64:21.21",
		f64:  21.21,
		i64:  2121,
	},
	{
		name: "f64:22.22",
		f64:  22.22,
		i64:  2222,
	},
	{
		name: "f64:23.23",
		f64:  23.23,
		i64:  2323,
	},
	{
		name: "f64:24.24",
		f64:  24.24,
		i64:  2424,
	},
	{
		name: "f64:25.25",
		f64:  25.25,
		i64:  2525,
	},
	{
		name: "f64:26.26",
		f64:  26.26,
		i64:  2626,
	},
	{
		name: "f64:27.27",
		f64:  27.27,
		i64:  2727,
	},
	{
		name: "f64:28.28",
		f64:  28.28,
		i64:  2828,
	},
	{
		name: "f64:29.29",
		f64:  29.29,
		i64:  2929,
	},
	{
		name: "f64:30.30",
		f64:  30.30,
		i64:  3030,
	},
	{
		name: "f64:31.31",
		f64:  31.31,
		i64:  3131,
	},
	{
		name: "f64:32.32",
		f64:  32.32,
		i64:  3232,
	},
	{
		name: "f64:33.33",
		f64:  33.33,
		i64:  3333,
	},
	{
		name: "f64:34.34",
		f64:  34.34,
		i64:  3434,
	},
	{
		name: "f64:35.35",
		f64:  35.35,
		i64:  3535,
	},
	{
		name: "f64:36.36",
		f64:  36.36,
		i64:  3636,
	},
	{
		name: "f64:37.37",
		f64:  37.37,
		i64:  3737,
	},
	{
		name: "f64:38.38",
		f64:  38.38,
		i64:  3838,
	},
	{
		name: "f64:39.39",
		f64:  39.39,
		i64:  3939,
	},
	{
		name: "f64:40.40",
		f64:  40.40,
		i64:  4040,
	},
	{
		name: "f64:41.41",
		f64:  41.41,
		i64:  4141,
	},
	{
		name: "f64:42.42",
		f64:  42.42,
		i64:  4242,
	},
	{
		name: "f64:43.43",
		f64:  43.43,
		i64:  4343,
	},
	{
		name: "f64:44.44",
		f64:  44.44,
		i64:  4444,
	},
	{
		name: "f64:45.45",
		f64:  45.45,
		i64:  4545,
	},
	{
		name: "f64:46.46",
		f64:  46.46,
		i64:  4646,
	},
	{
		name: "f64:47.47",
		f64:  47.47,
		i64:  4747,
	},
	{
		name: "f64:48.48",
		f64:  48.48,
		i64:  4848,
	},
	{
		name: "f64:49.49",
		f64:  49.49,
		i64:  4949,
	},
	{
		name: "f64:50.50",
		f64:  50.50,
		i64:  5050,
	},
	{
		name: "f64:51.51",
		f64:  51.51,
		i64:  5151,
	},
	{
		name: "f64:52.52",
		f64:  52.52,
		i64:  5252,
	},
	{
		name: "f64:53.53",
		f64:  53.53,
		i64:  5353,
	},
	{
		name: "f64:54.54",
		f64:  54.54,
		i64:  5454,
	},
	{
		name: "f64:55.55",
		f64:  55.55,
		i64:  5555,
	},
	{
		name: "f64:56.56",
		f64:  56.56,
		i64:  5656,
	},
	{
		name: "f64:57.57",
		f64:  57.57,
		i64:  5757,
	},
	{
		name: "f64:58.58",
		f64:  58.58,
		i64:  5858,
	},
	{
		name: "f64:59.59",
		f64:  59.59,
		i64:  5959,
	},
	{
		name: "f64:60.60",
		f64:  60.60,
		i64:  6060,
	},
	{
		name: "f64:61.61",
		f64:  61.61,
		i64:  6161,
	},
	{
		name: "f64:62.62",
		f64:  62.62,
		i64:  6262,
	},
	{
		name: "f64:63.63",
		f64:  63.63,
		i64:  6363,
	},
	{
		name: "f64:64.64",
		f64:  64.64,
		i64:  6464,
	},
	{
		name: "f64:65.65",
		f64:  65.65,
		i64:  6565,
	},
	{
		name: "f64:66.66",
		f64:  66.66,
		i64:  6666,
	},
	{
		name: "f64:67.67",
		f64:  67.67,
		i64:  6767,
	},
	{
		name: "f64:68.68",
		f64:  68.68,
		i64:  6868,
	},
	{
		name: "f64:69.69",
		f64:  69.69,
		i64:  6969,
	},
	{
		name: "f64:70.70",
		f64:  70.70,
		i64:  7070,
	},
	{
		name: "f64:71.71",
		f64:  71.71,
		i64:  7171,
	},
	{
		name: "f64:72.72",
		f64:  72.72,
		i64:  7272,
	},
	{
		name: "f64:73.73",
		f64:  73.73,
		i64:  7373,
	},
	{
		name: "f64:74.74",
		f64:  74.74,
		i64:  7474,
	},
	{
		name: "f64:75.75",
		f64:  75.75,
		i64:  7575,
	},
	{
		name: "f64:76.76",
		f64:  76.76,
		i64:  7676,
	},
	{
		name: "f64:77.77",
		f64:  77.77,
		i64:  7777,
	},
	{
		name: "f64:78.78",
		f64:  78.78,
		i64:  7878,
	},
	{
		name: "f64:79.79",
		f64:  79.79,
		i64:  7979,
	},
	{
		name: "f64:80.80",
		f64:  80.80,
		i64:  8080,
	},
	{
		name: "f64:81.81",
		f64:  81.81,
		i64:  8181,
	},
	{
		name: "f64:82.82",
		f64:  82.82,
		i64:  8282,
	},
	{
		name: "f64:83.83",
		f64:  83.83,
		i64:  8383,
	},
	{
		name: "f64:84.84",
		f64:  84.84,
		i64:  8484,
	},
	{
		name: "f64:85.85",
		f64:  85.85,
		i64:  8585,
	},
	{
		name: "f64:86.86",
		f64:  86.86,
		i64:  8686,
	},
	{
		name: "f64:87.87",
		f64:  87.87,
		i64:  8787,
	},
	{
		name: "f64:88.88",
		f64:  88.88,
		i64:  8888,
	},
	{
		name: "f64:89.89",
		f64:  89.89,
		i64:  8989,
	},
	{
		name: "f64:90.90",
		f64:  90.90,
		i64:  9090,
	},
	{
		name: "f64:91.91",
		f64:  91.91,
		i64:  9191,
	},
	{
		name: "f64:92.92",
		f64:  92.92,
		i64:  9292,
	},
	{
		name: "f64:93.93",
		f64:  93.93,
		i64:  9393,
	},
	{
		name: "f64:94.94",
		f64:  94.94,
		i64:  9494,
	},
	{
		name: "f64:95.95",
		f64:  95.95,
		i64:  9595,
	},
	{
		name: "f64:96.96",
		f64:  96.96,
		i64:  9696,
	},
	{
		name: "f64:97.97",
		f64:  97.97,
		i64:  9797,
	},
	{
		name: "f64:98.98",
		f64:  98.98,
		i64:  9898,
	},
	{
		name: "f64:99.99",
		f64:  99.99,
		i64:  9999,
	},
	{
		name: "f64:100.105",
		f64:  100.10,
		i64:  10010,
	},
	{
		name: "f64:1.15",
		f64:  -1.15,
		i64:  -115,
	},
	{
		name: "f64:2.25",
		f64:  -2.25,
		i64:  -225,
	},
	{
		name: "f64:3.35",
		f64:  -3.35,
		i64:  -335,
	},
	{
		name: "f64:4.45",
		f64:  -4.45,
		i64:  -445,
	},
	{
		name: "f64:5.55",
		f64:  -5.55,
		i64:  -555,
	},
	{
		name: "f64:6.65",
		f64:  -6.65,
		i64:  -665,
	},
	{
		name: "f64:7.75",
		f64:  -7.75,
		i64:  -775,
	},
	{
		name: "f64:8.85",
		f64:  -8.85,
		i64:  -885,
	},
	{
		name: "f64:9.95",
		f64:  -9.95,
		i64:  -995,
	},
	{
		name: "f64:10.10",
		f64:  -10.10,
		i64:  -1010,
	},
	{
		name: "f64:11.11",
		f64:  -11.11,
		i64:  -1111,
	},
	{
		name: "f64:12.12",
		f64:  -12.12,
		i64:  -1212,
	},
	{
		name: "f64:13.13",
		f64:  -13.13,
		i64:  -1313,
	},
	{
		name: "f64:14.14",
		f64:  -14.14,
		i64:  -1414,
	},
	{
		name: "f64:15.15",
		f64:  -15.15,
		i64:  -1515,
	},
	{
		name: "f64:16.16",
		f64:  -16.16,
		i64:  -1616,
	},
	{
		name: "f64:17.17",
		f64:  -17.17,
		i64:  -1717,
	},
	{
		name: "f64:18.18",
		f64:  -18.18,
		i64:  -1818,
	},
	{
		name: "f64:19.19",
		f64:  -19.19,
		i64:  -1919,
	},
	{
		name: "f64:20.20",
		f64:  -20.20,
		i64:  -2020,
	},
	{
		name: "f64:21.21",
		f64:  -21.21,
		i64:  -2121,
	},
	{
		name: "f64:22.22",
		f64:  -22.22,
		i64:  -2222,
	},
	{
		name: "f64:23.23",
		f64:  -23.23,
		i64:  -2323,
	},
	{
		name: "f64:24.24",
		f64:  -24.24,
		i64:  -2424,
	},
	{
		name: "f64:25.25",
		f64:  -25.25,
		i64:  -2525,
	},
	{
		name: "f64:26.26",
		f64:  -26.26,
		i64:  -2626,
	},
	{
		name: "f64:27.27",
		f64:  -27.27,
		i64:  -2727,
	},
	{
		name: "f64:28.28",
		f64:  -28.28,
		i64:  -2828,
	},
	{
		name: "f64:29.29",
		f64:  -29.29,
		i64:  -2929,
	},
	{
		name: "f64:30.30",
		f64:  -30.30,
		i64:  -3030,
	},
	{
		name: "f64:31.31",
		f64:  -31.31,
		i64:  -3131,
	},
	{
		name: "f64:32.32",
		f64:  -32.32,
		i64:  -3232,
	},
	{
		name: "f64:33.33",
		f64:  -33.33,
		i64:  -3333,
	},
	{
		name: "f64:34.34",
		f64:  -34.34,
		i64:  -3434,
	},
	{
		name: "f64:35.35",
		f64:  -35.35,
		i64:  -3535,
	},
	{
		name: "f64:36.36",
		f64:  -36.36,
		i64:  -3636,
	},
	{
		name: "f64:37.37",
		f64:  -37.37,
		i64:  -3737,
	},
	{
		name: "f64:38.38",
		f64:  -38.38,
		i64:  -3838,
	},
	{
		name: "f64:39.39",
		f64:  -39.39,
		i64:  -3939,
	},
	{
		name: "f64:40.40",
		f64:  -40.40,
		i64:  -4040,
	},
	{
		name: "f64:41.41",
		f64:  -41.41,
		i64:  -4141,
	},
	{
		name: "f64:42.42",
		f64:  -42.42,
		i64:  -4242,
	},
	{
		name: "f64:43.43",
		f64:  -43.43,
		i64:  -4343,
	},
	{
		name: "f64:44.44",
		f64:  -44.44,
		i64:  -4444,
	},
	{
		name: "f64:45.45",
		f64:  -45.45,
		i64:  -4545,
	},
	{
		name: "f64:46.46",
		f64:  -46.46,
		i64:  -4646,
	},
	{
		name: "f64:47.47",
		f64:  -47.47,
		i64:  -4747,
	},
	{
		name: "f64:48.48",
		f64:  -48.48,
		i64:  -4848,
	},
	{
		name: "f64:49.49",
		f64:  -49.49,
		i64:  -4949,
	},
	{
		name: "f64:50.50",
		f64:  -50.50,
		i64:  -5050,
	},
	{
		name: "f64:51.51",
		f64:  -51.51,
		i64:  -5151,
	},
	{
		name: "f64:52.52",
		f64:  -52.52,
		i64:  -5252,
	},
	{
		name: "f64:53.53",
		f64:  -53.53,
		i64:  -5353,
	},
	{
		name: "f64:54.54",
		f64:  -54.54,
		i64:  -5454,
	},
	{
		name: "f64:55.55",
		f64:  -55.55,
		i64:  -5555,
	},
	{
		name: "f64:56.56",
		f64:  -56.56,
		i64:  -5656,
	},
	{
		name: "f64:57.57",
		f64:  -57.57,
		i64:  -5757,
	},
	{
		name: "f64:58.58",
		f64:  -58.58,
		i64:  -5858,
	},
	{
		name: "f64:59.59",
		f64:  -59.59,
		i64:  -5959,
	},
	{
		name: "f64:60.60",
		f64:  -60.60,
		i64:  -6060,
	},
	{
		name: "f64:61.61",
		f64:  -61.61,
		i64:  -6161,
	},
	{
		name: "f64:62.62",
		f64:  -62.62,
		i64:  -6262,
	},
	{
		name: "f64:63.63",
		f64:  -63.63,
		i64:  -6363,
	},
	{
		name: "f64:64.64",
		f64:  -64.64,
		i64:  -6464,
	},
	{
		name: "f64:65.65",
		f64:  -65.65,
		i64:  -6565,
	},
	{
		name: "f64:66.66",
		f64:  -66.66,
		i64:  -6666,
	},
	{
		name: "f64:67.67",
		f64:  -67.67,
		i64:  -6767,
	},
	{
		name: "f64:68.68",
		f64:  -68.68,
		i64:  -6868,
	},
	{
		name: "f64:69.69",
		f64:  -69.69,
		i64:  -6969,
	},
	{
		name: "f64:70.70",
		f64:  -70.70,
		i64:  -7070,
	},
	{
		name: "f64:71.71",
		f64:  -71.71,
		i64:  -7171,
	},
	{
		name: "f64:72.72",
		f64:  -72.72,
		i64:  -7272,
	},
	{
		name: "f64:73.73",
		f64:  -73.73,
		i64:  -7373,
	},
	{
		name: "f64:74.74",
		f64:  -74.74,
		i64:  -7474,
	},
	{
		name: "f64:75.75",
		f64:  -75.75,
		i64:  -7575,
	},
	{
		name: "f64:76.76",
		f64:  -76.76,
		i64:  -7676,
	},
	{
		name: "f64:77.77",
		f64:  -77.77,
		i64:  -7777,
	},
	{
		name: "f64:78.78",
		f64:  -78.78,
		i64:  -7878,
	},
	{
		name: "f64:79.79",
		f64:  -79.79,
		i64:  -7979,
	},
	{
		name: "f64:80.80",
		f64:  -80.80,
		i64:  -8080,
	},
	{
		name: "f64:81.81",
		f64:  -81.81,
		i64:  -8181,
	},
	{
		name: "f64:82.82",
		f64:  -82.82,
		i64:  -8282,
	},
	{
		name: "f64:83.83",
		f64:  -83.83,
		i64:  -8383,
	},
	{
		name: "f64:84.84",
		f64:  -84.84,
		i64:  -8484,
	},
	{
		name: "f64:85.85",
		f64:  -85.85,
		i64:  -8585,
	},
	{
		name: "f64:86.86",
		f64:  -86.86,
		i64:  -8686,
	},
	{
		name: "f64:87.87",
		f64:  -87.87,
		i64:  -8787,
	},
	{
		name: "f64:88.88",
		f64:  -88.88,
		i64:  -8888,
	},
	{
		name: "f64:89.89",
		f64:  -89.89,
		i64:  -8989,
	},
	{
		name: "f64:90.90",
		f64:  -90.90,
		i64:  -9090,
	},
	{
		name: "f64:91.91",
		f64:  -91.91,
		i64:  -9191,
	},
	{
		name: "f64:92.92",
		f64:  -92.92,
		i64:  -9292,
	},
	{
		name: "f64:93.93",
		f64:  -93.93,
		i64:  -9393,
	},
	{
		name: "f64:94.94",
		f64:  -94.94,
		i64:  -9494,
	},
	{
		name: "f64:95.95",
		f64:  -95.95,
		i64:  -9595,
	},
	{
		name: "f64:96.96",
		f64:  -96.96,
		i64:  -9696,
	},
	{
		name: "f64:97.97",
		f64:  -97.97,
		i64:  -9797,
	},
	{
		name: "f64:98.98",
		f64:  -98.98,
		i64:  -9898,
	},
	{
		name: "f64:99.99",
		f64:  -99.99,
		i64:  -9999,
	},
	{
		name: "f64:100.105",
		f64:  -100.10,
		i64:  -10010,
	},
}

//go:noinline
func useInt64(i int64) {}

//go:noinline
func useError(e error) {}
