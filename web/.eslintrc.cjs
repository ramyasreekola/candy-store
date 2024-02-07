module.exports = {
  root: true,
  env: { browser: true, es2021: true },
  extends: [
    "eslint:recommended",
    "plugin:react/recommended",
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react-hooks/recommended",
  ],
  ignorePatterns: [
    "dist",
    ".eslintrc.cjs",
    "vite.config.ts",
    "src/vite-env.d.ts",
    "**/*.test.tsx",
  ],
  parser: "@typescript-eslint/parser",
  plugins: ["react"],
  settings: {
    "react": {
      "version": "detect",
    },
  }
};
