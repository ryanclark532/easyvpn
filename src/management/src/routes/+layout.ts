import { createAuthStore } from '$lib/stores/auth';
import { createUserStore } from '$lib/stores/users';

export const prerender = false;
export const ssr = false

const _authStore = await createAuthStore();
const _userStore = await createUserStore();

export { _authStore, _userStore };
