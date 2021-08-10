<script>
    import { JSONEditor } from 'svelte-jsoneditor'
    import { isSchema, isValidSchema } from "jtd";

    export let schema

    let json = {
      "properties": {
        "username": { "type": "string" },
        "password": { "type": "string" }
      },
      "optionalProperties": {
        "age": { "type": "int32" }
      }
    }

    $: isSchema(json) && isValidSchema(json) ? schema = json : schema = null

</script>
<JSONEditor bind:json />

{#if schema == null }
  <p class="error">Invalid JSON Type Definition</p>
{/if}
<style>
  .error {
    color: #ff1300;
  }
</style>