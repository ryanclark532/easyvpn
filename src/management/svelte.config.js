import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/kit/vite';

export default {
	preprocess: [vitePreprocess()],
	kit: {
		adapter: adapter({
			pages: '../app',
			assets: '../app',
			fallback: undefined,
			precompress: false,
			strict: true
		})
	}
};
