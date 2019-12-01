function captcha(input) {
	return input.reduce((accumulator, currentValue, currentIndex, arr) => {
		let nextIndex = (currentIndex >= arr.length - 1) ? 0 : currentIndex + 1;
		if (currentValue == arr[nextIndex]) {
			return accumulator + currentValue;
		} else {
			return accumulator;
		}
	}, 0);
}

function captcha2(input) {
	return input.reduce((accumulator, currentValue, currentIndex, arr) => {
		
		let nextIndex = ((arr.length / 2) + currentIndex) % arr.length;
		
		if (currentValue == arr[nextIndex]) {
			return accumulator + currentValue;
		} else {
			return accumulator;
		}
	}, 0);
}

module.exports = { captcha, captcha2 };
