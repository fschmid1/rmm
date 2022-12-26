
import { writable } from 'svelte/store';
import type { Device, User } from './types';
import type { Websocket } from './ws';

export const userStore = writable<User>(null);

export const deviceStore = writable<Device[]>([]);

export const device = writable<Device>(null);

export const ws = writable<Websocket>(null);

export const confirmModal = writable<{open: boolean, text: string, desision: boolean}>({open: false, text: '', desision: false});