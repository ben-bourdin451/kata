const checksum = require('./main.js').checksum;
const checksum2 = require('./main.js').checksum2;
const checksum3 = require('./main.js').checksum3;

describe('Part 1 tests', () => {
	test('spreadsheet with 3 lines', () => {
		expect(checksum("5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8")).toEqual(18);
	});
});
describe('Part 2 tests', () => {
	test('functional solution', () => {
		expect(checksum2("5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5")).toEqual(9);
	});
	test('easy solution', () => {
		expect(checksum3("5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5")).toEqual(9);
	});
});
