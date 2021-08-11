<script>
	import { onMount } from 'svelte';
	import Router from 'svelte-spa-router'
	import {wrap} from 'svelte-spa-router/wrap'
	import CreateForm from './create/CreateForm.svelte'
	import DisplayForm from './display/DisplayForm.svelte'

	export let apiurl

	const routes = {
		'/create': wrap({
			component: CreateForm,
			props: {
				endpoint: apiurl + "/forms/"
			}
    	}),
		'/forms/:formid': wrap({
			component: DisplayForm,
			props: {
				endpoint: apiurl + "/forms/"
			}
    	})
	}

	let forms = []

	const getForms = async () => {
		var resp = await fetch(apiurl + "/forms/")
		forms = await resp.json()
	}

	onMount(getForms)

	const routeEvent = (event) => {
		switch(event.detail.type) {
			case 'form-created':
				forms.push(event.detail.form)
				forms = forms
				break
			case 'form-deleted':
				getForms()
				break
		}
		
	}
</script>

<main>
	<div id="leftpanel">
		<!-- <h1><img src="https://coopgo.fr/images/coopgo-logo-whitegreen.svg" alt="COOPGO" /> Forms Service</h1> -->
		<h1>COOPGO Forms Service</h1>
		<ul id="mainmenu">
			<li><a href="#/create">+ Create a new form</a></li>
		</ul>

		<h2>Forms :</h2>
		<ul id="formslist">
			{#each forms as form}
				<li><a href="#/forms/{form.id}">{form.id}</a></li>
			{/each}
		</ul>
	</div>
	<div id="content">
		<Router {routes} on:routeEvent={routeEvent} />
	</div>
</main>

<style>
	main {
		height: 100%;
		display: flex;
		flex-direction: row;
	}

	#leftpanel {
		width: 300px;
		background-color: #243887;
		color: #fff;
		overflow: auto;
	}

	#leftpanel a {
		display: block;
		text-decoration: none;
		color: #fff;
		padding: 10px;
	}

	#leftpanel a:hover {
		background-color: rgba(255,255,255,0.1);
	}

	#leftpanel h1 {
		color: #fff;	
		border-bottom: 1px solid rgba(255,255,255,0.1);
		font-size: 1em;
		font-weight: normal;
		margin: 0;
		padding: 10px;
		text-align: center;
	}

	#leftpanel h1 img {
		height: 1em;
	}

	#leftpanel h2 {
		margin: 0;
		padding: 10px;
		font-size: 1em;
	}

	#leftpanel ul {
		margin: 0px;
		padding: 0px;
	}

	#leftpanel ul li {
		/* padding: 10px; */
		list-style: none;
	}

	/* #leftpanel ul li:hover {
		background-color: rgba(255,255,255,0.1);
	} */

	ul#mainmenu {
		margin-top: 30px;
		margin-bottom: 30px;
	}

	ul#formslist {
		margin-top: 0px;
		margin-bottom: 0px;
	}

	#content {
		flex: 1;
	}
</style>