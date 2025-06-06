import test from "@playwright/test";

test("Shop app", async ({ page }) => {
    await page.goto("https://rahulshettyacademy.com/angularpractice");

    await page.getByLabel(/icecreams/i).check();
    await page.getByLabel("Employed").check();
    await page.getByLabel("Gender").selectOption("Female");

    await page.getByPlaceholder("Password").fill("abc123");
    await page.getByRole("button", { name: "Submit" }).click();
    await page.getByText(/submitted successfully/i).isVisible();
    await page.getByRole("link", { name: "Shop" }).click();

    await page
        .locator("app-card")
        .filter({
            hasText: "Nokia Edge",
        })
        .getByRole("button")
        .click();
});
