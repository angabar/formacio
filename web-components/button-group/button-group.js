class ButtonGroup extends HTMLElement {
    connectedCallback() {
        this.innerHTML = `
            <style>
                .change-state-button.disabled {
                    opacity: 0.5;
                }
            </style>
            <button class="disabler-button">Disable other button</button>
            <button class="change-state-button">Enabled</button>
        `;

        document
            .querySelector(".disabler-button")
            .addEventListener("click", (event) =>
                this.changeButtonState(event),
            );
    }

    changeButtonState(event) {
        const changeStateButton = document.querySelector(
            ".change-state-button",
        );

        if (changeStateButton.classList.contains("disabled")) {
            changeStateButton.classList.remove("disabled");
        } else {
            changeStateButton.classList.add("disabled");
        }
    }
}

if (!customElements.get("button-group")) {
    customElements.define("button-group", ButtonGroup);
}
