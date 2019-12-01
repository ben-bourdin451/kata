const Robot = require('./robot.js');

// crawler will map out moves until it reaches its target number
// move - number (x, y) [steps]
// r - 2 (0, 1) [1]
// u - 3 (1, 1) [2]
// l l - 4 (0, 1) [2]; 5 (-1, 1) [2]
// d d - 6 (-1, 0) [1]; 7 (-1, -1) [2] 
// r r r
// u u u
// l l l l
// d d d d
// r r r r r
function crawler(target) {
	robot = new Robot();
	
	let moves = ['right', 'up', 'left', 'down'];
	let movesIndex = 0;
	let steps = 1;
	
	// increase steps and repeat movements until we reach our target number
	for (let n = 1; n < target; steps++) {
		
		// we need to make the same movements twice before incrementing steps
		for (let i = 0; i < 2 && n < target; i++) {
			for (let stepsTaken = 0; stepsTaken < steps && n < target; stepsTaken++) {
				robot.move(moves[movesIndex]);
				n++;
			}
			movesIndex++;
			
			// wrap moves (we're going round & round)
			if (movesIndex >= moves.length) {
				movesIndex = 0;
			}
		}
	}
	
	return robot.getDistance();
}

function superCrawler(target) {
	robot = new Robot();
	
	let moves = ['right', 'up', 'left', 'down'];
	let movesIndex = 0;
	let steps = 1;
	
	// increase steps and repeat movements until we reach our target number
	for (let n = 1; n <= target; steps++) {
		
		// we need to make the same movements twice before incrementing steps
		for (let i = 0; i < 2 && n <= target; i++) {
			for (let stepsTaken = 0; stepsTaken < steps && n <= target; stepsTaken++) {
				n = robot.move(moves[movesIndex]);
			}
			movesIndex++;
			
			// wrap moves (we're going round & round)
			if (movesIndex >= moves.length) {
				movesIndex = 0;
			}
		}
	}
	
	return robot.getCurrentValue();
}

module.exports = { crawler, superCrawler };
