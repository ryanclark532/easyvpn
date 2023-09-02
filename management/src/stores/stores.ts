import { writable } from 'svelte/store';

//auth
export const authResponseMessage = writable<string>();

//users
export const userValidationMessage = writable<string>();

//misc
export const checkboxMaster = writable<boolean>(false);
