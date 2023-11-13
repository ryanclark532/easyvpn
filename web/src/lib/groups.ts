import { writable } from 'svelte/store';
import { getToken } from './auth';
import type { User } from '../types/users';
import { invalidate } from '$app/navigation';
import type { Group } from '../types/groups';
export const groupsFilter = writable<string>();
export const createGroupResponse = writable<string>();

export const groupMembershipMasterCheckbox = writable<boolean>();
export const selectedGroupMemberships = writable<string[]>([]);
export const selectedGroups = writable<Group[]>([]);

export async function getGroupMembers(groupId: string) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	const response = await fetch(`http://localhost:8080/group/membership/${groupId}`, {
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
	await invalidate('admin:group');
}

export async function createGroupMembership(userIds: string[], groupId: string) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	await fetch(`http://localhost:8080/group/membership/${groupId}`, {
		headers,
		method: 'POST',
		credentials: 'include',
		body: JSON.stringify(userIds)
	});

	await invalidate('admin:group');
}

export async function deleteGroupMembership(userIds: string[], groupId: string) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	await fetch(`http://localhost:8080/group/membership/${groupId}`, {
		headers,
		method: 'DELETE',
		credentials: 'include',
		body: JSON.stringify(userIds)
	});

	await invalidate('admin:group');
}

export async function deleteGroup(groupId: string) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	await fetch(`http://localhost:8080/group/${groupId}`, {
		headers,
		method: 'DELETE',
		credentials: 'include'
	});
	await invalidate('admin:group');
}

export async function updateGroup(groupId: string, e: Event) {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const name = formData.get('name')?.toString();
	const is_admin = formData.get('is_admin')?.toString() === 'on';
	const enabled = formData.get('enabled')?.toString() === 'on';

	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	const response = await fetch(`http://localhost:8080/group/${groupId}`, {
		headers,
		body: JSON.stringify({ name, is_admin, enabled }),
		method: 'PUT'
	}).then((res) => res.status);
	if (response === 409) {
		return new Error(`There is already a group with name ${name}`);
	}
	if (response >= 400) {
		return new Error('There was an error updating the group, please try again later');
	}
	await invalidate('admin:group');
}
