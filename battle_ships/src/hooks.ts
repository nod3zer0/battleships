/** @type {import('@sveltejs/kit').Handle} */
export async function handle({ event, resolve }) {

    if (event.cookies.get('sessionid') === undefined) {
        const response = await resolve(event);
        return response;
    }

    const user = await fetch('http://127.0.0.1:8081/user', {
        method: 'POST',
        body: JSON.stringify({ sessionid: event.cookies.get('sessionid') }),
        headers: {
            'Content-Type': 'application/json'
        }
    });



    event.locals.user = await user.json();

	const response = await resolve(event);
	return response;
}


