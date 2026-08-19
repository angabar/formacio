import WorkoutExercise from "../workout-exercise/workout-exercise.js";

export default {
    render(props) {
        return `
            ${this.css(props)}
            ${this.html(props)}
        `;
    },
    html(props) {
        let exercises = ``;

        props.exercises.forEach((exercice) => {
            console.log(
                "WorkoutExercise.toAttributeString(exercice)__ ",
                WorkoutExercise.toAttributeString(exercice),
            );

            exercises += `<workout-exercise class="${exercice.type}" ${WorkoutExercise.toAttributeString(exercice)}></workout-exercise>`;
        });

        return `
            <h1>Exercises</h1>
            <div id="container">
                ${exercises}
            </div>
        `;
    },
    css(props) {
        return `
            <style>
                host {
                    display: flex;
                    flex-direction: column;
                }

                #container {
                    overflow-y: scroll;
                    height: calc(100% - 60px);
                }
            </style>
        `;
    },
};
