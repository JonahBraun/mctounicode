package mctounicode

import (
	"bufio"
	"io"
	"strings"

	H "github.com/JonahBraun/mctounicode/hebrew"
)

var mcHebrew = map[string]string{
	// characters
	")": H.Alef,
	"B": H.Bet,
	"G": H.Gimel,
	"D": H.Dalet,
	"H": H.He,
	"W": H.Vav,
	"Z": H.Zayin,
	"X": H.Het,
	"+": H.Tet,
	"Y": H.Yod,
	"K": H.Kaf,
	"L": H.Lamed,
	"M": H.Mem,
	"N": H.Nun,
	"S": H.Samekh,
	"(": H.Ayin,
	"P": H.Pe,
	"C": H.Tsadi,
	"Q": H.Qof,
	"R": H.Resh,
	"#": H.Shin,
	"&": H.Shin + H.SinDot,
	"$": H.Shin + H.ShinDot,
	"T": H.Tav,
	// points
	"A":  H.Patah,
	"F":  H.Qamats,
	"I":  H.Hiriq,
	"E":  H.Segol,
	"\"": H.Tsere,
	"O":  H.Holam,
	"U":  H.Qubuts,
	"W.": H.Dagesh, // =shuruq
	":":  H.Sheva,
	"OW": H.Vav + H.Holam,
	":A": H.HatafPatah,
	":F": H.HatafQamats,
	":E": H.HatafSegol,
	"-":  H.Maqaf,
	".":  H.Dagesh,
	",":  H.Rafe,
	"*":  "", // ketiv
	"**": "", // qere
	// Cantillation
	// at end of word above
	"00": H.SofPasuq,
	"01": H.Segolta,
	"02": H.Zarqa,
	"03": H.Pashta,
	"04": H.TelishaQetana, // =telisha parvum (parvum latin for small)
	"05": H.Paseq,
	// at start of word below
	"10": H.Yetiv,
	"13": H.Dehi,
	// at start of word above
	"11": H.GereshMuqdam,  // =mugrash ??
	"14": H.TelishaGedola, // =telisha magnum
	// above word
	"24": H.TelishaQetana,
	"33": H.Pashta,
	"44": H.TelishaGedola, // =telisha magnum
	"60": H.Ole,
	"61": H.Geresh,
	"62": H.Gershayim,
	"63": H.Qadma, // also azla
	"64": H.Iluy,
	"65": H.Shalshelet,
	"80": H.ZaqefQatan,
	"81": H.Revia,
	"82": H.Zarqa, // =zinorit
	"83": H.Pazer,
	"84": H.QarneyPara, // =pazer gadol
	"85": H.ZaqefGadol,
	// below word
	"35": H.Meteg,
	"70": H.Mahapakh,
	"71": H.Merkha,
	"72": H.MerkhaKefula,
	"73": H.Tipeha,
	"74": H.Munah,
	"75": H.Meteg, // =silluq
	"91": H.Tevir,
	"92": H.Etnahta,      // =atnah
	"93": H.YerahBenYomo, // =galgal
	"94": H.Darga,
	"95": H.Meteg,
}

type MCScanner struct {
	*bufio.Scanner
	text          string
	skipComments  bool
	skipReference bool
}

func NewMCScanner(r io.Reader) *MCScanner {
	return &MCScanner{
		bufio.NewScanner(r),
		"",
		true,
		true,
	}
}

// BUG(JonahBraun): Source text must be ASCII, single bytes are treated as characters.
func (s *MCScanner) Scan() bool {
	if !s.Scanner.Scan() {
		return false
	}

	s.text = ""
	src := s.Scanner.Text()

	// Ignore comment lines
	if s.skipComments && strings.HasPrefix(src, "#") {
		s.text = src
		return true
	}

	pos := 0
	if s.skipReference {
		pos = strings.Index(src, " ")
		if pos > -1 {
			s.text = src[:pos]
		}
	}

	// This block is pretty ugly
	// Bail at len-2 because we first perform a read ahead and will panic at end of string
	for ; pos < len(src)-1; pos++ {
		two := string(string(src[pos : pos+2]))
		one := string(two[0])

		if u, ok := mcHebrew[two]; ok {
			s.text += u
			// two letter replacement, advance twice
			pos++
		} else if u, ok := mcHebrew[one]; ok {
			s.text += u
		} else {
			s.text += one
		}

	}
	// handle the last byte if necessary
	if pos < len(src) {
		if u, ok := mcHebrew[string(src[pos])]; ok {
			s.text += u
		} else {
			s.text += string(src[pos])
		}
	}

	return true
}

func (s *MCScanner) Text() string {
	return s.text
}
