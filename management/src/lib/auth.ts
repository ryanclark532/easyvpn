import { writable } from 'svelte/store';

export const token = writable<string>();
export const is_admin = writable<boolean>()

export async function handleLogin(e: any): Promise<void> {
	const formData = new FormData(e.target);
	const body: { [key: string]: string } = {};

	formData.forEach((value, key) => {
		body[key] = value.toString();
	});

	const response = await fetch('http://localhost:8080/user/sign-in', {
		body: JSON.stringify(body),
		method: 'POST'
	});
	if (!response.ok) {
	}

	const json = await response.json();
	console.log(json)
	token.set(json.token);
	is_admin.set(json.is_admin === 'true')

	if(json.is_admin === 'true'){
		window.location.href ="/admin"
	} else {
		window.location.href ="/"
	}
}
