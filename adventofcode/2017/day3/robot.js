module.exports = class Robot {
	constructor() {
		this.x = 0;
		this.y = 0;
		this.map = { 0: { 0: 1 } };
		this.currentValue = 1;
	}
	
	move(direction) {
		switch (direction) {
			case 'right':
				this.x++;
				break;
			case 'up':
				this.y++;
				break;
			case 'left':
				this.x--;
				break;
			case 'down':
				this.y--;
				break;
		}
		
		this.currentValue = this.getNeighborsSum();
		if (this.map[this.x] === undefined) {
			this.map[this.x] = {};
		}
		this.map[this.x][this.y] = this.currentValue;
		
		return this.currentValue;
	}
	
	getNeighborsSum() {
		let sum = 0;
		
		// x, y-1
		// x, y+1
		if (this.map[this.x] !== undefined) {
			sum += this.map[this.x][this.y - 1] || 0;
			sum += this.map[this.x][this.y + 1] || 0;
		}
		
		// x+1, y
		// x+1, y-1
		// x+1, y+1
		if (this.map[this.x + 1] !== undefined) {
			sum += this.map[this.x + 1][this.y] || 0;
			sum += this.map[this.x + 1][this.y - 1] || 0;
			sum += this.map[this.x + 1][this.y + 1] || 0;
		}
		
		// x-1, y
		// x-1, y-1
		// x-1, y+1
		if (this.map[this.x - 1] !== undefined) {
			sum += this.map[this.x - 1][this.y] || 0;
			sum += this.map[this.x - 1][this.y - 1] || 0;
			sum += this.map[this.x - 1][this.y + 1] || 0;
		}
		
		return sum;
	}
	
	getDistance() {
		return Math.abs(this.x) + Math.abs(this.y);
	}
	
	getCurrentValue() {
		return this.currentValue;
	}
}
