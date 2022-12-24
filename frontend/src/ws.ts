import { apiBase } from "./vars";

export class Websocket {
	private ws;
	private events: Map<string, (...data: any) => void> = new Map();

	constructor(userId: number, token: string) {
		this.ws = new WebSocket(`${apiBase.replace(/^http/, 'ws')}/ws/user/${userId}?token=${token}`);

		this.ws.onmessage = (rawEvent) => {
			try {
				const event = JSON.parse(rawEvent.data);

				if (this.events.has(event.event)) {
					this.events.get(event.event)(event.data);
				}
			} catch (error) {
				
			}
		}
	}

	public on(event: string, handler: (...data:any) => void) {
		this.events.set(event, handler);
	}
}