module.exports = (input) => {
	let map = {};
	let isValid = true;
	input.split(' ').forEach(word => {
		if (map[word] === undefined) {
			map[word] = 1;
		} else {
			isValid = false;
		}
	});
	
	return isValid;
};
