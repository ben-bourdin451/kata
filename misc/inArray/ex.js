function inArray(a1, a2) {
	return a1.filter(e => a2.contains(e));
}

module.exports = inArray;
