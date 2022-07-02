<script>
    import FilePond, { registerPlugin, supported } from 'svelte-filepond';

    import FilePondPluginImageExifOrientation from 'filepond-plugin-image-exif-orientation';
    import FilePondPluginImagePreview from 'filepond-plugin-image-preview';
    import FilePondPluginImageTransform from 'filepond-plugin-image-transform';
    import FilePondPluginImageResize from 'filepond-plugin-image-resize';
    import FilePondPluginImageCrop from 'filepond-plugin-image-crop';
    
    registerPlugin(FilePondPluginImageExifOrientation);
    registerPlugin(FilePondPluginImagePreview);
    registerPlugin(FilePondPluginImageTransform);
    registerPlugin(FilePondPluginImageResize);
    registerPlugin(FilePondPluginImageCrop);

    let pond;

    let name = 'image-upload';

    function handleInit() {
	    console.log('FilePond has initialised');
    }

    function handleAddFile(err, fileItem) {
	    console.log('A file has been added', fileItem);
    }

    function afterUpload(err, file) {
        window.location = `${window.location.origin}/${file.serverId}`;
    }
</script>

<style>
    @import 'filepond/dist/filepond.css';
    @import 'filepond-plugin-image-preview/dist/filepond-plugin-image-preview.css';

</style>

<svelte:head>
    <title>thread-art-website</title>
</svelte:head>

<div class="section is-family-monospace">
    <div class="container is-max-desktop">
        <h1 class="title">thread-art-website</h1>
        <p class="subtitle">Hello</p>
    </div>
</div>

<div class="section is-family-monospace">
    <div class="container is-max-desktop">
        <div class="app">

            <FilePond bind:this={pond} {name}
                server="http://localhost:8001/upload"
                allowMultiple={false}
                oninit={handleInit}
                onaddfile={handleAddFile}
                onprocessfile={afterUpload}
                dropValidation={true}
                instantUpload={false}
                allowRevert={false}
                imageResizeTargetWidth={400}
                imageResizeTargetHeight={400}
                imageCropAspectRatio={1}
                allowReplace={false}/>
        </div>
    </div>
</div>
