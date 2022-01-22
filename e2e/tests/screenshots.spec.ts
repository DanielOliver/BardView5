import { test, expect } from '@playwright/test';
import * as crypto from 'crypto';


// function uuidv4() {
//   // @ts-ignore
//   return ([1e7]+-1e3+-4e3+-8e3+-1e11).replace(/[018]/g, c =>
//   // @ts-ignore
//   (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
//   );
// }

function createGuid(): string {
  return crypto.randomBytes(16).toString('hex');
}

test('landing page to register', async ({ page }) => {
  await page.goto('http://proxy.local/');
  await page.locator('text=BardView5').click();
  await expect(page).toHaveTitle(/BardView5/);
  await page.screenshot({ path: 'screenshots/unauth_landing.png' });

  await page.locator('text=Login').first().click();
  await page.locator('text=ID').isVisible();
  await page.waitForTimeout(200);
  await page.screenshot({ path: 'screenshots/unauth_login.png' });
  
  await page.locator('text=Register').click();
  await page.locator('text=Username').isVisible();
  await page.waitForTimeout(200);
  await page.screenshot({ path: 'screenshots/unauth_register.png' });

  const userId = (new Date()).getTime().toString();
  const password = createGuid(); // (Math.floor(Math.random() * 10000) * 10000 + Math.floor(Math.random() * 5000)).toString();
  const email = userId + "@test.com";

  const emailField = await page.locator('input[name="traits.email"]');
  const usernameField = await page.locator('input[name="traits.username"]');
  const passwordField = await page.locator('input[name="password"]');
  await emailField.click();
  await emailField.type(email)
  await passwordField.click();
  await passwordField.type(password);
  await usernameField.click();
  await usernameField.type(userId);
  await page.locator('.btn-primary').click();
   
  await page.locator('text=Worlds').isVisible();
  await page.waitForTimeout(2200);
  await page.screenshot({ path: 'screenshots/registered.png' });
});

