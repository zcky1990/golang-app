// Storybook reuses vite.config.js (Vue + Tailwind CSS v4 plugins), so importing the
// stylesheet here is enough to get the compiled Tailwind output in stories.
import '../static/css/tailwind.css';

/** @type { import('@storybook/vue3').Preview } */
const preview = {
  parameters: {
    docs: {
      toc: true, // 👈 Enables the table of contents
    },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
};

export default preview;
