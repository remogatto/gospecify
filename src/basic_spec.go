package main

import "./specify"

var runner TestRunner;

func init() {
	initSpec();
	
	spec.Describe("Specification", func() {
		runner = &testRunner{};

		spec.Before(func () {
			s := specify.New();
			s.Describe("Foo", func() {
				s.It("pass", func() {
					s.That(7 * 6).Should().Be(42);
				});
			});
			s.Run(runner);
		});

		spec.It("indicates a passing test", func() {
			spec.That(runner.PassCount()).Should().Be(1);
		});
	});
}