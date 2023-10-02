import type { Updater } from 'svelte/store';
import { writable } from 'svelte/store';
import { getToken } from '$lib/auth';
import { typedFetch } from '$lib/fetch';
import type {
	CreateUserRequest,
	GetUsersResponse,
	UpdateUserResponse,
	User,
	UserCreationResponse
} from '../../types/users';

export async function createUserStore() {
	const { subscribe, set, update } = writable<User[]>([]);
	const users = await getUsers();
	if (!(users instanceof Error)) {
		set(users);
	}
	return {
		subscribe,
		update,
		set,
		delete: (u: User[]) => deleteUser(u, update),
		updateUser: (u: User[]) => updateUser(u, update),
		setTempPw: (u: User[]) => setTempPw(u),
		create: (e: Event) => create(e, update)
	};
}

async function getUsers() {
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${getToken()}`);
	const response = await typedFetch<GetUsersResponse>('http://localhost:8080/api/user', {
		method: 'GET',
		headers
	});

	if (response.status >= 400) {
		return new Error('Error retrieving users, please try again later');
	}
	return response.data.users;
}

async function deleteUser(u: User[], update: (this: void, updater: Updater<User[]>) => void) {
	const body = {
		ID: u.map((u) => u.ID)
	};
	if (body.ID.length === 0) {
		return new Error('No Users Selected');
	}
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${getToken()}`);

	const response = await fetch('http://localhost:8080/api/user', {
		body: JSON.stringify(body),
		method: 'DELETE',
		headers
	});

	if (response.status >= 400) {
		return new Error('Error retrieving users, please try again later');
	}
	update((prev: User[]) => {
		return prev.filter((u) => {
			return !body.ID.includes(u.ID);
		});
	});
}

async function updateUser(users: User[], update: (this: void, updater: Updater<User[]>) => void) {
	const body = {
		users
	};
	if (body.users.length === 0) {
		return new Error('No Users Selected');
	}
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${getToken()}`);
	const response = await typedFetch<UpdateUserResponse>('http://localhost:8080/api/user', {
		body: JSON.stringify(body),
		method: 'PUT',
		headers
	});

	if (response.status >= 400) {
		return new Error('Error updating users, please try again later');
	}

	update((prev: User[]) => {
		return prev.map((u) => {
			if (!body.users.includes(u)) {
				return u;
			}
			return body.users.find((x) => x.ID === u.ID) ?? u;
		});
	});
}

async function setTempPw(u: User[]) {
	const body = {
		ID: u.map((u) => u.ID)
	};
	if (body.ID.length === 0) {
		return new Error('No Users Selected');
	}
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${getToken()}`);

	const response = await fetch('http://localhost:8080/api/user/set-temporary-password', {
		body: JSON.stringify(body),
		method: 'PUT',
		headers
	});

	if (response.status >= 400) {
		return new Error('Error setting users temporary password, please try again later');
	}
}

async function create(e: Event, update: (this: void, updater: Updater<User[]>) => void) {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const body = {
		username: formData.get('username') as string,
		name: formData.get('name') as string,
		password: formData.get('password') as string,
		is_admin: (formData.get('isAdmin') as string) === 'on',
		enabled: (formData.get('enabled') as string) === 'on'
	};

	if (!validateNewUser(body)) {
		return new Error('Fields Are Invalid');
	}
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${getToken()}`);

	const response = await typedFetch<UserCreationResponse>('http://localhost:8080/api/user', {
		body: JSON.stringify(body),
		method: 'POST',
		headers
	});
	if (response.status === 409) {
		return new Error(`User with username ${body.username} already exists`);
	}
	if (response.status >= 400) {
		return new Error(`Something went wrong while creating ${body.username}`);
	}

	update((users) => [...users, response.data.user]);
}

function validateNewUser(user: CreateUserRequest) {
	return true;
}
