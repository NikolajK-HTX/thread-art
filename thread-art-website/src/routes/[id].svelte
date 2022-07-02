<script>
    import { onMount } from 'svelte';

    let id = ""

    onMount(async () => {
        id = window.location.pathname.substring(1)

    })
</script>

<style>
    .round-image {
        border-radius: 10px;
    }
</style>

<svelte:head>
    <title>Your personal page</title>
</svelte:head>

<div class="container is-max-desktop">
    <div class="box mt-6">
        <div class="columns is-flex is-justify-content-center">
            <div class="column is-narrow">
                <h1 class="is-size-2 is-family-monospace"> {id} </h1>

            </div>
        </div>
    </div>

    <figure class="image is-128x128 my-4">
        <img src="/template.png" class="is-rounded" alt="" srcset="">
    </figure>
    <h1 class="title">Din personlige kunst side</h1>
    <p>Billederne er ikke klar endnu. Kom igen senere!</p>
    <div class="columns mt-4">
        <div class="column">
            <img class="round-image" src="/template.png" alt="" srcset="">
        </div>
        <div class="column">
            <img class="round-image" src="/template.png" alt="" srcset="">
        </div>
        <div class="column">
            <img class="round-image" src="/template.png" alt="" srcset="">
        </div>
    </div>

</div>