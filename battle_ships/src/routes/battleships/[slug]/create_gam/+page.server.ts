import { redirect } from "@sveltejs/kit";

export const load = async (event) => {

    const sessionid = event.cookies.get('sessionid');
    console.log(sessionid);
    if (!sessionid) {
        throw redirect(301,'/login');
    }
};

export const actions = {
    default: async (event) => {
        const formData = await event.request.formData();
        const name = formData.get('gameName');
        const sessionid = event.cookies.get('sessionid');

        if (!sessionid) {
            throw redirect(301,'/login');
        }

        const res = await fetch('http://127.0.0.1:8081/create_game', {
            method: 'POST',
            body: JSON.stringify({ sessionid }),
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (res.ok) {
            redirect(303,'/battleships');
        }
    }
};