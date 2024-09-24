<script>
    let username = "";
    let password = "";

    function register() {
        // Add your registration logic here
        console.log("Registering user:", username);
        console.log("Password:", password);
    }
</script>

<main>
    <h1>Register</h1>
    <form method="POST">
        <label>
            Username:
            <input type="text" name="username" />
        </label>
        <label>
            Password:
            <input type="password" name="password" />
        </label>
        <button type="submit">Register</button>
    </form>
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100vh;
        font-family: Arial, sans-serif;
    }

    h1 {
        margin-bottom: 1rem;
    }

    form {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    label {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    input {
        padding: 0.5rem;
        border: 1px solid #ccc;
        border-radius: 4px;
    }

    button {
        padding: 0.5rem 1rem;
        background-color: #007bff;
        color: #fff;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }
</style>
