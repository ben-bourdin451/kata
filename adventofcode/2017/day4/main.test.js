const main = require('./main.js');
const fs = require('fs');
const readline = require('readline');

describe('Part 1 tests', () => {
	it('should return valid on no duplicate words', () => {
		expect(main('aa bb cc dd ee')).toBeTruthy();
	});
	
	it('should return invalid with duplicate words', () => {
		expect(main('aa bb cc dd aa')).toBeFalsy();
	});
	
	it('should return valid on words that are similar', () => {
		expect(main('aa bb cc dd aaa')).toBeTruthy();
	});
	
	test('final test', (done) => {
		const rl = readline.createInterface({
			input: fs.createReadStream(process.cwd() + '/day4/input.txt'),
			crlfDelay: Infinity
		});

		sum = 0;
		rl.on('line', (line) => {
			if (main(line)) {
				sum++;
			}
		});

		rl.on('close', () => {
			console.log(sum);
			done();
		});
	});
});
