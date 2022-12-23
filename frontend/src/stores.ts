
import { writable } from 'svelte/store';
import type { User } from './types';

export const userStore = writable<User>(null);