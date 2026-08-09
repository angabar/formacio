export default {
    mapDOM(scope) {
        return {
            thumb: scope.querySelector(".thumb"),
            bgOverlay: scope.querySelector(".bg-overlay"),
        };
    },
    render(props) {
        return `
            ${this.html(props)}
            ${this.css(props)}
        `;
    },
    html(props) {
        return `<div class='bg-overlay'></div><div class='thumb'></div>`;
    },
    css(props) {
        return `
            <style>
                wcia-slider {
                    display: inline-block;
                    position: relative;
                    border-radius: 3px;
                    height: 50px;
                    width: 400px;
                }

                .bg-overlay {
                    width: 100%;
                    height: 100%;
                    position: absolute;
                    border-radius: 3px;
                    background-color: red;
                }

                .thumb {
                    margin-top: -1px;
                    width: 5px;
                    height: calc(100% - 5px);
                    position: absolute;
                    border-style: solid;
                    border-width: 3px;
                    border-color: white;
                    border-radius: 3px;
                    pointer-events: none;
                }
            </style>
        `;
    },
};
