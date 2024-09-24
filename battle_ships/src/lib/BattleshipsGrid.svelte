<script>
    let gridSize = 10; // Default grid size
    let gridArray = Array(gridSize * gridSize).fill(null); // Create a grid array

    // Function to change grid size
    function setGridSize(size) {
        gridSize = size;
        gridArray = Array(gridSize * gridSize).fill(null);
    }

    // Generate alphabet letters for header (up to gridSize)
    function getAlphabetArray(size) {
        return Array.from({ length: size }, (_, i) =>
            String.fromCharCode(65 + i),
        );
    }

    let alphabetArray = getAlphabetArray(gridSize);
</script>

<div>
    <label for="size">Select grid size:</label>
    <input
        id="size"
        type="number"
        bind:value={gridSize}
        min="5"
        max="20"
        on:input={() => {
            setGridSize(gridSize);
            alphabetArray = getAlphabetArray(gridSize);
        }}
    />
</div>

<div class="grid-container" style="--grid-size: {gridSize}">
    <!-- Top-left corner empty -->
    <div></div>

    <!-- Alphabet Header -->
    {#each alphabetArray as letter}
        <div class="header-item">{letter}</div>
    {/each}

    <!-- Numbered Rows and Grid Items -->
    {#each Array(gridSize) as _, rowIndex}
        <!-- Side Numbers -->
        <div class="side-item">{rowIndex + 1}</div>

        <!-- Grid Items -->
        {#each Array(gridSize) as _, colIndex}
            <div class="grid-item"></div>
        {/each}
    {/each}
</div>

<style>
    .grid-container {
        display: grid;
        grid-template-columns: auto repeat(var(--grid-size), 1fr);
        grid-template-rows: auto repeat(var(--grid-size), 1fr);
        gap: 1px;
    }

    .grid-item {
        background-color: lightblue;
        aspect-ratio: 1;
        display: flex;
        justify-content: right;
        align-items: top;
    }

    .header-item,
    .side-item {
        background-color: lightgray;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>
