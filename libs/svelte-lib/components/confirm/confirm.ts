
import { writable } from 'svelte/store';


export const confirmModal = writable<{open: boolean, text: string, desision: boolean}>({open: false, text: '', desision: false});


export const customConfirm = async (text: string) => {
	confirmModal.set({open: true, text, desision: false});
	return new Promise<boolean>((resolve) => {
		let sub = confirmModal.subscribe((value) => {
			if (!value.open) {
				sub();
				resolve(value.desision);
			}
		});
	});
}