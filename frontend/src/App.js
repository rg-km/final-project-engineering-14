import React from "react";
import { Routes, Route } from "react-router-dom";
import Navbar from "./Components/Navbar/Navbar";
import { Landing, Login, Register } from "./Pages";

export const App = () => {
	return (
		<Routes>
			<Route path="/" element={<Landing />}>
				{" "}
			</Route>
			<Route path="/login" element={<Login />}></Route>
			<Route path="/register" element={<Register />}></Route>
		</Routes>
	);
};
export default App;
