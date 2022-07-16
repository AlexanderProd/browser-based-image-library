import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import FilesystemPage from "./pages/Filesystem";
import ShufflePage from "./pages/Shuffle";
import "./App.css";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<FilesystemPage />} />
        <Route path="shuffle" element={<ShufflePage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
