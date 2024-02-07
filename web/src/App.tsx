import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import styles from "./styles/global.module.css";
import Home from "./views/Home";
import { NotFound } from './views/404'

function App() {
  return (
    <div className={styles.container}>
      <BrowserRouter>
        <Routes>
          <Route index element={<Home />} />
            <Route path="*" element={<NotFound />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
