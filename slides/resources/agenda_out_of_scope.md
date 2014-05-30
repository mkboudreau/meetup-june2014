## Out of Scope

* Test Strategy & Practices
	* (i.e. when to mock, stub, spy, dummy, or fake)
* TDD & BDD
* How to test specific Go types
	* (i.e. channels, interfaces, structs, slices)

Note:
1. why test: 
	* I don't want to assume anything about the audience
	* ... so I'd like to spend a minute or two just briefly touching on the importance of testing
2. idiomatic: 
	* maybe it is just me, but i feel like i've heard the word "idiomatic" more times in the past 2 1/2 months than I have in the past 15 years of being a developer
	* I've also seen a few misuse/abuses of the term to bolster an argument. 
	* so let's make sure we are on the same page
3. builtin facilities:
	* what does the language provide?
	* test CLI, benchmarks, coverage, detecting race conditions
4: test frameworks
	* ginkgo
	* goconvey
	* gocheck