import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/kit/vite';

export default {
	preprocess: [vitePreprocess()],
	kit: {
		adapter: adapter({
			pages: '../../dist/static',
			assets: '../../dist/static',
			fallback: undefined,
			precompress: false,
			strict: true
		})
	}
};
