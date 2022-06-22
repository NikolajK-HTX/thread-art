<script>
    let files;
    let errorMessage = null;

    function upload() {
        if(!files || !files[0]) {
            return
        }
        
        let url = "http://localhost:8001/upload";
        let formData = new FormData();

        formData.append("imageupload", files[0]);
        console.log("hello")

        fetch(url, {
            method: "POST",
            body: formData,
        })
            .then(() => {
                /* Done. Inform the user */
            })
            .catch(() => {
                /* Error. Inform the user */
            });
    }
</script>

<svelte:head>
    <title>thread-art-website</title>
</svelte:head>

<div class="section">
    <div class="container is-max-desktop">
        <h1 class="title">thread-art-website</h1>
        <p class="subtitle">Hello</p>
    </div>
</div>

<div class="section">
    <div class="container is-max-desktop">
        <div class="file is-boxed is-primary is-large">
            <label class="file-label">
                <input
                    class="file-input"
                    id="fileupload"
                    type="file"
                    name="resume"
                    accept="image/png, image/jpeg, image/webp"
                    bind:files
                    on:change="{upload}"
                />
                <span class="file-cta">
                    <span class="file-label"> Choose an image... </span>
                </span>
                {#if errorMessage}
                    <span class="file-label">{errorMessage}</span>
                {/if}
            </label>
        </div>
    </div>
</div>
