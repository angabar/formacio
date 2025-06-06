import test, { expect } from "@playwright/test";

test("Browser context test", async ({ browser }) => {
    const context = await browser.newContext();
    const page = await context.newPage();

    await page.goto("https://rahulshettyacademy.com/loginpagePractise/");
});

test("Page test", async ({ page }) => {
    await page.goto("https://www.google.com/");

    await expect(page).toHaveTitle("Google");
});

test("Can fill fields", async ({ page }) => {
    const userNameInput = page.locator("#username");
    const signInButton = page.locator("#signInBtn");
    const cardTitles = page.locator(".card-body a");

    await page.goto("https://rahulshettyacademy.com/loginpagePractise/");

    await userNameInput.fill("rahulshetty");
    await page.locator("[type='password']").fill("learning");
    await signInButton.click();

    await expect(page.locator("[style*='block']")).toContainText(
        "Incorrect username/password.",
    );

    await userNameInput.fill("");
    await userNameInput.fill("rahulshettyacademy");
    await signInButton.click();

    // console.log(await cardTitles.first().textContent());

    const allCardTitles = await cardTitles.allTextContents();
    console.log(allCardTitles);
});

test("UI Controls", async ({ page }) => {
    const userNameInput = page.locator("#username");
    const signInButton = page.locator("#signInBtn");
    const dropdworn = page.locator("select.form-control");
    const documentLink = page.locator("[href*='documents-request']");

    await page.goto("https://rahulshettyacademy.com/loginpagePractise/");

    await dropdworn.selectOption("consult");

    await page.locator(".radiotextsty").last().click();
    await page.locator("#okayBtn").click();

    await expect(page.locator(".radiotextsty").last()).toBeChecked();

    await page.locator("#terms").click();
    await expect(page.locator("#terms")).toBeChecked();
    expect(await page.locator("#terms").isChecked()).toBeFalsy();

    await expect(documentLink).toHaveAttribute("class", "blinkingText");
});

test("Child windows", async ({ browser }) => {
    const context = await browser.newContext();
    const page = await context.newPage();

    const userNameInput = page.locator("#username");
    const documentLink = page.locator("[href*='documents-request']");

    await page.goto("https://rahulshettyacademy.com/loginpagePractise/");

    const [newPage] = await Promise.all([
        context.waitForEvent("page"),
        documentLink.click(),
    ]);

    const text = await newPage.locator(".red").textContent();
    const arrayText = text.split("@");
    const domain = arrayText[1].split(" ")[0];

    await userNameInput.fill(domain);
});
