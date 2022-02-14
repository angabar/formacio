import React, { useState } from "react";

export const replaceCamelCaseWithSpaces = (color) => {
    return color.replace(/\B([A-Z])\B/g, " $1");
};

const App = () => {
    const [buttonColor, setButtonColor] = useState("MediumVioletRed");
    const [disabledMode, setDisabledMode] = useState(false);

    const newButtonColor =
        buttonColor === "MediumVioletRed" ? "MidnightBlue" : "MediumVioletRed";

    return (
        <div className="App">
            <button
                onClick={() => setButtonColor(newButtonColor)}
                style={{ backgroundColor: disabledMode ? "gray" : buttonColor }}
                disabled={disabledMode}
            >
                Change to {replaceCamelCaseWithSpaces(newButtonColor)}
            </button>
            <label htmlFor="diable-color-button">Disable color button</label>
            <input
                id="diable-color-button"
                type="checkbox"
                onClick={() => setDisabledMode(!disabledMode)}
            />
        </div>
    );
};

export default App;
