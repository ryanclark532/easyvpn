import { typedFetch } from "$lib/fetch";
import { writable } from "svelte/store";
import type { GetGroupsResponse, Group } from "../../types/groups";
import { getToken } from "./auth";

export const groupsFilter = writable<string>()

export async function createGroupsStore(){
    const { set, update, subscribe} = writable<Group[]>([])
const groups = await getGroups()
if(!(groups instanceof Error)){
    set(groups)
}

    return{
        set,
        update,
        subscribe
    }
}



async function getGroups() {
	const headers = new Headers();
	const token = getToken();
	if (!token) {
		return new Error('No Token Provided');
	}
	headers.append('Authorization', `Bearer ${token}`);
	const response = await typedFetch<GetGroupsResponse>('http://localhost:8080/api/groups', {
		method: 'GET',
		headers
	});

	if (response instanceof Error || response.status >= 400) {
		return new Error('Error retrieving users, please try again later');
	}
	return response.data.groups || [];
}