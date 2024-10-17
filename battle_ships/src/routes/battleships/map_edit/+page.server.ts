export async function load({ locals }) {

//let ships = [ { "name": "Carrier", "position": [0,0], "parts": [ [0,0], [0,1], [0,2], [0,3], [0,4] ] }, { "name": "Battleship", "position": [0,0], "parts": [ [0,0], [0,1], [0,2], [0,3] ] }, { "name": "Cruiser", "position": [0,0], "parts": [ [0,0], [0,1], [0,2] ] }, { "name": "Submarine", "position": [0,0], "parts": [ [0,0], [0,1], [0,2] ] }, { "name": "Destroyer", "position": [0,0], "parts": [ [0,0], [0,1] ] } ];


const ships_json = await fetch('http://127.0.0.1:8081/ships', {
    method: 'GET',
    headers: {
        'Content-Type': 'application/json'
    }
});

let ships = await ships_json.json();

console.log(ships_json);

console.log(ships);

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