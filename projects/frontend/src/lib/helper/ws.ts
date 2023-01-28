import {  wsBase } from '../../vars';

export class Websocket {
    private ws;
    private events: Map<string, (...data: any) => void> = new Map();

    private userId: number;

    constructor(userId: number) {
        this.userId = userId;

        this.connect();
    }

    private connect() {
        this.ws = new WebSocket(`${wsBase}/ws/user/${this.userId}`);

        this.ws.onmessage = (rawEvent) => {
            try {
                const event = JSON.parse(rawEvent.data);

                if (this.events.has(event.event)) {
                    this.events.get(event.event)(event.data);
                }
            } catch (error) {
                console.log(error);
            }
        };
        this.ws.onclose = () => {
            setTimeout(() => {
                this.connect();
            }, 5000);
        };
    }

    public on(event: string, handler: (...data: any) => void) {
        this.events.set(event, handler);
    }

    public off(event: string) {
        this.events.delete(event);
    }
}
