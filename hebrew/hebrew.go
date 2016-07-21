// String constants for the Hebriew Unicode set
//
// http://unicode.org/charts/PDF/U0590.pdf
package hebrew

// Cantillation marks 0x0591
const (
	Etnahta = string(0x0591 + iota)
	Segolta
	Shalshelet
	ZaqefQatan
	ZaqefGadol
	Tipeha
	Revia
	Zarqa
	Pashta
	Yetiv
	Tevir
	Geresh
	GereshMuqdam
	Gershayim
	QarneyPara
	TelishaGedola
	Pazer
	AtnahHafukh
	Munah
	Mahapakh
	Merkha
	MerkhaKefula
	Darga
	Qadma
	TelishaQetana
	YerahBenYomo
	Ole
	Iluy
	Dehi
	Zinor
	MasoraCircle
	// Points and punctuation 0x5b0
	Sheva
	HatafSegol
	HatafPatah
	HatafQamats
	Hiriq
	Tsere
	Segol
	Patah
	Qamats
	Holam
	HolamHaser
	Qubuts
	Dagesh // dagesh, mapiq, shuruq
	Meteg
	Maqaf
	Rafe
	Paseq
	ShinDot
	SinDot
	SofPasuq
	// Puncta extraordinaria 0x05c4
	UpperDot
	LowerDot
	// Points and punctuation 0x05c6
	NunHafukha
	QamatsQatan
)

// Alphabet 0x05d0
const (
	Alef = string(0x05d0 + iota)
	Bet
	Gimel
	Dalet
	He
	Vav
	Zayin
	Het
	Tet
	Yod
	FinalKaf
	Kaf
	Lamed
	FinalMem
	Mem
	FinalNun
	Nun
	Samekh
	Ayin
	FinalPe
	Pe
	FinalTsadi
	Tsadi
	Qof
	Resh
	Shin
	Tav
)

// Yiddish digraphs 0x05F0
const (
	YiddishDoubleVav = string(0x05f0 + iota)
	YiddishVavYod
	YiddishDoubleYod
	// Additional punctuation 0x05F3
	YiddishGeresh
	YiddishGershayim
)
