import { StrictMode } from "react";
import ReactDOM from "react-dom/client";

import "./index.css";
import { App } from "./App";

// biome-ignore lint/style/noNonNullAssertion: always there
const rootElement = document.getElementById("root")!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <StrictMode>
      <App />
    </StrictMode>
  );
}
