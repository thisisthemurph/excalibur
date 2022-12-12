/** @type {import('tailwindcss').Config} */
module.exports = {
  purge: ["./src/**/*.{js,jsx,ts,tsx}", "./public/index.html"],
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        brand: {
          one: "color: rgb(129 140 248);", // indigo-400
          dark: "rgb(67 56 202)", // indigo-700
          DEFAULT: "rgb(79 70 229)", // indigo-600
        },
      },
      spacing: {
        wrap: "2rem",
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
