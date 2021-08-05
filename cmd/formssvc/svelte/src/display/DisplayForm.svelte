<script>
    import { JSONEditor } from 'svelte-jsoneditor'
    import {push} from 'svelte-spa-router'
    import JSONTree from 'svelte-json-tree';
    import {createEventDispatcher} from 'svelte'
    const dispatch = createEventDispatcher()

    export let endpoint
    export let params = {}

    let form = {}
    let formresponses = []
    let schema = {}

    const getForm = async () => {
        //TODO display a loading overlay while retrieving data
        let response = await fetch(endpoint + params.formid)
        if(response.ok) {
            let f = await response.json()
            form = f
            console.log(form)
        }
    }

    const getFormResponses = async () => {
        formresponses = []
        let r = await fetch(endpoint + params.formid + "/responses")
        if(r.ok) {
            let responses = await r.json()
            formresponses = responses 
            console.log(responses)
        }
    }

    const deleteForm = async () => {

        if(!confirm("Are you sure you want to delete this form and it's data ? (cannot be undone)")) {
            return
        }

        let response = await fetch(endpoint + params.formid, {
            method: "DELETE"
        })

        if(response.ok) {
            dispatch('routeEvent', {
                "type": "form-deleted",
                "form": form
            })

            push("/")
        }
    }

    $: if(params.formid) { getForm(); getFormResponses() }
    $: if(form.schema) { schema = {"properties": form.schema.properties, "optionalProperties": form.schema.optionalProperties} }

</script>
<h1>{form.id}</h1>
<div class="formdefinition">
    <p>Form type : {form.type}</p>
    <!-- {#if form.type == "Structured"}<p>Schema :</p><JSONEditor json="{form.schema}" />{/if} -->
    {#if form.type == "Structured"}<p>Schema :</p><pre><code><JSONTree value={schema} /></code></pre>{/if}
    <button on:click={deleteForm}>Delete form</button>
</div>
<h2>Responses</h2>



{#each formresponses as response}
<pre>
    <code>
        <!-- {JSON.stringify(response, null, 2)} -->
        <JSONTree value={response} />
    </code>
</pre>
{/each}

<style>

	h1 {	
		border-bottom: 1px solid rgba(0,0,0,0.2);
		font-size: 1.1em;
		font-weight: bold;
		margin: 0;
		padding: 10px;
        text-align: center;
        color: #243887;
	}

    h2 {	
        font-size: 1.1em;
        font-weight: bold;
        margin: 0;
        padding: 10px;
        text-align: center;
    }

    .formdefinition {
        border-bottom: 1px solid rgba(0,0,0,0.2);
        text-align: center;
        padding: 30px;
    }

    button {
        margin-top: 20px;
    }

    pre code {
        background-color: #eee;
        border: 1px solid #243887;
        display: block;
        padding: 20px;
		padding-bottom: 5px;
        margin: auto;
        width: 500px;
        text-align: left;
    }
</style>