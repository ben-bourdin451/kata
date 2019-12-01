const captcha = require('./main.js').captcha;
const captcha2 = require('./main.js').captcha2;

describe('Part 1 tests', () => {
	it('2 consecutive pairs', () => {
		expect(captcha([1, 1, 2, 2])).toEqual(3);
	});
	it('all the same', () => {
		expect(captcha([1, 1, 1, 1])).toEqual(4);
	});
	it('increment returns 0', () => {
		expect(captcha([1, 2, 3, 4])).toEqual(0);
	});
	it('long matching at the end', () => {
		expect(captcha([9, 1, 2, 1, 2, 1, 2, 9])).toEqual(9);
	});
});

describe('Part 2 tests', () => {
	it('2 consecutive pairs', () => {
		expect(captcha2([1, 2, 1, 2])).toEqual(6);
	});
	it('all the same', () => {
		expect(captcha2([1, 2, 2, 1])).toEqual(0);
	});
	it('increment returns 0', () => {
		expect(captcha2([1, 2, 3, 4, 2, 5])).toEqual(4);
	});
	it('increment returns 0', () => {
		expect(captcha2([1, 2, 3, 1, 2, 3])).toEqual(12);
	});
	it('long matching at the end', () => {
		expect(captcha2([1, 2, 1, 3, 1, 4, 1, 5])).toEqual(4);
	});
});
