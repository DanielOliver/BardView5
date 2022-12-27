module.exports = {
    extends: ['eslint:recommended', 'plugin:@typescript-eslint/recommended'],
    ignorePatterns: [".eslintrc.cjs", "*config.js", "temp.js", "**/vendor/*.js"],
    parser: '@typescript-eslint/parser',
    plugins: ['@typescript-eslint'],
    root: true,
  };