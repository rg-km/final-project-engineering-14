import React from "react";
import { Routes, Route } from "react-router-dom";
// import Navbar from "./Components/Navbar/Navbar";
import { Landing, Login, Register, HomePage } from "./Pages";
// import axios from "axios";

export const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Landing />}>
        {" "}
      </Route>
      <Route path="/login" element={<Login />}></Route>
      <Route path="/Register" element={<Register />}></Route>
      <Route path="/HomePage" element={<HomePage />}></Route>
    </Routes>
  );
};
export default App;
