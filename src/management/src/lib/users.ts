import type { User } from '../types/users';
import { writable } from 'svelte/store';

export const searchFilter = writable<string>();
export const selectedUsers = writable<User[]>([]);
export const masterCheckbox = writable<boolean>(false);
