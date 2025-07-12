/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./cmd/web/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};
