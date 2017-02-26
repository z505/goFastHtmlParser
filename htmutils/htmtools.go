/* Fast HTML parser tools converted to Go. Originally by James Azarja and 
   Lars Olson for fpc. This works but is not finished yet, as just a simple
   conversion from freepascal was done. Not correctly go formatted, not 
   aesthetically in best shape. Needs polishing and cleaning up. */

package htmutils

import (
	"fmt"
	"strings"
	"strconv"
)

// checks whether a character (b byte) matches other characters given as arguments
func isChar(b byte, chars ...string) bool {
    for _, c := range chars {
        if string(b) == c { return true }
    }
    return false
}

// Return name=value pair ignore case of NAME, preserve case of VALUE
func GetNameValPair(tag, name string) string {
	var s, UpperName string
	var i, idx int
	var C string // char

	if tag == "" || name == "" {
		return ""
	}
	// must be space before case insensitive NAME, i.e. <a HREF= STYLE=
	UpperName = " " + strings.ToUpper(name)
	s = strings.ToUpper(tag)

	idx = strings.Index(s, UpperName)
	// no name value pair found
	if idx == -1 {
		return "" // todo: could return error since name was not found in tag
	}

	idx++ // skip space
	i = idx

	// Skip until hopefully equal sign
	for !isChar(s[i], "=", " ", ">") {
		if i == len(s)-1 { break }
		i++
	}

	if string(s[i]) == "=" { i++ }

	for !isChar(s[i], " ", ">") {
		if i == len(s)-1 { break }
		if isChar(s[i], "\"", "'") {
			C = string(s[i])
			i++ // Skip quote
		} else {
			C = " "
		}

		for !isChar(s[i], C, ">") {
			if i == len(s)-1 { break }
			i++
		}

		if string(s[i]) != ">" { i++ } // Skip current character, except '>'
		break
	}

    // extract the string slice where the name value pair is, return it
    return tag[idx:i]
}

// Get value of attribute, e.g WIDTH=36 returns 36, preserves case of value
func GetValFromNameVal(namevalpair string) string {
	var i int
	var idx int
	var C string // char
	var result string
	var nv string  // shortform for namevaluepair

	result = "";
	if namevalpair == "" { return result }
	nv = namevalpair
	idx = strings.Index(nv, "=")

	if idx == -1 {
		return "" // note: could return error
	}

	idx++ // skip equal
	i = idx  // set to a character after =

	if isChar(nv[i], `"`, `'`) {
		C = string(nv[i])
		i++ // Skip current character
	} else {
		C = " "
	}

	idx = i
	for string(nv[i])!=C {
		i++
		if i == len(nv) { break }
	}

	if i != idx {
		result = nv[idx:i]
	} else {
		result = ""
	}

	return result
}


