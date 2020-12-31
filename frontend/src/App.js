import React, { useState, useEffect, useRef } from 'react';

function App() {
  const [status, setStatus] = useState(false);
  const [username, setUsername] = useState('');
  const [messages, setMessages] = useState([]);
  const ws = useRef(null);

  useEffect(() => {
    if (!ws.current) return;

    ws.current.onmessage = (e) => {
      const message = JSON.parse(e.data);

      setMessages(messages.concat(message));

      console.log("e", message);
    };
  });

  const handleUsername = (e) => {
    e.preventDefault();
    const username = e.target.username.value;
    setUsername(username);
    setStatus(prevStatus => !prevStatus);

    ws.current = new WebSocket("ws://localhost:8000/ws");
  }

  const handleMessages = (e) => {
    e.preventDefault();
    const message = e.target.message.value;
    const toSend = {"username": username, "message": message};

    ws.current.send(JSON.stringify(toSend))
  }

  return (
    <body>
      <header>
        <nav>
          <div>
            <p>Simple Chat</p>
          </div>
        </nav>
      </header>
      <main>
        <div>
          {status &&
          <div>
            <div>
            {messages.map((message) => {
              return (
              <li>
                  <p>{message.username}</p>
                  <p>{message.message}</p>
              </li>
              )
            })
            }
            </div>
            <form onSubmit={handleMessages}>
              <input type='text' name='message' placeholder='enter message' />
              <input type='submit' value='enter' />
            </form>
          </div>
          }
          {!status &&
          <form onSubmit={handleUsername}>
            <input type='text' name='username' placeholder='enter username' />
            <input type='submit' value='enter' />
          </form>
          }
        </div>
      </main>
    </body>
  )
}

export default App;