document.addEventListener('DOMContentLoaded', () => {
  const input = document.getElementById('longUrl');
  const button = document.getElementById('shortenBtn');
  const result = document.getElementById('result');

  button.addEventListener('click', async () => {
    const url = input.value.trim();
    if (!url) {
      result.textContent = 'Please enter a valid URL.';
      return;
    }

    try {
      const res = await fetch('https://url-shortener-production-1ab1.up.railway.app/shorten', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ url })
      });

      const data = await res.json();

      if (data.short_url) {
        const fullURL = `https://url-shortener-production-1ab1.up.railway.app/redirect/${data.short_url}`;
        result.innerHTML = `
          <strong>Short URL:</strong><br>
          <a href="${fullURL}" target="_blank">${fullURL}</a>
          <br>
          <button id="copyBtn">Copy</button>
        `;
        document.getElementById('copyBtn').onclick = () => {
          navigator.clipboard.writeText(fullURL);
          result.innerHTML += "<br><em>Copied to clipboard!</em>";
        };
      } else {
        result.textContent = 'Shortening failed.';
      }
    } catch (error) {
      result.textContent = 'Error: ' + error.message;
    }
  });
});
