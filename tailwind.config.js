/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./public/*.html', './app/views/*/*.go'],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["light", "dark", "cupcake"],
  },
  plugins: [require("daisyui")],
}

