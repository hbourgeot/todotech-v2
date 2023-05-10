import type { Handle } from '@sveltejs/kit';
import { getAccessToken } from '$lib/auth/auth';
const authHandler: Handle = async ({ event, resolve }) => {
	console.log('http://localhost:5173/');
	console.log(
		'\x1b[35m',
		'HOOK RUN (This function runs every time the SvelteKit server receives a request — whether that happens while the app is running, or during prerendering)',
		'https://kit.svelte.dev/docs/hooks'
	);
	console.log(
		'\x1b[0;36m',
		'Requests for static assets — which includes pages that were already prerendered — are not handled by SvelteKit.'
	);

	const access_token = await getAccessToken(event);

	if (access_token) {
		event.locals.user = (await getUser(access_token)) as unknown as IUser;
	}

	return await resolve(event);
};
