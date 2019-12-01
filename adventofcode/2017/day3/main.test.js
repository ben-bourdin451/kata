const magicMumboJumbo = require('./magic.js').mumboJumbo;
const {crawler, superCrawler} = require('./crawler.js');

describe('Part 1 tests', () => {
	test('1 requires 0 steps', () => {
		expect(crawler(1)).toEqual(0);
		expect(magicMumboJumbo(1)).toEqual(0);
	});
	test('12 requires 3 steps', () => {
		expect(crawler(12)).toEqual(3);
		expect(magicMumboJumbo(12)).toEqual(3);
	});
	test('23 requires 2 steps', () => {
		expect(crawler(23)).toEqual(2);
		expect(magicMumboJumbo(23)).toEqual(2);
	});
	test('1024 requires 31 steps', () => {
		expect(crawler(1024)).toEqual(31);
		expect(magicMumboJumbo(1024)).toEqual(31);
	});
	test('final answer', () => {
		expect(crawler(289326)).toEqual(419);
		expect(magicMumboJumbo(289326)).toEqual(419);
	});
});

describe('Part 2 tests', () => {
	test('> 10', () => {
		expect(superCrawler(10)).toEqual(11);
	});
	test('> 20', () => {
		expect(superCrawler(20)).toEqual(23);
	});
	test('> 50', () => {
		expect(superCrawler(50)).toEqual(54);
	});
	test('> 100', () => {
		expect(superCrawler(100)).toEqual(122);
	});
	test('> 300', () => {
		expect(superCrawler(300)).toEqual(304);
	});
	test('final answer', () => {
		console.log(superCrawler(289326));
	});
});
