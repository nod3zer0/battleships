<script>
    import BattleshipsGrid from "$lib/BattleshipsGrid.svelte";
    export let data;

    let gridSize = 10; // Default grid size
    let gridArray = Array(gridSize)
        .fill(0)
        .map((x) => Array(gridSize).fill(0)); // Create a grid array

    /** @type {{ name: string, parts: number[][], bitmap: number[][] } | null } */
    let pickedShip = null;
    let pickedShips = [];
    // Function to change grid size
    function setGridSize(size) {
        gridSize = size;
        gridArray = Array(gridSize)
            .fill(0)
            .map((x) => Array(gridSize).fill(0));
    }

    // Generate alphabet letters for header (up to gridSize)
    function getAlphabetArray(size) {
        return Array.from({ length: size }, (_, i) =>
            String.fromCharCode(65 + i),
        );
    }

    function enter(event, rowIndex, colIndex) {
        event.preventDefault();
        gridArray[rowIndex][colIndex] = 1;

        if (!pickedShip) {
            return;
        }
        pickedShip.parts.forEach((part) => {
            gridArray[rowIndex + part[0]][colIndex + part[1]] = 1;
        });
    }

    function leave(event, rowIndex, colIndex) {
        event.preventDefault();
        updateGrid();
    }

    function drop(event, rowIndex, colIndex) {
        event.preventDefault();
        pickedShips.push({ ship: pickedShip, position: [rowIndex, colIndex] });
        pickedShip = null;
        updateGrid();
    }

    function updateGrid() {
        gridArray = Array(gridSize)
            .fill(0)
            .map((x) => Array(gridSize).fill(0));

        if (!pickedShips) {
            return;
        }

        pickedShips.forEach((ship) => {
            ship.ship.parts.forEach((part) => {
                gridArray[ship.position[0] + part[0]][
                    ship.position[1] + part[1]
                ] = 1;
            });
        });
    }

    let alphabetArray = getAlphabetArray(gridSize);
</script>

<h1>user session ID</h1>
<!-- <div>{data.user.username}</div> -->

<div class="parent">
    <div class="ship-container">
        <h2>Ships</h2>
        <ul>
            {#each data.ships as ship}
                <li>
                    {ship.name}
                    <div
                        on:dragstart={() => {
                            pickedShip = ship;
                        }}
                        draggable={true}
                        class="grid-container-ship"
                        style="--grid-size: {4}"
                    >
                        <!-- Numbered Rows and Grid Items -->
                        {#each Array(5) as _, rowIndex}
                            <!-- Grid Items -->

                            {#each Array(5) as _, colIndex}
                                {#if ship.bitmap[rowIndex][colIndex] == 1}
                                    <div class="grid-item-ship"></div>
                                {/if}
                            {/each}
                        {/each}
                    </div>
                </li>
            {/each}
        </ul>
    </div>
    <div class="ship-grid-container child">
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
                    {#if gridArray[rowIndex][colIndex] == 1}
                        <div
                            class="grid-item-gray"
                            on:dragleave={(event) =>
                                leave(event, rowIndex, colIndex)}
                            on:drop={(event) => drop(event, rowIndex, colIndex)}
                            ondragover="return false"
                        ></div>
                    {:else}
                        <div
                            class="grid-item"
                            on:dragenter={(event) =>
                                enter(event, rowIndex, colIndex)}
                            on:drop={(event) => drop(event, rowIndex, colIndex)}
                            ondragover="return false"
                        ></div>
                    {/if}
                {/each}
            {/each}
        </div>
    </div>
</div>

<style>
    .ship-container {
        float: left;
        width: 25%;
    }

    .ship-grid-container {
        float: left;
        width: 60%;
    }

    .grid-container-ship {
        display: grid;
        grid-template-columns: auto repeat(var(--grid-size), 1fr);
        grid-template-rows: auto repeat(var(--grid-size), 1fr);
        gap: 1px;
        width: 75%;
    }

    .grid-item-ship {
        background-color: rgb(65, 74, 86);
        aspect-ratio: 1;
        display: flex;
    }

    .header-item,
    .side-item {
        background-color: lightgray;
        display: flex;
        justify-content: center;
        align-items: center;
    }

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
    .grid-item-gray {
        background-color: gray;
        aspect-ratio: 1;
        display: flex;
        justify-content: right;
        align-items: top;
    }
</style>
