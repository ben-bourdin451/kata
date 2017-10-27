const PRESET_COLORS = {
	limegreen: '#32cd32'
}

function parseHTMLColor(color) {
	if (color[0] === "#") {
		color = color.slice(1);
		let map = {};
		if (color.length > 4) {
			// 6 digit hex
			map.r = Number.parseInt(color.substr(0,2), 16);
			map.g = Number.parseInt(color.substr(2,2), 16);
			map.b = Number.parseInt(color.substr(4), 16);
			
		} else {
			// 3 digit hex
			map.r = Number.parseInt(color.substr(0,1) + color.substr(0,1), 16);
			map.g = Number.parseInt(color.substr(1,1) + color.substr(1,1), 16);
			map.b = Number.parseInt(color.substr(2) + color.substr(2), 16);
		}
		return map;
	} else {
	 return parseHTMLColor(PRESET_COLORS[color.toLowerCase()]);
	}
}

module.exports = parseHTMLColor;
