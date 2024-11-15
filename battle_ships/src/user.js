 import GUN from 'gun';
 import 'gun/sea';
 import 'gun/axe';
 import { writable } from 'svelte/store';


 export const db = GUN();

 export const user = db.user().recall({ sessionStorage: true });


 user.get('alias')

 export const username = writable('');

 user.get('alias').on(v => username.set(v));

 db.on('auth', async(at) => {
     const alias = await user.get('alias');
     username.set(alias);

     console.log(`Signed in as ${alias}`);
 });