module.exports = (array) => {
	let index = 0;
	let steps = 1;
	while (true) {
		let instruction = array[index];
		if (instruction + index < 0 // is that possible?
			|| instruction + index > array.length - 1) {
				break;
		} else {
			array[index]++;
			steps++;
			index += instruction;
		}
	}
	
	return steps;
};
