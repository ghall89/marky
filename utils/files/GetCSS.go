package files

func GetCSS() []byte {
  css := `
  :root {
    --blue: #3274ce;
    --red: #ce3446;
    --orange: #e56045;
    --text: #757575;
    --background: white;
  }
  
  .dark {
    --blue: #5996e9;
    --red: #e75969;
    --orange: #e87c66;
    --text: white;
    --background: #2d2d2d;
  }
  
  * {
    font-family: sans-serif;
    color: var(--text);
  }
  
  body {
    background-color: var(--background);
    margin: 0;
  }
  
  header {
    margin: 0;
    padding: 0.5rem 1rem;
    background-color: var(--blue);
  }
  
  header h1 {
    color: white;
  }
  
  nav ul {
    display: flex;
  }
  
  nav ul li a {
    color: white;
    text-decoration: underline;
    margin-right: 1rem;
    font-size: 1.2rem;
  }
  
  nav ul li::before {
    content: '';
  }
  
  main {
    padding: 1rem;
  }
  
  /* Typography */
  
  h1,
  h2,
  h3,
  h4 {
    font-weight: bold;
  }
  
  /* Links */
  
  a {
    color: var(--blue);
    text-decoration: none;
    transition: opacity 0.15s;
  }
  
  a:hover {
    opacity: 0.6;
  }
  
  /* List Styles */
  
  ul,
  ol {
    list-style: none;
  }
  
  ol li {
    counter-increment: li;
  }
  
  ul li::before {
    content: '‣';
    color: var(--orange);
    display: inline-block;
    width: 1em;
    margin-left: -1em;
  }
  
  ol li::before {
    content: '.' counter(li);
    color: var(--orange);
    display: inline-block;
    width: 1em;
    margin-left: -1.5em;
    margin-right: 0.5em;
    text-align: right;
    direction: rtl;
  }
  
  li {
    margin: 0.5rem 0;
  }
  
  /* Button Styles */
  button {
    background-color: gray;
    color: white;
    font-size: 1rem;
    padding: 0.6rem 1.2rem;
    outline: none;
    border: none;
    border-radius: 5px;
    transition: all 0.15s;
  }
  
  button:hover {
    opacity: 0.8;
  }
  
  button:active {
    opacity: 0.9;
    scale: 0.98;
  }
  
  /* Text Inputs */
  
  input,
  textarea {
    appearance: none;
    outline: none;
    background-color: #e6e6e6;
    color: --text;
    border: 2px solid #dadada;
    font-size: 1rem;
    padding: 0.6rem;
    border-radius: 5px;
    transition: background 0.15s;
  }
  
  input:hover,
  textarea:hover {
    background-color: #eaeaea;
  }
  
  input:focus,
  textarea:focus {
    background-color: #eaeaea;
    box-shadow: 0px 0px 2px 1px var(--blue);
  }
  
  dialog {
    border: none;
    border-radius: 5px;
    box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
    padding: 1.4rem;
    text-align: center;
    max-width: 900px;
  }
  
  /* Classes */
  
  .width-full {
    width: 100%;
  }
  
  .blue-btn {
    background-color: var(--blue);
  }
  
  .red-btn {
    background-color: var(--red);
  }
`
  
  return []byte(css)
}

