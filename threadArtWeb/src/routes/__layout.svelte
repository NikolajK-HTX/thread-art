<script>
    import "../app.css";
</script>

<div class="navbar mb-2 shadow-lg bg-neutral text-neutral-content rounded-box">
    <div class="px-2 mx-2 navbar-start">
        <span class="text-lg font-bold"> ThreadArt </span>
    </div>
    <div class="hidden px-2 mx-2 navbar-center lg:flex">
        <div class="flex items-stretch">
            <a href="/" class="btn btn-ghost btn-sm rounded-btn"> Hjem </a>
            <a href="/about" class="btn btn-ghost btn-sm rounded-btn"> Om </a>
            <a href="/contact" class="btn btn-ghost btn-sm rounded-btn">
                Kontakt
            </a>
        </div>
    </div>
    <div class="navbar-end" />
</div>

<slot />
