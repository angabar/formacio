import { render, screen, fireEvent } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";

import App, { replaceCamelCaseWithSpaces } from "./App";

test("button has correct initial color", () => {
    render(<App />);

    // La invocacion de render nos da acceso a
    // la variable global screen
    const colorButton = screen.getByRole("button", {
        name: "Change to Midnight Blue",
    });

    expect(colorButton).toHaveStyle({ backgroundColor: "MediumVioletRed" });
});

test("button turns blue when clicked", () => {
    render(<App />);

    const colorButton = screen.getByRole("button", {
        name: "Change to Midnight Blue",
    });

    expect(colorButton).toHaveStyle({ backgroundColor: "MediumVioletRed" });

    fireEvent.click(colorButton);

    expect(colorButton).toHaveStyle({ backgroundColor: "MidnightBlue" });
    expect(colorButton).toHaveTextContent("Change to Medium Violet Red");
});

test("initial conditions", () => {
    render(<App />);

    const colorButton = screen.getByRole("button", {
        name: "Change to Midnight Blue",
    });

    expect(colorButton).toBeEnabled();

    const checkbox = screen.getByRole("checkbox");

    expect(checkbox).not.toBeChecked();
});

test("button is disabled and turns gray when click checkbox", () => {
    render(<App />);

    const colorButton = screen.getByRole("button", {
        name: "Change to Midnight Blue",
    });
    const checkbox = screen.getByRole("checkbox", {
        name: "Disable color button",
    });
    expect(colorButton).toHaveStyle({ backgroundColor: "MediumVioletRed" });

    fireEvent.click(colorButton);
    fireEvent.click(checkbox);

    expect(checkbox).toBeChecked();
    expect(colorButton).toBeDisabled();
    expect(colorButton).toHaveStyle({ backgroundColor: "gray" });

    fireEvent.click(checkbox);

    expect(checkbox).not.toBeChecked();
    expect(colorButton).toBeEnabled();
    expect(colorButton).toHaveStyle({ backgroundColor: "MidnightBlue" });
});

test("replaceCamelCaseWithSpaces", () => {
    expect(replaceCamelCaseWithSpaces("Red")).toBe("Red");
    expect(replaceCamelCaseWithSpaces("MidnightBlue")).toBe("Midnight Blue");
    expect(replaceCamelCaseWithSpaces("MediumVioletRed")).toBe(
        "Medium Violet Red",
    );
});
