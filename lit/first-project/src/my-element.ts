import { LitElement, html } from "lit";
import { customElement, state } from "lit/decorators.js";
import { increment, store } from "./redux/store";
import type { Unsubscribe } from "@reduxjs/toolkit";
import { api } from "./api/api";

@customElement("my-element")
export class MyElement extends LitElement {
    @state()
    ubsuscribe?: Unsubscribe;
    @state()
    querySubscription?: unknown;

    connectedCallback() {
        super.connectedCallback();

        this.ubsuscribe = store.subscribe(() => {
            this.requestUpdate();
        });

        this.loadUser(1);
    }

    disconnectedCallback(): void {
        super.disconnectedCallback();

        this.ubsuscribe?.();
        this.querySubscription?.unsubscribe?.();
    }

    protected render() {
        const state = store.getState();

        return html`
            <div>hola! ${state.counter.value}</div>
            <button @click=${() => store.dispatch(increment())}>jaja</button>
            ${this.renderUsers()}
        `;
    }

    private loadUser(id: number) {
        this.querySubscription?.unsubscribe();

        this.querySubscription = store.dispatch(
            api.endpoints.getUser.initiate(id),
        );
    }

    private renderUsers() {
        const result = api.endpoints.getUser.select(1)(store.getState());

        if (result.isLoading) {
            return html`<p>Cargando...</p>`;
        }

        if (result.isError) {
            return html`<p>Error</p>`;
        }

        if (!result.data) {
            return html`<p>Sin datos</p>`;
        }

        return html`<span>${result.data.name}</span>`;
    }
}
