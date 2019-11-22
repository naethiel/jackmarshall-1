import React from "react";
import ReactDOM from "react-dom";
import { login } from "./services/auth-api";

ReactDOM.render(<App />, document.getElementById("root"));

function App() {
  const usernameInput = React.useRef();
  const passwordInput = React.useRef();

  return (
    <div>
      <h1>Welcom to JackMarshall</h1>
      <h2>Please log in</h2>
      <form
        onSubmit={e => {
          e.preventDefault();
          return handleSubmit(usernameInput, passwordInput);
        }}
      >
        <label htmlFor="username">Username</label>
        <input
          name="username"
          type="text"
          placeholder="Drax"
          ref={usernameInput}
        />
        <label htmlFor="password">Password</label>
        <input name="password" type="password" ref={passwordInput} />
        <button type="submit">log in</button>
      </form>
    </div>
  );
}

function handleSubmit(usernameInput, passwordInput) {
  const username = usernameInput.current.value;
  const password = passwordInput.current.value;

  return login(username, password);
}
