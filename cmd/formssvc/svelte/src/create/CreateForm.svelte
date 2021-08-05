<script>
    import SchemaBuilder from './SchemaBuilder.svelte'
    import {push} from 'svelte-spa-router'
    import {createEventDispatcher} from 'svelte'
    const dispatch = createEventDispatcher()

    export let endpoint = ""
    let id = ''
    let name = ''
    let type = ''
    let schema

    let availableTypes = [
        "Structured",
        "Unstructured"
    ]

    const create = async () => {

        let req = {
            "id": id,
            "name": name,
            "type": type,
        }

        if(type == "Structured") {
            req["schema"] = schema
        }

        let response = await fetch(endpoint, {
            method: "POST",
            body: JSON.stringify(req)
        })
        
        if(response.ok) {
            let form = await response.json()
            dispatch('routeEvent', {
                "type": "form-created",
                "form": form
            })

            push("/forms/"+form.id)
        }
        
        return
    }
</script>
<h1>Create a new form</h1>
<form>
    <div class="line">
        <label for="id">Name (letters and numbers) :</label><input class="full-w" type="text" name="id" bind:value={id} />
    </div>
    <div class="line">
        <label for="id">Form type :</label>
        {#each availableTypes as availableType}
            <div class="row">
                <input type="radio" bind:group={type} name="type" value={availableType} />
                {availableType}
            </div>
        {/each}
    </div>
    <div class="line">
        {#if type == "Structured"}
            <label for="schema">Form schema (following <a href="https://jsontypedef.com/">JSON Type Definition</a> specification)</label>
            <SchemaBuilder id="schema" bind:schema />
        {/if}
    </div>
    <button on:click|preventDefault={create}>
        Create form
    </button>
</form>
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

    form {
        text-align: center;
        padding: 50px;
        max-width: 800px;
        margin: auto;
    }

    form .line {
        margin: 10px;
    }

    form .line label {
        font-weight: bold;
        color: #243887;
    }

    form input.full-w {
        width: 100%;
    }
</style>