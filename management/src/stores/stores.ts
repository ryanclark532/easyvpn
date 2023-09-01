import { writable } from 'svelte/store';

export const authResponseMessage = writable<string>();
export const checkboxMaster = writable<boolean>(false);
