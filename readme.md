This stream editor transforms Hebrew from the 1982 Michigan-Claremont encoding to Unicode.

Use it on the command line as you would any other stream editor, e.g.:
`cat wlc420_ps.txt | mctounicode > wlc420_unicode.txt`

Features (these may be made optional in the future):
- Ignores comment lines beginning with #
- Ignores first word references (gn1:1), as found in the Tanakh.

MC Encoding also specifies Greek, this may be added in the future.

More on the MC encoding:
- http://ccat.sas.upenn.edu/gopher/text/religion/biblical/mbhs/002.MichiganManual
- http://ccat.sas.upenn.edu/beta/key.html
