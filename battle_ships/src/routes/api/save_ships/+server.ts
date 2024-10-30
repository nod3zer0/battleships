import { json } from '@sveltejs/kit';

/** @type {import('./$types').RequestHandler} */
export async function POST({ request }) {
	const data  = await request.json();


    fetch("http://127.0.0.1:8081/save_ships", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            ships: data.ships,
            id: data.id,
            user_session: data.session,
        }),
    });

    console.log(a);
    return json({ message: 'ok' });
}