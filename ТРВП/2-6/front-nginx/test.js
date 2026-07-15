async function testApi() {
    const response = await fetch("/api");
    const data = await response.json();
    document.getElementById("result").innerText = data.message;
}