import { writable } from 'svelte/store';
import { getToken } from './auth';
import type { User } from '../types/users';
import { invalidate, invalidateAll } from '$app/navigation';

export const groupsFilter = writable<string>();
export const createGroupResponse = writable<string>();

export const groupMembershipMasterCheckbox = writable<boolean>();
export const selectedGroupMemberships = writable<string[]>([]);

export async function getGroupMembers(groupId: string) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	const response = await fetch(`http://localhost:8080/group/${groupId}`, {
		headers,
		credentials: 'include'
	}).then((res) => res.json());
	return response as User[];
}

export async function createGroup(e: Event) {
	e.preventDefault();

	const formData = new FormData(e.target as HTMLFormElement);
	const name = formData.get('name');
	const enabled = formData.get('enabled')?.toString() === 'on';
	const is_admin = formData.get('is_admin')?.toString() === 'on';

	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	let response: number;
	try {
		response = await fetch('http://localhost:8080/group', {
			headers,
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify({
				name,
				enabled,
				is_admin
			})
		}).then((res) => res.status);
	} catch (e) {
		createGroupResponse.set('Error creating group, please try again later');
	}
	invalidate('admin:group');
}

export async function createGroupMembership(userIds: string[], groupId: string) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	const response = await fetch(`http://localhost:8080/group/${groupId}`, {
		headers,
		method: 'POST',
		credentials: 'include',
		body: JSON.stringify(userIds)
	}).then((res) => res.json());

}

export async function deleteGroupMembership(userIds: string[], groupId: string) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	const response = await fetch(`http://localhost:8080/group/${groupId}`, {
		headers,
		method: 'DELETE',
		credentials: 'include',
		body: JSON.stringify(userIds)
	}).then((res) => res.status);
}
