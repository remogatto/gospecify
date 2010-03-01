/*
Copyright (c) 2009 Samuel Tesla <samuel.tesla@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package specify

import (
	"fmt"
)

type specdocReporter struct {
	*basicReporter
}

func makeSpecdocReporter() (result *specdocReporter) { return &specdocReporter{NewBasicReporter()} }

func (self *specdocReporter) Error(r Report) {
	self.basicReporter.Error(r)
	fmt.Printf("- %s (ERROR)\n", r.Title())
}

func (self *specdocReporter) Fail(r Report) {
	self.basicReporter.Fail(r)
	fmt.Printf("- %s (FAILED)\n", r.Title())
}

func (self *specdocReporter) printList(label string, reports <-chan Report) {
	fmt.Printf("\n%v:\n", label)
	for r := range reports {
		fmt.Printf("\n- %v - %v\n  %v\n", r.Title(), r.Error(), r.Location())
	}
}

func (self *specdocReporter) Finish() {
	fmt.Printf("\nPassing: %v  Failing: %v  Pending: %v  Errors: %v\n", self.PassingCount(), self.FailingCount(), self.PendingCount(), self.ErrorCount())
	if self.ErrorCount() > 0 {
		self.printList("Errors", self.EachError())
	}
	if self.FailingCount() > 0 {
		self.printList("Failing Examples", self.EachFailure())
	}
	if self.PendingCount() > 0 {
		self.printList("Pending Examples", self.EachPending())
	}
}

func (self *specdocReporter) Pass(r Report) {
	self.basicReporter.Pass(r)
	fmt.Printf("- %s\n", r.Title())
}

func (self *specdocReporter) Pending(r Report) {
	self.basicReporter.Pending(r)
	fmt.Printf("- %s (Not yet implemented)\n", r.Title())
}
