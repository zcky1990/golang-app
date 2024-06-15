/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/views/**/*.{vue,html,js}",
    "./static/javascript/**/*.{vue,html,js}"
  ],
  //comment this when you want to build prod
  safelist: [
    {
      pattern: /./, // the "." means "everything"
    },
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

