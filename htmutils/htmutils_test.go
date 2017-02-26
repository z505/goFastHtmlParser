/* Fast HTML parser tools converted to Go. 
   This shows some sample tests */


package htmutils

import (
	"fmt"
)

type other struct {
	s string
}

func (o *other) onFoundTag(tag string) {
	fmt.Printf("Tag Found: " + tag +"\n");
	o.s = "other test"
	fmt.Printf(o.s+"\n")
//	test = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
//	fmt.Printf(test+"\n")
}

func onFoundText(s string) {
	fmt.Printf("Text Found: " + s +"\n")
}


func test() {
	var h State
	var o other
	// o = new(other)
	h.OnFoundTag = o.onFoundTag
	h.OnFoundText = onFoundText
	s := "<html><b>testing</b> and this is <i>a</i> test</html>  this is not shown because it is after closing tag"
	h.Parse(s)
//	printDone()
	s = "<html><b>testing</b> <a href=\"http://test.com\">link</a> and this is <i>a</i> test</html>  this is not shown because it is after closing tag"
	h.Parse(s)
//	printDone()
}
