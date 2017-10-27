const parseHTMLColor = require('./ex.js');

describe('Example tests', () => {
	it('should parse example colors', () => {
		expect(parseHTMLColor('#80FFA0')).toEqual({ r: 128, g: 255, b: 160 });
		expect(parseHTMLColor('#3B7')).toEqual({ r: 51, g: 187, b: 119 });
		expect(parseHTMLColor('LimeGreen')).toEqual({ r: 50, g: 205, b: 50 });
	});
});
