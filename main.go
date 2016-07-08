package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
const(
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
const(
	YiddishDoubleVav = string(0x05f0 + iota)
	YiddishVavYod
	YiddishDoubleYod
// Additional punctuation 0x05F3
	YiddishGeresh
	YiddishGershayim
)


var mcHebrew = map[string]string{
	// characters
	")": Alef,
	"B": Bet,
	"G": Gimel,
	"D": Dalet,
	"H": He,
	"W": Vav,
	"Z": Zayin,
	"X": Het,
	"+": Tet,
	"Y": Yod,
	"K": Kaf,
	"L": Lamed,
	"M": Mem,
	"N": Nun,
	"S": Samekh,
	"(": Ayin,
	"P": Pe,
	"C": Tsadi,
	"Q": Qof,
	"R": Resh,
	"#": Shin,
	"&": Shin+SinDot,
	"$": Shin+ShinDot,
	"T": Tav,
	// points
	"A":  Patah,
	"F":  Qamats,
	"I":  Hiriq,
	"E":  Segol,
	"\"": Tsere,
	"O":  Holam,
	"U":  Qubuts,
	"W.": Dagesh, // =shuruq
	":":  Sheva,
	"OW": Vav+Holam,
	":A": HatafPatah,
	":F": HatafQamats,
	":E": HatafSegol,
	"-":  Maqaf,
	".":  Dagesh,
	",":  Rafe,
	"*": "", // ketiv
	"**": "", // qere
	// Cantillation
	// at end of word above
	"00": SofPasuq,
	"01": Segolta,
	"02": Zarqa,
	"03": Pashta,
	"04": TelishaQetana, // =telisha parvum (parvum latin for small)
	"05": Paseq,
	// at start of word below
	"10": Yetiv,
	"13": Dehi,
	// at start of word above
	"11": GereshMuqdam, // =mugrash ??
	"14": TelishaGedola, // =telisha magnum 
	// above word
	"24": TelishaQetana,
	"33": Pashta,
	"44": TelishaGedola, // =telisha magnum
	"60": Ole,
	"61": Geresh,
	"62": Gershayim,
	"63": Qadma, // also azla
	"64": Iluy,
	"65": Shalshelet,
	"80": ZaqefQatan,
	"81": Revia,
	"82": Zarqa, // =zinorit
	"83": Pazer,
	"84": QarneyPara, // =pazer gadol
	"85": ZaqefGadol,
	// below word
	"35": Meteg,
	"70": Mahapakh,
	"71": Merkha,
	"72": MerkhaKefula,
	"73": Tipeha,
	"74": Munah,
	"75": Meteg, // =silluq
	"91": Tevir,
	"92": Etnahta, // =atnah
	"93": YerahBenYomo, // =galgal
	"94": Darga,
	"95": Meteg,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "#") {
			continue
		}
		a := strings.Split(s, " ")
		s = strings.Join(a[1:], "")

		for k := range mcHebrew {
			s = strings.Replace(s, k, mcHebrew[k], -1)
		}
		fmt.Println(s)
	}
}
