// Country names and ISO codes kit.
package locales

import (
	"errors"
	"strings"

	"golang.org/x/text/language"
)

type Locale string

// All available ISO codes.
const (
	LocaleWildcard Locale = "*"
	LocaleAD       Locale = "AD"
	LocaleAF       Locale = "AF"
	LocaleAG       Locale = "AG"
	LocaleAI       Locale = "AI"
	LocaleAL       Locale = "AL"
	LocaleAM       Locale = "AM"
	LocaleAO       Locale = "AO"
	LocaleAQ       Locale = "AQ"
	LocaleAE       Locale = "AE"
	LocaleAR       Locale = "AR"
	LocaleAS       Locale = "AS"
	LocaleAT       Locale = "AT"
	LocaleAU       Locale = "AU"
	LocaleAW       Locale = "AW"
	LocaleAX       Locale = "AX"
	LocaleAZ       Locale = "AZ"
	LocaleBA       Locale = "BA"
	LocaleBB       Locale = "BB"
	LocaleBD       Locale = "BD"
	LocaleBE       Locale = "BE"
	LocaleBF       Locale = "BF"
	LocaleBG       Locale = "BG"
	LocaleBH       Locale = "BH"
	LocaleBI       Locale = "BI"
	LocaleBJ       Locale = "BJ"
	LocaleBL       Locale = "BL"
	LocaleBM       Locale = "BM"
	LocaleBN       Locale = "BN"
	LocaleBO       Locale = "BO"
	LocaleBQ       Locale = "BQ"
	LocaleBR       Locale = "BR"
	LocaleBS       Locale = "BS"
	LocaleBT       Locale = "BT"
	LocaleBV       Locale = "BV"
	LocaleBW       Locale = "BW"
	LocaleBY       Locale = "BY"
	LocaleBZ       Locale = "BZ"
	LocaleCA       Locale = "CA"
	LocaleCC       Locale = "CC"
	LocaleCD       Locale = "CD"
	LocaleCF       Locale = "CF"
	LocaleCG       Locale = "CG"
	LocaleCH       Locale = "CH"
	LocaleCI       Locale = "CI"
	LocaleCK       Locale = "CK"
	LocaleCL       Locale = "CL"
	LocaleCM       Locale = "CM"
	LocaleCN       Locale = "CN"
	LocaleCO       Locale = "CO"
	LocaleCR       Locale = "CR"
	LocaleCU       Locale = "CU"
	LocaleCV       Locale = "CV"
	LocaleCW       Locale = "CW"
	LocaleCX       Locale = "CX"
	LocaleCY       Locale = "CY"
	LocaleCZ       Locale = "CZ"
	LocaleDE       Locale = "DE"
	LocaleDJ       Locale = "DJ"
	LocaleDK       Locale = "DK"
	LocaleDM       Locale = "DM"
	LocaleDO       Locale = "DO"
	LocaleDZ       Locale = "DZ"
	LocaleEC       Locale = "EC"
	LocaleEE       Locale = "EE"
	LocaleEG       Locale = "EG"
	LocaleEH       Locale = "EH"
	LocaleER       Locale = "ER"
	LocaleES       Locale = "ES"
	LocaleET       Locale = "ET"
	LocaleFI       Locale = "FI"
	LocaleFJ       Locale = "FJ"
	LocaleFK       Locale = "FK"
	LocaleFM       Locale = "FM"
	LocaleFO       Locale = "FO"
	LocaleFR       Locale = "FR"
	LocaleGA       Locale = "GA"
	LocaleGB       Locale = "GB"
	LocaleGD       Locale = "GD"
	LocaleGE       Locale = "GE"
	LocaleGF       Locale = "GF"
	LocaleGG       Locale = "GG"
	LocaleGH       Locale = "GH"
	LocaleGI       Locale = "GI"
	LocaleGL       Locale = "GL"
	LocaleGM       Locale = "GM"
	LocaleGN       Locale = "GN"
	LocaleGP       Locale = "GP"
	LocaleGQ       Locale = "GQ"
	LocaleGR       Locale = "GR"
	LocaleGS       Locale = "GS"
	LocaleGT       Locale = "GT"
	LocaleGU       Locale = "GU"
	LocaleGW       Locale = "GW"
	LocaleGY       Locale = "GY"
	LocaleHK       Locale = "HK"
	LocaleHM       Locale = "HM"
	LocaleHN       Locale = "HN"
	LocaleHR       Locale = "HR"
	LocaleHT       Locale = "HT"
	LocaleHU       Locale = "HU"
	LocaleID       Locale = "ID"
	LocaleIE       Locale = "IE"
	LocaleIL       Locale = "IL"
	LocaleIM       Locale = "IM"
	LocaleIN       Locale = "IN"
	LocaleIO       Locale = "IO"
	LocaleIQ       Locale = "IQ"
	LocaleIR       Locale = "IR"
	LocaleIS       Locale = "IS"
	LocaleIT       Locale = "IT"
	LocaleJE       Locale = "JE"
	LocaleJM       Locale = "JM"
	LocaleJO       Locale = "JO"
	LocaleJP       Locale = "JP"
	LocaleKE       Locale = "KE"
	LocaleKG       Locale = "KG"
	LocaleKH       Locale = "KH"
	LocaleKI       Locale = "KI"
	LocaleKM       Locale = "KM"
	LocaleKN       Locale = "KN"
	LocaleKP       Locale = "KP"
	LocaleKR       Locale = "KR"
	LocaleKW       Locale = "KW"
	LocaleKY       Locale = "KY"
	LocaleKZ       Locale = "KZ"
	LocaleLA       Locale = "LA"
	LocaleLB       Locale = "LB"
	LocaleLC       Locale = "LC"
	LocaleLI       Locale = "LI"
	LocaleLK       Locale = "LK"
	LocaleLR       Locale = "LR"
	LocaleLS       Locale = "LS"
	LocaleLT       Locale = "LT"
	LocaleLU       Locale = "LU"
	LocaleLV       Locale = "LV"
	LocaleLY       Locale = "LY"
	LocaleMA       Locale = "MA"
	LocaleMC       Locale = "MC"
	LocaleMD       Locale = "MD"
	LocaleME       Locale = "ME"
	LocaleMF       Locale = "MF"
	LocaleMG       Locale = "MG"
	LocaleMH       Locale = "MH"
	LocaleMK       Locale = "MK"
	LocaleML       Locale = "ML"
	LocaleMM       Locale = "MM"
	LocaleMN       Locale = "MN"
	LocaleMO       Locale = "MO"
	LocaleMP       Locale = "MP"
	LocaleMQ       Locale = "MQ"
	LocaleMR       Locale = "MR"
	LocaleMS       Locale = "MS"
	LocaleMT       Locale = "MT"
	LocaleMU       Locale = "MU"
	LocaleMV       Locale = "MV"
	LocaleMW       Locale = "MW"
	LocaleMX       Locale = "MX"
	LocaleMY       Locale = "MY"
	LocaleMZ       Locale = "MZ"
	LocaleNA       Locale = "NA"
	LocaleNC       Locale = "NC"
	LocaleNE       Locale = "NE"
	LocaleNF       Locale = "NF"
	LocaleNI       Locale = "NI"
	LocaleNL       Locale = "NL"
	LocaleNO       Locale = "NO"
	LocaleNP       Locale = "NP"
	LocaleNR       Locale = "NR"
	LocaleNU       Locale = "NU"
	LocaleNZ       Locale = "NZ"
	LocaleOM       Locale = "OM"
	LocalePA       Locale = "PA"
	LocalePE       Locale = "PE"
	LocalePF       Locale = "PF"
	LocalePG       Locale = "PG"
	LocalePH       Locale = "PH"
	LocalePK       Locale = "PK"
	LocalePL       Locale = "PL"
	LocalePM       Locale = "PM"
	LocalePN       Locale = "PN"
	LocalePR       Locale = "PR"
	LocalePS       Locale = "PS"
	LocalePT       Locale = "PT"
	LocalePW       Locale = "PW"
	LocalePY       Locale = "PY"
	LocaleQA       Locale = "QA"
	LocaleRE       Locale = "RE"
	LocaleRO       Locale = "RO"
	LocaleRS       Locale = "RS"
	LocaleRU       Locale = "RU"
	LocaleRW       Locale = "RW"
	LocaleSA       Locale = "SA"
	LocaleSB       Locale = "SB"
	LocaleSC       Locale = "SC"
	LocaleSD       Locale = "SD"
	LocaleSE       Locale = "SE"
	LocaleSG       Locale = "SG"
	LocaleSH       Locale = "SH"
	LocaleSI       Locale = "SI"
	LocaleSJ       Locale = "SJ"
	LocaleSK       Locale = "SK"
	LocaleSL       Locale = "SL"
	LocaleSM       Locale = "SM"
	LocaleSN       Locale = "SN"
	LocaleSO       Locale = "SO"
	LocaleSR       Locale = "SR"
	LocaleSS       Locale = "SS"
	LocaleST       Locale = "ST"
	LocaleSV       Locale = "SV"
	LocaleSX       Locale = "SX"
	LocaleSY       Locale = "SY"
	LocaleSZ       Locale = "SZ"
	LocaleTC       Locale = "TC"
	LocaleTD       Locale = "TD"
	LocaleTF       Locale = "TF"
	LocaleTG       Locale = "TG"
	LocaleTH       Locale = "TH"
	LocaleTJ       Locale = "TJ"
	LocaleTK       Locale = "TK"
	LocaleTL       Locale = "TL"
	LocaleTM       Locale = "TM"
	LocaleTN       Locale = "TN"
	LocaleTO       Locale = "TO"
	LocaleTR       Locale = "TR"
	LocaleTT       Locale = "TT"
	LocaleTV       Locale = "TV"
	LocaleTW       Locale = "TW"
	LocaleTZ       Locale = "TZ"
	LocaleUA       Locale = "UA"
	LocaleUG       Locale = "UG"
	LocaleUM       Locale = "UM"
	LocaleUS       Locale = "US"
	LocaleUY       Locale = "UY"
	LocaleUZ       Locale = "UZ"
	LocaleVA       Locale = "VA"
	LocaleVC       Locale = "VC"
	LocaleVE       Locale = "VE"
	LocaleVG       Locale = "VG"
	LocaleVI       Locale = "VI"
	LocaleVN       Locale = "VN"
	LocaleVU       Locale = "VU"
	LocaleWF       Locale = "WF"
	LocaleWS       Locale = "WS"
	LocaleYE       Locale = "YE"
	LocaleYT       Locale = "YT"
	LocaleZA       Locale = "ZA"
	LocaleZM       Locale = "ZM"
	LocaleZW       Locale = "ZW"
)

