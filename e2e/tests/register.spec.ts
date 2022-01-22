import { test, expect } from '@playwright/test';

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
});

