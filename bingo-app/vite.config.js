import { sveltekit } from '@sveltejs/kit/vite';

/** @type {import('vite').UserConfig} */
const config = {
	plugins: [sveltekit()],
	// Repo-root .env is the single source of truth for both docker-compose
	// and Vite. envDir tells Vite to load it from one level up.
	envDir: '..',
	server: {
		port: 3000,
		strictPort: false,
		// Proxy /api and /ws to the local backend so the browser sees
		// everything as same-origin. Avoids CORS friction (Safari ITP /
		// preflight caching) entirely during dev.
		proxy: {
			'/api': {
				target: 'http://localhost:8082',
				changeOrigin: true,
				ws: true,
			},
		},
	},
	preview: {
		port: 3000,
		strictPort: false,
	},
};

export default config;
