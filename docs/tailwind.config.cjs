const defaultTheme = require("tailwindcss/defaultTheme");

const config = {
	content: ["./src/**/*.{html,js,svelte,ts}"],
	theme: {
		extend: {
			typography: () => ({
				DEFAULT: {
					css: {
						"code::before": {
							content: "none",
						},
						"code::after": {
							content: "none",
						},
						"blockquote p:first-of-type::before": {
							content: "none",
						},
						"blockquote p:last-of-type::after": {
							content: "none",
						},
						code: {
							padding: "4px 8px",
							fontWeight: "400",
						},
					},
				},
			}),
			colors: {
				main: {
					primary: "#022964",
					secondary: "#1d7aeb",
					black: "#1B1B1B",
				},
				gray: {
					100: "#FCFAFA",
					200: "#EBEBEB",
				},
				blue: {
					100: "#EFF6FF",
					200: "#eaf3ff",
					500: "#03499F",
				},
				green: {
					100: "#DAF6DB",
					500: "#027149",
				},
				red: {
					100: "#FFEFEF",
					400: "#D82F2F",
				},
				yellow: {
					100: "#FFEAB6",
					500: "#B68505",
				},
			},
			fontFamily: {
				sans: ["Ubuntu", ...defaultTheme.fontFamily.sans],
				heading: ["Lato", ...defaultTheme.fontFamily.sans],
				mono: ["Ubuntu Mono", ...defaultTheme.fontFamily.mono],
			},
			borderWidth: {
				1.5: "1.5px",
			},
			spacing: {
				0.5: "0.125rem",
				1.5: "0.375rem",
				2.5: "0.625rem",
				3.5: "0.875rem",
				4.5: "1.125rem",
				13: "3.25rem",
			},
			borderRadius: {
				"md-alt": "0.625rem",
			},
			minHeight: {
				"screen-with-footer": "calc(100vh - 48px)",
			},
			boxShadow: {
				"sm-alt": "0px 2px 4px rgba(0, 0, 0, 0.25)",
				"md-alt": "0px 2px 8px rgba(0, 0, 0, 0.1)",
			},
		},
	},
	plugins: [require("@tailwindcss/typography")],
};

module.exports = config;
