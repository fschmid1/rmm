import { confirmModal } from "../../stores"

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