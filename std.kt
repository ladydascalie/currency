data class Currency(val currencyCode: String, val minorUnits: Int, val factor: Int) {
    companion object {
        val currencyMap = mapOf<String, Currency>(
        "AED" to Currency("AED", 2, 100),
        "AFN" to Currency("AFN", 2, 100),
        "ALL" to Currency("ALL", 2, 100),
        "AMD" to Currency("AMD", 2, 100),
        "ANG" to Currency("ANG", 2, 100),
        "AOA" to Currency("AOA", 2, 100),
        "ARS" to Currency("ARS", 2, 100),
        "AUD" to Currency("AUD", 2, 100),
        "AWG" to Currency("AWG", 2, 100),
        "AZN" to Currency("AZN", 2, 100),
        "BAM" to Currency("BAM", 2, 100),
        "BBD" to Currency("BBD", 2, 100),
        "BDT" to Currency("BDT", 2, 100),
        "BGN" to Currency("BGN", 2, 100),
        "BHD" to Currency("BHD", 3, 1000),
        "BIF" to Currency("BIF", 0, 1),
        "BMD" to Currency("BMD", 2, 100),
        "BND" to Currency("BND", 2, 100),
        "BOB" to Currency("BOB", 2, 100),
        "BOV" to Currency("BOV", 2, 100),
        "BRL" to Currency("BRL", 2, 100),
        "BSD" to Currency("BSD", 2, 100),
        "BTN" to Currency("BTN", 2, 100),
        "BWP" to Currency("BWP", 2, 100),
        "BYN" to Currency("BYN", 2, 100),
        "BZD" to Currency("BZD", 2, 100),
        "CAD" to Currency("CAD", 2, 100),
        "CDF" to Currency("CDF", 2, 100),
        "CHE" to Currency("CHE", 2, 100),
        "CHF" to Currency("CHF", 2, 100),
        "CHW" to Currency("CHW", 2, 100),
        "CLF" to Currency("CLF", 4, 10000),
        "CLP" to Currency("CLP", 0, 1),
        "CNY" to Currency("CNY", 2, 100),
        "COP" to Currency("COP", 2, 100),
        "COU" to Currency("COU", 2, 100),
        "CRC" to Currency("CRC", 2, 100),
        "CUC" to Currency("CUC", 2, 100),
        "CUP" to Currency("CUP", 2, 100),
        "CVE" to Currency("CVE", 2, 100),
        "CZK" to Currency("CZK", 2, 100),
        "DJF" to Currency("DJF", 0, 1),
        "DKK" to Currency("DKK", 2, 100),
        "DOP" to Currency("DOP", 2, 100),
        "DZD" to Currency("DZD", 2, 100),
        "EGP" to Currency("EGP", 2, 100),
        "ERN" to Currency("ERN", 2, 100),
        "ETB" to Currency("ETB", 2, 100),
        "EUR" to Currency("EUR", 2, 100),
        "FJD" to Currency("FJD", 2, 100),
        "FKP" to Currency("FKP", 2, 100),
        "GBP" to Currency("GBP", 2, 100),
        "GEL" to Currency("GEL", 2, 100),
        "GHS" to Currency("GHS", 2, 100),
        "GIP" to Currency("GIP", 2, 100),
        "GMD" to Currency("GMD", 2, 100),
        "GNF" to Currency("GNF", 0, 1),
        "GTQ" to Currency("GTQ", 2, 100),
        "GYD" to Currency("GYD", 2, 100),
        "HKD" to Currency("HKD", 2, 100),
        "HNL" to Currency("HNL", 2, 100),
        "HTG" to Currency("HTG", 2, 100),
        "HUF" to Currency("HUF", 2, 100),
        "IDR" to Currency("IDR", 2, 100),
        "ILS" to Currency("ILS", 2, 100),
        "INR" to Currency("INR", 2, 100),
        "IQD" to Currency("IQD", 3, 1000),
        "IRR" to Currency("IRR", 2, 100),
        "ISK" to Currency("ISK", 0, 1),
        "JMD" to Currency("JMD", 2, 100),
        "JOD" to Currency("JOD", 3, 1000),
        "JPY" to Currency("JPY", 0, 1),
        "KES" to Currency("KES", 2, 100),
        "KGS" to Currency("KGS", 2, 100),
        "KHR" to Currency("KHR", 2, 100),
        "KMF" to Currency("KMF", 0, 1),
        "KPW" to Currency("KPW", 2, 100),
        "KRW" to Currency("KRW", 0, 1),
        "KWD" to Currency("KWD", 3, 1000),
        "KYD" to Currency("KYD", 2, 100),
        "KZT" to Currency("KZT", 2, 100),
        "LAK" to Currency("LAK", 2, 100),
        "LBP" to Currency("LBP", 2, 100),
        "LKR" to Currency("LKR", 2, 100),
        "LRD" to Currency("LRD", 2, 100),
        "LSL" to Currency("LSL", 2, 100),
        "LYD" to Currency("LYD", 3, 1000),
        "MAD" to Currency("MAD", 2, 100),
        "MDL" to Currency("MDL", 2, 100),
        "MGA" to Currency("MGA", 2, 100),
        "MKD" to Currency("MKD", 2, 100),
        "MMK" to Currency("MMK", 2, 100),
        "MNT" to Currency("MNT", 2, 100),
        "MOP" to Currency("MOP", 2, 100),
        "MRU" to Currency("MRU", 2, 100),
        "MUR" to Currency("MUR", 2, 100),
        "MVR" to Currency("MVR", 2, 100),
        "MWK" to Currency("MWK", 2, 100),
        "MXN" to Currency("MXN", 2, 100),
        "MXV" to Currency("MXV", 2, 100),
        "MYR" to Currency("MYR", 2, 100),
        "MZN" to Currency("MZN", 2, 100),
        "NAD" to Currency("NAD", 2, 100),
        "NGN" to Currency("NGN", 2, 100),
        "NIO" to Currency("NIO", 2, 100),
        "NOK" to Currency("NOK", 2, 100),
        "NPR" to Currency("NPR", 2, 100),
        "NZD" to Currency("NZD", 2, 100),
        "OMR" to Currency("OMR", 3, 1000),
        "PAB" to Currency("PAB", 2, 100),
        "PEN" to Currency("PEN", 2, 100),
        "PGK" to Currency("PGK", 2, 100),
        "PHP" to Currency("PHP", 2, 100),
        "PKR" to Currency("PKR", 2, 100),
        "PLN" to Currency("PLN", 2, 100),
        "PYG" to Currency("PYG", 0, 1),
        "QAR" to Currency("QAR", 2, 100),
        "RON" to Currency("RON", 2, 100),
        "RSD" to Currency("RSD", 2, 100),
        "RUB" to Currency("RUB", 2, 100),
        "RWF" to Currency("RWF", 0, 1),
        "SAR" to Currency("SAR", 2, 100),
        "SBD" to Currency("SBD", 2, 100),
        "SCR" to Currency("SCR", 2, 100),
        "SDG" to Currency("SDG", 2, 100),
        "SEK" to Currency("SEK", 2, 100),
        "SGD" to Currency("SGD", 2, 100),
        "SHP" to Currency("SHP", 2, 100),
        "SLE" to Currency("SLE", 2, 100),
        "SOS" to Currency("SOS", 2, 100),
        "SRD" to Currency("SRD", 2, 100),
        "SSP" to Currency("SSP", 2, 100),
        "STN" to Currency("STN", 2, 100),
        "SVC" to Currency("SVC", 2, 100),
        "SYP" to Currency("SYP", 2, 100),
        "SZL" to Currency("SZL", 2, 100),
        "THB" to Currency("THB", 2, 100),
        "TJS" to Currency("TJS", 2, 100),
        "TMT" to Currency("TMT", 2, 100),
        "TND" to Currency("TND", 3, 1000),
        "TOP" to Currency("TOP", 2, 100),
        "TRY" to Currency("TRY", 2, 100),
        "TTD" to Currency("TTD", 2, 100),
        "TWD" to Currency("TWD", 2, 100),
        "TZS" to Currency("TZS", 2, 100),
        "UAH" to Currency("UAH", 2, 100),
        "UGX" to Currency("UGX", 0, 1),
        "USD" to Currency("USD", 2, 100),
        "USN" to Currency("USN", 2, 100),
        "UYI" to Currency("UYI", 0, 1),
        "UYU" to Currency("UYU", 2, 100),
        "UYW" to Currency("UYW", 4, 10000),
        "UZS" to Currency("UZS", 2, 100),
        "VED" to Currency("VED", 2, 100),
        "VES" to Currency("VES", 2, 100),
        "VND" to Currency("VND", 0, 1),
        "VUV" to Currency("VUV", 0, 1),
        "WST" to Currency("WST", 2, 100),
        "XAF" to Currency("XAF", 0, 1),
        "XAG" to Currency("XAG", 0, 1),
        "XAU" to Currency("XAU", 0, 1),
        "XBA" to Currency("XBA", 0, 1),
        "XBB" to Currency("XBB", 0, 1),
        "XBC" to Currency("XBC", 0, 1),
        "XBD" to Currency("XBD", 0, 1),
        "XCD" to Currency("XCD", 2, 100),
        "XDR" to Currency("XDR", 0, 1),
        "XOF" to Currency("XOF", 0, 1),
        "XPD" to Currency("XPD", 0, 1),
        "XPF" to Currency("XPF", 0, 1),
        "XPT" to Currency("XPT", 0, 1),
        "XSU" to Currency("XSU", 0, 1),
        "XTS" to Currency("XTS", 0, 1),
        "XUA" to Currency("XUA", 0, 1),
        "XXX" to Currency("XXX", 0, 1),
        "YER" to Currency("YER", 2, 100),
        "ZAR" to Currency("ZAR", 2, 100),
        "ZMW" to Currency("ZMW", 2, 100),
        "ZWL" to Currency("ZWL", 2, 100)
        )
    }
}