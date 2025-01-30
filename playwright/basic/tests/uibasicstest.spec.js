const { test, expect } = require("@playwright/test");

test("Fist playwright test", async ({ page }) => {
    // page nos proporciona por defecto estos dos pasos

    // const context = await browser.newContext();
    // const page = await context.newPage();

    await page.goto("https://rahulshettyacademy.com/loginpagePractise/");

    const title = await page.title();
    await expect(page).toHaveTitle(title);

    const userName = page.locator("#username");
    const userPassword = page.locator("[type='password']");
    const submitButton = page.locator("#signInBtn");

    await userName.fill("jajaja");
    await userPassword.fill("learning");
    await submitButton.click();

    // const warningText = await page.locator("[style*='block']").textContent();
    await expect(page.locator("[style*='block']")).toContainText("Incorrect");

    await userName.fill("");

    await userName.fill("rahulshettyacademy");
    await submitButton.click();

    const cardTitle = page.locator(".card-body a");

    // Cont nth accedemos al indice del elemento que queremos
    // await cardTitle.nth(0).textContent();
    await cardTitle.first().textContent();

    // Para obtener todos los titulos en un array usamos allTextContents pero
    // este metodo es curioso, debemos usar un elmento que espere a que salga
    // todo antes (invocando a un first por ejemplo) de este o de lo contrario
    // devolvera 0 elementos siempre ya que allTextContents por defecto no
    // espera a que el DOM este cargado o que el elemento sea visible

    // await cardTitle.allTextContents();
});
