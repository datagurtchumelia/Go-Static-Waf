<h1>Go Security Middleware for SQL Injection and XSS Detection</h1>

<p>This Go code is a simple security middleware that can be integrated into your web application to detect and block SQL injection and XSS (Cross-Site Scripting) attacks. It also sends alerts to a Telegram channel when an attack is detected.</p>

<h2>How It Works</h2>

<ol>
    <li>The middleware listens for incoming HTTP requests on port 8080 (you can customize the port as needed).</li>
    <li>It reads the request body to analyze the content.</li>
    <li>If it detects a potential SQL injection attack using SQLmap-like payloads, it blocks the request, logs the attack, and sends an alert to a specified Telegram channel.</li>
    <li>If it detects a potential XSS attack using example payloads, it blocks the request, logs the attack, and sends an alert to a specified Telegram channel.</li>
    <li>If no attacks are detected, it allows the request to proceed.</li>
</ol>

<h2>Getting Started</h2>

<ol>
    <li>Clone or download the repository.</li>
    <li>Open the <code>main.go</code> file and configure the following variables:</li>
</ol>

<pre>
<code>
telegramBotToken := "Telegram_Bot_Token"
telegramChatID := "Telegram_Chat_ID"
</code>
</pre>

<ol start="3">
    <li>Run the Go code using the following command:</li>
</ol>

<pre>
<code>
go run main.go
</code>
</pre>

<ol start="4">
    <li>Your Go application will now listen for incoming HTTP requests on port 8080 (you can change the port as needed).</li>
</ol>

<h2>Usage</h2>

<p>Integrate this security middleware into your web application by setting it up as a reverse proxy or middleware to analyze incoming requests. When an attack is detected, the middleware will block the request, log the incident, and send an alert to your specified Telegram channel.</p>

<h2>Telegram Alerts</h2>

<p>To receive Telegram alerts, you'll need to create a Telegram bot and obtain the bot token. Additionally, you'll need to specify the chat ID where you want to receive alerts in the <code>telegramChatID</code> variable.</p>
