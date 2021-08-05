import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		apiurl: 'http://localhost:8080'
	}
});

export default app;