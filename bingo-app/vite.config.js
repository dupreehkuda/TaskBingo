import { sveltekit } from '@sveltejs/kit/vite';

/** @type {import('vite').UserConfig} */
const config = {
	plugins: [sveltekit()],
	server:{
		port:3000,
		strictPort:false,
	},
	preview:{
	port:4173,
		strictPort:false,
	}
};

export default config;
