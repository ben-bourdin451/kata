const inArray = require('./ex.js');

describe('Example tests', () => {
	a2 = ["lively", "alive", "harp", "sharp", "armstrong"]

	it('should parse example colors', () => {
		a1 = ["xyz", "live", "strong"]
		expect(inArray(a1, a2)).toEqual(["live", "strong"]);
		a1 = ["live", "strong", "arp"]
		expect(inArray(a1, a2)).toEqual(["arp", "live", "strong"]);
		a1 = ["tarp", "mice", "bull"]
		expect(inArray(a1, a2)).toEqual([]);
	});
});
