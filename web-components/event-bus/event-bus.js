class EventBus {
    constructor() {
        this.target = new EventTarget();
    }

    subscribe(eventName, callback) {
        const listener = (event) => callback(event.detail);
        this.target.addEventListener(eventName, listener);

        return () => this.target.removeEventListener(eventName, listener);
    }

    publish(eventName, detail) {
        this.target.dispatchEvent(new CustomEvent(eventName, { detail }));
    }
}

const eventBus = new EventBus();

export default eventBus;