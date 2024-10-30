import { json } from '@sveltejs/kit';


export async function load({ params, cookies }) {

//let ships = [ { "name": "Carrier", "position": [0,0], "parts": [ [0,0], [0,1], [0,2], [0,3], [0,4] ] }, { "name": "Battleship", "position": [0,0], "parts": [ [0,0], [0,1], [0,2], [0,3] ] }, { "name": "Cruiser", "position": [0,0], "parts": [ [0,0], [0,1], [0,2] ] }, { "name": "Submarine", "position": [0,0], "parts": [ [0,0], [0,1], [0,2] ] }, { "name": "Destroyer", "position": [0,0], "parts": [ [0,0], [0,1] ] } ];


const ships_json = await fetch('http://127.0.0.1:8081/ships', {
    method: 'GET',
    headers: {
        'Content-Type': 'application/json'
    }
});

let ships = await ships_json.json();

// console.log(ships_json);

// console.log(ships);

ships.forEach(ship => {
    ship.bitmap = createBitmap(ship, 5);
});

const sessionid = cookies.get('sessionid');


return {
    ships: ships,
    id: params.slug,
    session: sessionid
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

export const actions = {
    default: async (event) => {
        const formData = await event.request.formData();

        console.log(formData.get("pickedShips"));


        // const res = await fetch('http://127.0.0.1:8081/login', {
        //     method: 'POST',
        //     body: JSON.stringify({ username, password }),
        //     headers: {
        //         'Content-Type': 'application/json'
        //     }
        // });

        // if (res.ok) {
        //     const sessionid = res.headers.get('Authorization');
        //     event.cookies.set('sessionid', sessionid, { path: '/' });2
        //     return { status: 302 };
        // }
    }
};


// /** @type {import('./$types').RequestHandler} */
// export async function POST({ request }) {
// 	const a  = await request.json();

//     console.log(a);
// }