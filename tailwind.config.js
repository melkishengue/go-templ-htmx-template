/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "selector",
  content: ["./views/**/*.templ", "./partials/**/*.templ"],
  theme: {
    extend: {
      colors: {
        primary: "#c2e7ff"
      }
    }
  },

  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: [
      {
        myapp: {
          primary: "#c2e7ff",
          error: "#fd5c63"
        }
      }
    ]
  }
};
