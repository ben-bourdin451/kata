const main = require('./main.js');
const fs = require('fs');
const readline = require('readline');

describe('Part 1 tests', () => {
	
	test('example takes 5 steps', () => {
		expect(main([0, 3, 0, 1, -3])).toEqual(5);
	});
	test('example takes 5 steps', () => {
		expect(main([0, 3, 0, 1, -3])).toEqual(5);
	});
	test('part 1 answer', (done) => {
		const rl = readline.createInterface({
			input: fs.createReadStream(process.cwd() + '/day5/input.txt'),
			crlfDelay: Infinity
		});

		input = [];
		rl.on('line', (line) => {
			if (main(line)) {
				input.push(line)
			}
		});

		rl.on('close', () => {
			console.log(main(input));
			done();
		});
	})
});
