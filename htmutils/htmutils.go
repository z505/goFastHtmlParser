/* Fast HTML parser tools converted to Go. Originally by James Azarja and
   Lars Olson for fpc. This works but is not finished yet, as just a simple
   conversion from freepascal was done. Not correctly go formatted, not
   aesthetically in best shape. Needs polishing and cleaning up. */


package htmutils

import "fmt"


func nilOnFoundText(s string) {
// do nothing by default if found text event is set to nil
}

func nilOnFoundTag(tag string) {
// do nothing by default if found tag event is set to nil
}

type State struct {
	OnFoundTag func(string)
	OnFoundText func(string)
}

func (h *State) Parse(htm string) {
	var TL int // text length
	var i int
	var done bool
	var tagslice, textslice string
	var TagStart, TextStart int
	var C string // char
	// if string empty then do nothing. Note: could report error here
	if htm == "" { return }
	// if no event(s) assigned, set nil procedural events that do nothing by default
	if h.OnFoundText == nil { h.OnFoundText = nilOnFoundText }
	if h.OnFoundTag == nil { h.OnFoundTag = nilOnFoundTag }

	i = 0
	TL = len(htm) // length of text string
	done = false

	TagStart = -1
	for {
		TextStart = i
		// Get next tag position
		for string(htm[i]) != "<" {
			i++
			if i >= TL {
				done = true
				break
			}
		}
		if done { break }

		// Is there any text before
		if TextStart != -1 && i > TextStart {
			textslice = htm[TextStart:i]
			// Yes, copy to buffer
			h.OnFoundText(textslice)
		} else {
			TextStart = -1
		}

		TagStart = i
		for string(htm[i]) != ">" {
			// Find quoted string in tag
			if string(htm[i]) == `"` || string(htm[i]) == "'" {
				C = string(htm[i])
				i++ // Skip current char " or '

				// Skip until quoted string end
				for string(htm[i]) != C {
					i++
					if i >= TL { break }
				}
			}

			i++
			if i >= TL {
				done = true
				break
			}
		}
		if done { break }

		// Copy this tag to buffer
		tagslice = htm[TagStart:i+1]

		h.OnFoundTag(tagslice)

		i++
		if i >= TL { break }
	}
}


