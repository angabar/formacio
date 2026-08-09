export default {
    // La función de mapDOM es la de definir unas referencias a nuestros
    // elementos importantes del componente para no tener que estar llamando y
    // re-llamado a querySelector constantemente en la clase del componente y
    // ahorrar así cálculos extra de computación
    mapDOM(scope) {
        return {
            title: scope.querySelector("h1"),
            subtitle: scope.querySelector("h3"),
        };
    },
    render(props) {
        return `${this.html(props)}
                ${this.css(props)}`;
    },
    html(props) {
        return `
            <h1>${props.title}</h1>
            <h3>${props.subtitle}</h3>
        `;
    },
    css(props) {
        return ``;
    },
};
