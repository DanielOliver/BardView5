import { test, expect } from '@playwright/test';
import * as crypto from 'crypto';

function createGuid(): string {
  return crypto.randomBytes(16).toString('hex');
}

test('landing page to register', async ({ page }) => {
  await page.goto('http://proxy.local/');
  await page.locator('text=BardView5').click();
  await expect(page).toHaveTitle(/BardView5/);
  await page.screenshot({ path: 'screenshots/unauth_landing.png' });

  await page.locator('text=Login').first().click();
  await page.waitForLoadState();
  const idSelector = await page.waitForSelector('text=ID');
  await idSelector.isVisible();
  await page.screenshot({ path: 'screenshots/unauth_login.png' });

  await page.locator('text=Register').click();
  await page.waitForLoadState();
  const usernameSelector = await page.waitForSelector('text=Username');
  await usernameSelector.isVisible();
  await page.screenshot({ path: 'screenshots/unauth_register.png' });

  const userId = "user" + (new Date()).getTime().toString();
  const password = createGuid();
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

  await page.waitForLoadState();
  const settingSelector = await page.waitForSelector('text=Settings');
  await settingSelector.isVisible();

  const layoutSelector = await page.waitForSelector(`text=${userId}`);
  await layoutSelector.isVisible();

  await page.screenshot({ path: 'screenshots/registered.png' });

});

