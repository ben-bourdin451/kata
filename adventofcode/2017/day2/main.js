function checksum(input) {
	return input.split('\n').reduce((acc, val) => {
		// split the string, parse to int and sort asc
		let sortedArr = val.split('\t').map(n => parseInt(n)).sort((a, b) => a - b);
		
		// return the difference: last (max) - first (min)
		return acc + sortedArr[sortedArr.length - 1] - sortedArr[0];
	}, 0);
}

// part 2 functional solution
function checksum2(input) {
	return input.split('\n').reduce((acc, val) => {
		// parse strings array to int array
		// figure out the divisible number by poping each element and figuring out if it
		// has a corresponding divisible number 
		let divisibleResult = val.split('\t').map(n => parseInt(n)).reduce((acc2, val2, index2, arr2) => {
			// copy arr and remove value at current index
			let arrCp = arr2.slice();
			arrCp.splice(index2, 1);
			
			// if any vals in our reduced arr has a 0 modulo (can be divided)
			// sum the result of that division
			// else return accumulater (add 0)
			let divider = arrCp.filter(x => val2 % x == 0);
			if (divider.length > 0) {
				return acc2 + (val2 / divider[0]);
			} else {
				return acc2;
			};
		}, 0);
		
		// sum the result of the division for each line
		return acc + divisibleResult;
	}, 0);
}

// part 2 solution with for loops
function checksum3(input) {
	return input.split('\n').reduce((acc, val) => {
		let parsedArr = val.split('\t').map(n => parseInt(n));
		for (let index1 = 0; index1 < parsedArr.length - 1; index1++) {
			for (let index2 = index1 + 1; index2 < parsedArr.length; index2++) {
				if (parsedArr[index1] % parsedArr[index2] == 0) {
					return acc + (parsedArr[index1] / parsedArr[index2])
				} else if (parsedArr[index2] % parsedArr[index1] == 0) {
					return acc + (parsedArr[index2] / parsedArr[index1])
				}
			}
		}
	}, 0);	
}

module.exports = { checksum, checksum2, checksum3 };
