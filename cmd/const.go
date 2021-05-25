package cmd

import (
	"regexp"
)

var regexRules = []regexRule {
	regexRule {name: "email", displayName: "Email", regex: regexp.MustCompile("([a-zA-Z0-9_\\-\\.]+)@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.)|(([a-zA-Z0-9\\-]+\\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\\]?)")},
	regexRule {name: "password0", displayName: "Password", regex: regexp.MustCompile("(?i)(['\"]?\\w*password['\"]? ?[:=] ?(['\"][^'\"]{4,}['\"]))")},
	regexRule {name: "password1", displayName: "Password", regex: regexp.MustCompile("(?i)(['\"]?pass['\"]? ?[:=] ?(['\"][^'\"]{4,}['\"]|[0-9a-z\\-_@#!%\\^\\?\\*&\\$~]{4,}))")},
	regexRule {name: "password2", displayName: "Password", regex: regexp.MustCompile("(?i)(<[^(><.)]*password[^(><.)]*>[^(><.)]+<\\/[^(><.)]*password[^(><.)]*>)")},
	regexRule {name: "password3", displayName: "Password", regex: regexp.MustCompile("(?i)(['\"]?\\w*password['\"]? ?[:=] ?[0-9a-z\\-_@#! %\\^\\?\\*&\\$~]{4,})")},
	regexRule {name: "password4", displayName: "Password", regex: regexp.MustCompile("(?i)(['\"]?passwd['\"]? ?[:=] ?(['\"][^'\"]{4,}['\"]|[0-9a-z\\-_@#!%\\^\\?\\*&\\$~]{4,}))")},
	regexRule {name: "consumerKey", displayName: "Consumer Key", regex: regexp.MustCompile("(?i)(<ConsumerKey>[^<\\n]*<\\/ConsumerKey>)")},
	regexRule {name: "consumerSecret", displayName: "Consumer Secret", regex: regexp.MustCompile("(?i)(<ConsumerSecret>[^<\\n]*<\\/ConsumerSecret>)")},
	regexRule {name: "privateKey", displayName: "Private Key", regex: regexp.MustCompile("(BEGIN (RSA|OPENSSH|DSA|EC|PGP) PRIVATE KEY)")},
	regexRule {name: "cloudKey", displayName: "Cloud Key", regex: regexp.MustCompile("(AKIA[0-9A-Z]{16})")},
	regexRule {name: "AWSSecretAccessKey", displayName: "AWS Secret Access Key", regex: regexp.MustCompile("(?i)(AWS Secret Access Key [^:]*: ?([^\\s]+))")},
	regexRule {name: "clientSecret", displayName: "Client Secret", regex: regexp.MustCompile("(?i)(['\"]client_secret['\"] ?[:=] ?(['\"][^'\"]{4,}['\"]|[0-9a-z\\-_@#!%\\^\\?\\*&\\$~]{4,}))")},
	regexRule {name: "token", displayName: "Token", regex: regexp.MustCompile("(?i)(['\"]?_token['\"]? ?[:=] ?['\"]?[0-9a-zA-Z\\.\\-_]{12,}['\"]?)")},
	regexRule {name: "potentialSecret0", displayName: "Potential Secret", regex: regexp.MustCompile("(?i)(ya29.[0-9a-zA-Z_\\-]{68})")},
	regexRule {name: "potentialSecret1", displayName: "Potential Secret", regex: regexp.MustCompile("(?i)(AIzaSy[0-9a-zA-Z_\\-]{33})")}, 
	regexRule {name: "deprecatedCrypto", displayName: "Deprecated Crypto algorithm", regex: regexp.MustCompile("(?i)(DESede|3DES|TRIPLEDES)[-_a-zA-Z]*\\(.*\\)")},
	regexRule {name: "insecureDbConn", displayName: "Insecure Database Connection", regex: regexp.MustCompile("(?i)(mysql|oracle|odbc|jdbc|postgresql|mongodb|mongo):\\/\\/\\w{3,}:\\w{3,}(@[^\\/]{3,}\\/)")},
	regexRule {name: "pgpPrivateKey", displayName: "PGP Private Key", regex: regexp.MustCompile("(BEGIN PGP PRIVATE KEY)")},
	regexRule {name: "secretKey0", displayName: "Secret Key", regex: regexp.MustCompile("(?i)(['\"]?secret_key['\"]? ?[:=] ?(['\"][^'\"]{4,}['\"]|[0-9a-z\\-_@#!%\\^\\?\\*&\\$~]{4,}))")},
	regexRule {name: "secretKey1", displayName: "Secret Key", regex: regexp.MustCompile("(?i)(['\"]?secretkey['\"]? ?[:=] ?(['\"][^'\"]{4,}['\"]|[0-9a-z\\-_@#!%\\^\\?\\*&\\$~]{4,}))")},
	regexRule {name: "secretKey2", displayName: "", regex: regexp.MustCompile("(?i)(['\"]?secret['\"]? ?[:=] ?(['\"][^'\"]{3,}['\"]|[0-9a-z\\-_@#!%\\^\\?\\*&\\$~]+))")},
	regexRule {name: "passphrase", displayName: "Passphrase", regex: regexp.MustCompile("(?i)(<[^(><.)]*passphrase[^(><.)]*>[^(><.)]+<\\/[^(><.)]*passphrase[^(><.)]*>)")},
	regexRule {name: "mobilenumber0", displayName: "Mobile Number", regex: regexp.MustCompile("(\\+91|\\+91\\-|0)?\\d{10}")},
	regexRule {name: "cardnumber", displayName: "Potential Credit/Debit Card", regex: regexp.MustCompile("(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|6(?:011|5[0-9]{2})[0-9]{12}(?:2131|1800|35\\d{3})\\d{11})")},
	regexRule {name: "zipcode", displayName: "Potential Zip Code", regex: regexp.MustCompile("\\b\\d{6}(?:[-\\s]\\d{4})?\\b")},
}