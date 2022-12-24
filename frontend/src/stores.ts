
import { writable } from 'svelte/store';
import { type Device, type User } from './types';
import type { Websocket } from './ws';

export const userStore = writable<User>(null);

export const deviceStore = writable<Device[]>([]);

export const ws = writable<Websocket>(null);