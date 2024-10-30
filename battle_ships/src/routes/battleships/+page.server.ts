export async function load({ locals }) {


    const res = await fetch('http://127.0.0.1:8081/get_games', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    });

    if (res.ok) {
        const games = await res.json();
        return {
            games: games,
        }
    }


	// return {
	// 	user: locals.user,
	// }
}