import React from "react";
import { Routes, Route } from "react-router-dom";
import Navbar from "./Components/Navbar/Navbar";
import { Landing, Login, Register, Home, Question, Language } from "./Pages";

export const App = () => {
	return (
		<Routes>
			<Route path="/" element={<Language />}>
				{" "}
			</Route>
			<Route path="/login" element={<Login />}></Route>
			<Route path="/register" element={<Register />}></Route>
			<Route path="/home" element={<Home />}></Route>
		</Routes>
	);
};
export default App;
