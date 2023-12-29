document.getElementById('chatForm').addEventListener('submit', function(e) {
    e.preventDefault();

    var userInput = document.getElementById('userInput').value;
    fetch('/ask', {
        method: 'POST',
        headers: {
            'Content-Type': 'text/plain',
        },
        body: userInput,
    })
    .then(response => response.text())
    .then(data => {
        document.getElementById('response').innerText = "Response: " + data;
    })
    .catch(error => {
        console.error('Error:', error);
        document.getElementById('response').innerText = "Error: " + error;
    });
});
