import { redirect } from "@sveltejs/kit";

export const load = async (event) => {
    event.cookies.delete('sessionid', { path: '/' });
    event.locals.user = null;
    throw redirect(301,'/login');
};


