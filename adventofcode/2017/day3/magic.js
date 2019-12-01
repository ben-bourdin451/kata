function mumboJumbo(target) {
	if (target == 1) {
		return 0;
	}
	
	let numSqrt = parseInt(Math.sqrt(target));
	let root = 0;
	if ((numSqrt % 2) == 0) {
		root = numSqrt + 1;
	} else {
		if (Math.pow(numSqrt, 2) == target) {
			root = numSqrt;
		} else {
			root = numSqrt + 2;
	 	}
	}
	 
	let squaredRoot = Math.pow(root, 2);
	let halfRoot = Math.trunc(root/2);
	let difference = squaredRoot - target;
	if (difference % halfRoot == 0) {
		if ((difference / halfRoot) % 2 != 0) {
			return halfRoot;
		} else {
			return halfRoot * 2;
		}
	} else {
		if (difference < root) {
			return (Math.abs((halfRoot % difference) - halfRoot) + halfRoot);
		} else {
			let divi = parseInt(difference / root);
			let real_diff = (difference  - (divi * root) + divi) % halfRoot;
			return (Math.abs(real_diff - halfRoot) + halfRoot);
		}
	}
	
	return 0;
}

module.exports = { mumboJumbo };
