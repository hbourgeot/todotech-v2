import { typeChecker } from '$lib/typeChecke';
import type { RequestEvent } from '@sveltejs/kit';
export const getAccessToken = async (event: RequestEvent) => {
	const { cookies } = event;

	const token_payload: tokenPayload = typeChecker(
		'string',
		cookies.get('token_payload'),
		JSON.parse
	);
	const refresh_token = cookies.get('refresh_token');

	//If accessToken is expired (they expire every 5 mins)
	if (token_payload && token_payload.exp * 1000 < Date.now()) {
		//then get a new access toke with the refresh token
		if (refresh_token) {
			let token: string = (await refreshToken(refresh_token)) as unknown as string;
			if (token) {
				setAuthCookies({ cookies, access_token: token });
			}
		}
	}

	const access_token = cookies.get('access_token');

	return access_token;
};

export const logIn = async (
	{ cookies }: RequestEvent,
	{ username, password }: { username?: string; password?: string }
) => {
	// const { ok, status, data } = await client.POST('/auth/login', { username, password });
	// console.log(data);
	// let location: string;
	// if (ok) {
	// 	location = '/dashboard';
	// 	return { ok, location };
	// }
	// return { ok, status, data };
};
