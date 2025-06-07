import test, { expect } from "@playwright/test";

test("Calendar validations", async ({ page }) => {
    const day = "15";
    const month = "6";
    const year = "2027";

    const expectedList = [month, day, year];

    await page.goto("https://rahulshettyacademy.com/seleniumPractise/#/offers");

    await page.locator(".react-date-picker__inputGroup").click();
    await page.locator(".react-calendar__navigation__label").click();
    await page.locator(".react-calendar__navigation__label").click();
    await page.getByText(year).click();
    await page
        .locator(".react-calendar__year-view__months__month")
        .nth(+month - 1)
        .click();
    await page.locator("//abbr[text()='" + day + "']").click();

    const inputs = page.locator(".react-date-picker__inputGroup input");

    for (let index = 0; index < inputs.length; index++) {
        const input = inputs[index];

        const inputValue = input.getAttribute("value");

        expect(inputValue).toEqual(expectedList[index]);
    }
});
