import { writable } from 'svelte/store';
import type { User } from '../types/types';

//auth
export const authResponseMessage = writable<string>();

//users
export const userValidationMessage = writable<string>();
export const users = writable<User[]>([]);
export const usersCount = writable<number>(0);
export const checkedUsers = writable<string[]>([]);
export const changedUsers = writable<User[]>([]);

//misc
export const checkboxMaster = writable<boolean>(false);