describe("dblLinear", () => {

	test("Binary search", function(){ 
		a = [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27];

		expect(mi(a, 12)).toBe(6);
	});

	test("Basic tests",function() {
		expect(dblLinear(10)).toBe(22);
		expect(dblLinear(20)).toBe(57);
		expect(dblLinear(30)).toBe(91);
		expect(dblLinear(50)).toBe(175);
		expect(dblLinear(100)).toBe(447);
		expect(dblLinear(500)).toBe(3355);
		expect(dblLinear(1000)).toBe(8488);
		expect(dblLinear(2000)).toBe(19773);
		expect(dblLinear(6000)).toBe(80914);
	});
});