var caountriesMap = map[Locale]string{
	LocaleAD: "Andorra",
	LocaleAE: "United Arab Emirates",
	LocaleAF: "Afghanistan",
	LocaleAG: "Antigua and Barbuda",
	LocaleAI: "Anguilla",
	LocaleAL: "Albania",
	LocaleAM: "Armenia",
	LocaleAO: "Angola",
	LocaleAQ: "Antarctica",
	LocaleAR: "Argentina",
	LocaleAS: "American Samoa",
	LocaleAT: "Austria",
	LocaleAU: "Australia",
	LocaleAW: "Aruba",
	LocaleAX: "Åland Islands",
	LocaleAZ: "Azerbaijan",
	LocaleBA: "Bosnia and Herzegovina",
	LocaleBB: "Barbados",
	LocaleBD: "Bangladesh",
	LocaleBE: "Belgium",
	LocaleBF: "Burkina Faso",
	LocaleBG: "Bulgaria",
	LocaleBH: "Bahrain",
	LocaleBI: "Burundi",
	LocaleBJ: "Benin",
	LocaleBL: "Saint Barthélemy",
	LocaleBM: "Bermuda",
	LocaleBN: "Brunei Darussalam",
	LocaleBO: "Bolivia",
	LocaleBQ: "Bonaire, Sint Eustatius and Saba",
	LocaleBR: "Brazil",
	LocaleBS: "Bahamas",
	LocaleBT: "Bhutan",
	LocaleBV: "Bouvet Island",
	LocaleBW: "Botswana",
	LocaleBY: "Belarus",
	LocaleBZ: "Belize",
	LocaleCA: "Canada",
	LocaleCC: "Cocos (Keeling) Islands",
	LocaleCD: "Democratic Republic of the Congo",
	LocaleCF: "Central African Republic",
	LocaleCG: "Congo",
	LocaleCH: "Switzerland",
	LocaleCI: "Côte d'Ivoire",
	LocaleCK: "Cook Islands",
	LocaleCL: "Chile",
	LocaleCM: "Cameroon",
	LocaleCN: "China",
	LocaleCO: "Colombia",
	LocaleCR: "Costa Rica",
	LocaleCU: "Cuba",
	LocaleCV: "Cabo Verde",
	LocaleCW: "Curaçao",
	LocaleCX: "Christmas Island",
	LocaleCY: "Cyprus",
	LocaleCZ: "Czech Republic",
	LocaleDE: "Germany",
	LocaleDJ: "Djibouti",
	LocaleDK: "Denmark",
	LocaleDM: "Dominica",
	LocaleDO: "Dominican Republic",
	LocaleDZ: "Algeria",
	LocaleEC: "Ecuador",
	LocaleEE: "Estonia",
	LocaleEG: "Egypt",
	LocaleEH: "Western Sahara",
	LocaleER: "Eritrea",
	LocaleES: "Spain",
	LocaleET: "Ethiopia",
	LocaleFI: "Finland",
	LocaleFJ: "Fiji",
	LocaleFK: "Falkland Islands (Malvinas)",
	LocaleFM: "Federated States of Micronesia",
	LocaleFO: "Faroe Islands",
	LocaleFR: "France",
	LocaleGA: "Gabon",
	LocaleGB: "United Kingdom",
	LocaleGD: "Grenada",
	LocaleGE: "Georgia",
	LocaleGF: "French Guiana",
	LocaleGG: "Guernsey",
	LocaleGH: "Ghana",
	LocaleGI: "Gibraltar",
	LocaleGL: "Greenland",
	LocaleGM: "Gambia",
	LocaleGN: "Guinea",
	LocaleGP: "Guadeloupe",
	LocaleGQ: "Equatorial Guinea",
	LocaleGR: "Greece",
	LocaleGS: "South Georgia and the South Sandwich Islands",
	LocaleGT: "Guatemala",
	LocaleGU: "Guam",
	LocaleGW: "Guinea-Bissau",
	LocaleGY: "Guyana",
	LocaleHK: "Hong Kong",
	LocaleHM: "Heard Island and McDonald Islands",
	LocaleHN: "Honduras",
	LocaleHR: "Croatia",
	LocaleHT: "Haiti",
	LocaleHU: "Hungary",
	LocaleID: "Indonesia",
	LocaleIE: "Ireland",
	LocaleIL: "Israel",
	LocaleIM: "Isle of Man",
	LocaleIN: "India",
	LocaleIO: "British Indian Ocean Territory",
	LocaleIQ: "Iraq",
	LocaleIR: "Iran",
	LocaleIS: "Iceland",
	LocaleIT: "Italy",
	LocaleJE: "Jersey",
	LocaleJM: "Jamaica",
	LocaleJO: "Jordan",
	LocaleJP: "Japan",
	LocaleKE: "Kenya",
	LocaleKG: "Kyrgyzstan",
	LocaleKH: "Cambodia",
	LocaleKI: "Kiribati",
	LocaleKM: "Comoros",
	LocaleKN: "Saint Kitts and Nevis",
	LocaleKP: "Korea",
	LocaleKR: "Korea",
	LocaleKW: "Kuwait",
	LocaleKY: "Cayman Islands",
	LocaleKZ: "Kazakhstan",
	LocaleLA: "Lao",
	LocaleLB: "Lebanon",
	LocaleLC: "Saint Lucia",
	LocaleLI: "Liechtenstein",
	LocaleLK: "Sri Lanka",
	LocaleLR: "Liberia",
	LocaleLS: "Lesotho",
	LocaleLT: "Lithuania",
	LocaleLU: "Luxembourg",
	LocaleLV: "Latvia",
	LocaleLY: "Libya",
	LocaleMA: "Morocco",
	LocaleMC: "Monaco",
	LocaleMD: "Moldova",
	LocaleME: "Montenegro",
	LocaleMF: "Saint Martin",
	LocaleMG: "Madagascar",
	LocaleMH: "Marshall Islands",
	LocaleMK: "North Macedonia",
	LocaleML: "Mali",
	LocaleMM: "Myanmar",
	LocaleMN: "Mongolia",
	LocaleMO: "Macao",
	LocaleMP: "Northern Mariana Islands",
	LocaleMQ: "Martinique",
	LocaleMR: "Mauritania",
	LocaleMS: "Montserrat",
	LocaleMT: "Malta",
	LocaleMU: "Mauritius",
	LocaleMV: "Maldives",
	LocaleMW: "Malawi",
	LocaleMX: "Mexico",
	LocaleMY: "Malaysia",
	LocaleMZ: "Mozambique",
	LocaleNA: "Namibia",
	LocaleNC: "New Caledonia",
	LocaleNE: "Niger",
	LocaleNF: "Norfolk Island",
	LocaleNI: "Nicaragua",
	LocaleNL: "Netherlands",
	LocaleNO: "Norway",
	LocaleNP: "Nepal",
	LocaleNR: "Nauru",
	LocaleNU: "Niue",
	LocaleNZ: "New Zealand",
	LocaleOM: "Oman",
	LocalePA: "Panama",
	LocalePE: "Peru",
	LocalePF: "French Polynesia",
	LocalePG: "Papua New Guinea",
	LocalePH: "Philippines",
	LocalePK: "Pakistan",
	LocalePL: "Poland",
	LocalePM: "Saint Pierre and Miquelon",
	LocalePN: "Pitcairn",
	LocalePR: "Puerto Rico",
	LocalePS: "Palestine",
	LocalePT: "Portugal",
	LocalePW: "Palau",
	LocalePY: "Paraguay",
	LocaleQA: "Qatar",
	LocaleRE: "Réunion",
	LocaleRO: "Romania",
	LocaleRS: "Serbia",
	LocaleRU: "Russian Federation",
	LocaleRW: "Rwanda",
	LocaleSA: "Saudi Arabia",
	LocaleSB: "Solomon Islands",
	LocaleSC: "Seychelles",
	LocaleSD: "Sudan",
	LocaleSE: "Sweden",
	LocaleSG: "Singapore",
	LocaleSH: "Saint Helena",
	LocaleSI: "Slovenia",
	LocaleSJ: "Svalbard and Jan Mayen",
	LocaleSK: "Slovakia",
	LocaleSL: "Sierra Leone",
	LocaleSM: "San Marino",
	LocaleSN: "Senegal",
	LocaleSO: "Somalia",
	LocaleSR: "Suriname",
	LocaleSS: "South Sudan",
	LocaleST: "Sao Tome and Principe",
	LocaleSV: "El Salvador",
	LocaleSX: "Sint Maarten (Dutch part)",
	LocaleSY: "Syrian Arab Republic",
	LocaleSZ: "Eswatini",
	LocaleTC: "Turks and Caicos Islands",
	LocaleTD: "Chad",
	LocaleTF: "French Southern Territories",
	LocaleTG: "Togo",
	LocaleTH: "Thailand",
	LocaleTJ: "Tajikistan",
	LocaleTK: "Tokelau",
	LocaleTL: "Timor-Leste",
	LocaleTM: "Turkmenistan",
	LocaleTN: "Tunisia",
	LocaleTO: "Tonga",
	LocaleTR: "Turkey",
	LocaleTT: "Trinidad and Tobago",
	LocaleTV: "Tuvalu",
	LocaleTW: "Taiwan",
	LocaleTZ: "Tanzania",
	LocaleUA: "Ukraine",
	LocaleUG: "Uganda",
	LocaleUM: "United States Minor Outlying Islands",
	LocaleUS: "United States",
	LocaleUY: "Uruguay",
	LocaleUZ: "Uzbekistan",
	LocaleVA: "Holy See (Vatican City State)",
	LocaleVC: "Saint Vincent and the Grenadines",
	LocaleVE: "Venezuela",
	LocaleVG: "Virgin Islands, British",
	LocaleVI: "Virgin Islands, U.S.",
	LocaleVN: "Viet Nam",
	LocaleVU: "Vanuatu",
	LocaleWF: "Wallis and Futuna",
	LocaleWS: "Samoa",
	LocaleYE: "Yemen",
	LocaleYT: "Mayotte",
	LocaleZA: "South Africa",
	LocaleZM: "Zambia",
	LocaleZW: "Zimbabwe",
}

// Check if such ISO country code exists.
func (l *Locale) IsValid() bool {
	for k := range caountriesMap {
		if k == *l {
			return true
		}
	}
	return false
}

// Returns country name for the ISO code.
func (l *Locale) Country() string {
	c, ok := caountriesMap[*l]
	if ok {
		return c
	} else if *l == LocaleWildcard {
		return "Global"
	}
	return ""
}

// Returns country ISO code as normalized string.
func (l *Locale) ISO() string {
	return strings.ToLower(string(*l))
}

// Safely cast string type ISO code to Locale type.
// If casting failed, en empty (zero value) Locale and error reurned.
func Make(isoCode string) (Locale, error) {
	isoCode = strings.TrimSpace(isoCode)
	isoCode = strings.ToUpper(isoCode)
	locale := Locale(isoCode)
	if locale.IsValid() {
		return locale, nil
	}
	parts := strings.Split(isoCode, ",")
	tag, err := language.Parse(parts[0])
	if err != nil {
		return "", err
	}
	reg, _ := tag.Region()
	locale = Locale(reg.String())
	if locale.IsValid() {
		return locale, nil
	}
	return "", errors.New("invalid ISO code")
}
