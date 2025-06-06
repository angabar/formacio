import test from "@playwright/test";

test("Calendar validations", async ({ page }) => {
    const day = "15";
    const month = "6";
    const year = "2027";

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

    await page.pause();
});
