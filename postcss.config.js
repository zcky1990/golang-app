// Deprecated: PostCSS is no longer in the pipeline.
// Tailwind CSS v4 is applied through the Vite plugin (@tailwindcss/vite in vite.config.js),
// so `postcss`, `postcss-loader`, `autoprefixer` and `@storybook/addon-postcss` were removed.
// This file is kept only as a no-op so nothing picks up the old Tailwind v3 PostCSS plugin.
// It can be deleted safely.
module.exports = {
  plugins: {},
};
