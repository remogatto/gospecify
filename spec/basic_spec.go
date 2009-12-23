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
package main

import (
	. "specify";
	t "../src/testspecify";
)

func init() {
	Describe("Running", func() {
		Before(func(the Example) {
			reporter := testRun(func(r t.Runner) {
				r.It("pass", func(the t.Example) {
					the.Value(7 * 6).Should(t.Be(42));
					the.Value(1).ShouldNot(t.Be(2));
				});

				r.It("fail", func(the t.Example) {
					the.Value(7 * 6).ShouldNot(t.Be(42));
					the.Value(1).Should(t.Be(2));
				});

				r.It("pending", nil);
			});
			the.SetField("reporter", reporter);
		});

		It("counts passing examples", func(the Example) { the.Field("reporter").Should(HavePassing(1)) });
		It("counts failing examples", func(the Example) { the.Field("reporter").Should(HaveFailing(1)) });
		It("counts pending examples", func(the Example) { the.Field("reporter").Should(HavePending(1)) });

		It("includes the example name in the failure message", func(the Example) { the.Field("reporter").Should(HaveFailureIncluding("fail")) });
		It("reports the example name when pending", nil);	// dotReporter should print them in Finish
	})
}
