import { writable } from 'svelte/store';
import type { Websocket } from './lib/helper/ws';
import type { Device, User } from './types';

export const userStore = writable<User>(null);

export const device = writable<Device>(null);

export const ws = writable<Websocket>(null);
