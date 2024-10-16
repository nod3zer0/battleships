export async function load({ locals }) {

let ships = [ { "name": "Carrier","placed": "false", "position": [0,0], "parts": [ [0,0], [0,1], [0,2], [0,3], [0,4] ] }, { "name": "Battleship","placed": "false", "position": [0,0], "parts": [ [0,0], [0,1], [0,2], [0,3] ] }, { "name": "Cruiser","placed": "false", "position": [0,0], "parts": [ [0,0], [0,1], [0,2] ] }, { "name": "Submarine","placed": "false", "position": [0,0], "parts": [ [0,0], [0,1], [0,2] ] }, { "name": "Destroyer","placed": "false", "position": [0,0], "parts": [ [0,0], [0,1] ] } ];

ships.forEach(ship => {
    ship.bitmap = createBitmap(ship, 5);
});


return {
    ships: ships
}
	// return {
	// 	user: locals.user,
	// }
}

// Function to create a grid for each ship
function createBitmap(ship, gridSize) {
    // Initialize the grid with false (empty)
    let grid = Array.from({ length: gridSize }, () => Array(gridSize).fill(false));

    // Mark ship parts as true
    ship.parts.forEach(([x, y]) => {
        grid[x][y] = true;
    });

    return grid;
}