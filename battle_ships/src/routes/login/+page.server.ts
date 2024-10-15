import { redirect } from "@sveltejs/kit";

export const load = async (event) => {

    const sessionid = event.cookies.get('sessionid');

    if (sessionid) {
        throw redirect(301,'/');
    }
};


export const actions = {
    default: async (event) => {
        const formData = await event.request.formData();
        const username = formData.get('username');
        const password = formData.get('password');



        const res = await fetch('http://127.0.0.1:8081/login', {
            method: 'POST',
            body: JSON.stringify({ username, password }),
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (res.ok) {
            const sessionid = res.headers.get('Authorization');
            event.cookies.set('sessionid', sessionid, { path: '/' });2
            return { status: 302 };
        }
    }
};