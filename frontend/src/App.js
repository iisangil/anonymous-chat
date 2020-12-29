import React, { useState, useEffect } from 'react';

function App() {
  const [status, setStatus] = useState(false);
  const [username, setUsername] = useState('');
  const [messages, setMessages] = useState([{ "username": "john", "message": "hi there" }, { "username": "josh", "message": "Never taalk to me again you fucking weirdo" }]);

  useEffect(() => {
    console.log(status);
    console.log(messages);
  });

  const handleUsername = (e) => {
    e.preventDefault();
    const username = e.target.username.value;
    setUsername(username);
    setStatus(prevStatus => !prevStatus);
  }

  const handleMessages = (e) => {
    e.preventDefault();
    console.log("not implemented yet");
  }

  return (
    <body>
      <header>
        <nav>
          <div>
            <p>Anonymous Chat</p>
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
              <input type='text' name='username' placeholder='enter message' />
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