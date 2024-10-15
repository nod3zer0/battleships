<script>
    import "./styles.css";
    export let data;
</script>

<nav>
    <a href="/battleships">Battleships</a>
    {#if data.user}
        <a href="/logout">Logout</a>
    {:else}
        <a href="/login">Login</a>
        <a href="/register">register</a>
    {/if}
</nav>

<slot></slot>
