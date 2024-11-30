/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/components/**/*.{html,js,templ,go}",
    "./node_modules/flowbite/**/*.js"
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
  ],
}