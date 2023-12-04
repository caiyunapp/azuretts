package azuretts

// https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/ai-services/speech-service/includes/language-support/tts.md
type Language string

const (
	LanguageAfZA         Language = "af-ZA"
	LanguageAmET         Language = "am-ET"
	LanguageArAE         Language = "ar-AE"
	LanguageArBH         Language = "ar-BH"
	LanguageArDZ         Language = "ar-DZ"
	LanguageArEG         Language = "ar-EG"
	LanguageArIQ         Language = "ar-IQ"
	LanguageArJO         Language = "ar-JO"
	LanguageArKW         Language = "ar-KW"
	LanguageArLB         Language = "ar-LB"
	LanguageArLY         Language = "ar-LY"
	LanguageArMA         Language = "ar-MA"
	LanguageArOM         Language = "ar-OM"
	LanguageArQA         Language = "ar-QA"
	LanguageArSA         Language = "ar-SA"
	LanguageArSY         Language = "ar-SY"
	LanguageArTN         Language = "ar-TN"
	LanguageArYE         Language = "ar-YE"
	LanguageAzAZ         Language = "az-AZ"
	LanguageBgBG         Language = "bg-BG"
	LanguageBnBD         Language = "bn-BD"
	LanguageBnIN         Language = "bn-IN"
	LanguageBsBA         Language = "bs-BA"
	LanguageCaES         Language = "ca-ES"
	LanguageCsCZ         Language = "cs-CZ"
	LanguageCyGB         Language = "cy-GB"
	LanguageDaDK         Language = "da-DK"
	LanguageDeAT         Language = "de-AT"
	LanguageDeCH         Language = "de-CH"
	LanguageDeDE         Language = "de-DE"
	LanguageElGR         Language = "el-GR"
	LanguageEnAU         Language = "en-AU"
	LanguageEnCA         Language = "en-CA"
	LanguageEnGB         Language = "en-GB"
	LanguageEnHK         Language = "en-HK"
	LanguageEnIE         Language = "en-IE"
	LanguageEnIN         Language = "en-IN"
	LanguageEnKE         Language = "en-KE"
	LanguageEnNG         Language = "en-NG"
	LanguageEnNZ         Language = "en-NZ"
	LanguageEnPH         Language = "en-PH"
	LanguageEnSG         Language = "en-SG"
	LanguageEnTZ         Language = "en-TZ"
	LanguageEnUS         Language = "en-US"
	LanguageEnZA         Language = "en-ZA"
	LanguageEsAR         Language = "es-AR"
	LanguageEsBO         Language = "es-BO"
	LanguageEsCL         Language = "es-CL"
	LanguageEsCO         Language = "es-CO"
	LanguageEsCR         Language = "es-CR"
	LanguageEsCU         Language = "es-CU"
	LanguageEsDO         Language = "es-DO"
	LanguageEsEC         Language = "es-EC"
	LanguageEsES         Language = "es-ES"
	LanguageEsGQ         Language = "es-GQ"
	LanguageEsGT         Language = "es-GT"
	LanguageEsHN         Language = "es-HN"
	LanguageEsMX         Language = "es-MX"
	LanguageEsNI         Language = "es-NI"
	LanguageEsPA         Language = "es-PA"
	LanguageEsPE         Language = "es-PE"
	LanguageEsPR         Language = "es-PR"
	LanguageEsPY         Language = "es-PY"
	LanguageEsSV         Language = "es-SV"
	LanguageEsUS         Language = "es-US"
	LanguageEsUY         Language = "es-UY"
	LanguageEsVE         Language = "es-VE"
	LanguageEtEE         Language = "et-EE"
	LanguageEuES         Language = "eu-ES"
	LanguageFaIR         Language = "fa-IR"
	LanguageFiFI         Language = "fi-FI"
	LanguageFilPH        Language = "fil-PH"
	LanguageFrBE         Language = "fr-BE"
	LanguageFrCA         Language = "fr-CA"
	LanguageFrCH         Language = "fr-CH"
	LanguageFrFR         Language = "fr-FR"
	LanguageGaIE         Language = "ga-IE"
	LanguageGlES         Language = "gl-ES"
	LanguageGuIN         Language = "gu-IN"
	LanguageHeIL         Language = "he-IL"
	LanguageHiIN         Language = "hi-IN"
	LanguageHrHR         Language = "hr-HR"
	LanguageHuHU         Language = "hu-HU"
	LanguageHyAM         Language = "hy-AM"
	LanguageIdID         Language = "id-ID"
	LanguageIsIS         Language = "is-IS"
	LanguageItIT         Language = "it-IT"
	LanguageJaJP         Language = "ja-JP"
	LanguageJvID         Language = "jv-ID"
	LanguageKaGE         Language = "ka-GE"
	LanguageKkKZ         Language = "kk-KZ"
	LanguageKmKH         Language = "km-KH"
	LanguageKnIN         Language = "kn-IN"
	LanguageKoKR         Language = "ko-KR"
	LanguageLoLA         Language = "lo-LA"
	LanguageLtLT         Language = "lt-LT"
	LanguageLvLV         Language = "lv-LV"
	LanguageMkMK         Language = "mk-MK"
	LanguageMlIN         Language = "ml-IN"
	LanguageMnMN         Language = "mn-MN"
	LanguageMrIN         Language = "mr-IN"
	LanguageMsMY         Language = "ms-MY"
	LanguageMtMT         Language = "mt-MT"
	LanguageMyMM         Language = "my-MM"
	LanguageNbNO         Language = "nb-NO"
	LanguageNeNP         Language = "ne-NP"
	LanguageNlBE         Language = "nl-BE"
	LanguageNlNL         Language = "nl-NL"
	LanguagePlPL         Language = "pl-PL"
	LanguagePsAF         Language = "ps-AF"
	LanguagePtBR         Language = "pt-BR"
	LanguagePtPT         Language = "pt-PT"
	LanguageRoRO         Language = "ro-RO"
	LanguageRuRU         Language = "ru-RU"
	LanguageSiLK         Language = "si-LK"
	LanguageSkSK         Language = "sk-SK"
	LanguageSlSI         Language = "sl-SI"
	LanguageSoSO         Language = "so-SO"
	LanguageSqAL         Language = "sq-AL"
	LanguageSrLATNRS     Language = "sr-LATN-RS"
	LanguageSrRS         Language = "sr-RS"
	LanguageSuID         Language = "su-ID"
	LanguageSvSE         Language = "sv-SE"
	LanguageSwKE         Language = "sw-KE"
	LanguageSwTZ         Language = "sw-TZ"
	LanguageTaIN         Language = "ta-IN"
	LanguageTaLK         Language = "ta-LK"
	LanguageTaMY         Language = "ta-MY"
	LanguageTaSG         Language = "ta-SG"
	LanguageTeIN         Language = "te-IN"
	LanguageThTH         Language = "th-TH"
	LanguageTrTR         Language = "tr-TR"
	LanguageUkUA         Language = "uk-UA"
	LanguageUrIN         Language = "ur-IN"
	LanguageUrPK         Language = "ur-PK"
	LanguageUzUZ         Language = "uz-UZ"
	LanguageViVN         Language = "vi-VN"
	LanguageWuuCN        Language = "wuu-CN"
	LanguageYueCN        Language = "yue-CN"
	LanguageZhCN         Language = "zh-CN"
	LanguageZhCNGUANGXI  Language = "zh-CN-GUANGXI"
	LanguageZhCNhenan    Language = "zh-CN-henan"
	LanguageZhCNliaoning Language = "zh-CN-liaoning"
	LanguageZhCNshaanxi  Language = "zh-CN-shaanxi"
	LanguageZhCNshandong Language = "zh-CN-shandong"
	LanguageZhCNsichuan  Language = "zh-CN-sichuan"
	LanguageZhHK         Language = "zh-HK"
	LanguageZhTW         Language = "zh-TW"
	LanguageZuZA         Language = "zu-ZA"
)

func (l Language) String() string {
	return string(l)
}
