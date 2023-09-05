import { users, usersCount, userValidationMessage } from '../stores/stores';
import type { UserCreationResponse, UserResponse } from '../types/types';
import { typedFetch } from '$lib/fetch';

function validateUser(name: string, username: string, password: string): boolean {
	if (!name) {
		userValidationMessage.set('Please Enter A Name');
		return false;
	} else if (!username) {
		userValidationMessage.set('Please Enter A Username');
		return false;
	} else if (!password) {
		userValidationMessage.set('Please Enter A Password');
		return false;
	}
	userValidationMessage.set('');
	return true;
}

export async function createUser(
	name: string,
	username: string,
	password: string,
	is_admin: boolean,
	enabled: boolean
) {
	const valid = validateUser(name, username, password);
	if (!valid) {
		return false;
	}
	const body = {
		username,
		name,
		password,
		is_admin,
		enabled
	};

	const response = await typedFetch<UserCreationResponse>('http://localhost:8080/user', {
		method: 'POST',
		body: JSON.stringify(body)
	});
	if (response.status === 409) {
		userValidationMessage.set('User Already Exists');
		return false;
	} else if (response.status !== 201) {
		userValidationMessage.set('An Error Occurred, Please Try Again');
		return false;
	}
	alert(`User ${username} Created`);
	users.update((prev) => [...prev, response.data.user]);
	return true;
}

export async function getUsers() {
	const response = await typedFetch<UserResponse>('http://localhost:8080/user');
	if (response.status !== 200) {
		//Some error occurred
		return;
	}
	users.set(response.data.users);
	usersCount.set(response.data.count);
	return response.data.users;
}
