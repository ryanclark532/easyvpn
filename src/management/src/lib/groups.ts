import { writable } from 'svelte/store';
import { getToken } from './auth';
import type { User } from '../types/users';

export const groupsFilter = writable<string>();

export async function getGroupMembers(groupId: number) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	const response = await fetch(`http://localhost:8080/group/${groupId}`, {
		headers,
		credentials: 'include'
	}).then((res) => res.json());
	return response as User[];
}

export async function createGroup(e: Event) {
    e.preventDefault()

    console.log("hello?");
}

